package gogit

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// Add handles adding files to the repository's index.
func Add(path string) error {
	ignorePatterns, err := readGogitignore()
	if err != nil {
		return fmt.Errorf("error reading .gogitignore: %w", err)
	}
	// 1. Read the index file once into memory.
	indexEntries, err := ReadIndex()
	if err != nil {
		return fmt.Errorf("error reading index: %w", err)
	}

	pathsChan := make(chan string, 100) // Buffered channel to prevent the walker from blocking immediately

	// Use a WaitGroup to wait for all workers to finish
	var wgWorkers sync.WaitGroup

	// Start a fixed number of workers (e.g., 4)
	numWorkers := 4
	for i := 1; i <= numWorkers; i++ {
		wgWorkers.Add(1)
		go worker(pathsChan, indexEntries, &wgWorkers)
	}

	// Start a goroutine to walk the directory and close the channel when done
	var wgWalker sync.WaitGroup
	wgWalker.Add(1)
	go func() {
		defer wgWalker.Done()
		discoverFiles(pathsChan, ignorePatterns, path)
	}()

	// Wait for the walker to complete its job of sending paths
	wgWalker.Wait()

	// Wait for all workers to finish processing all paths from the closed channel
	wgWorkers.Wait()

	// 3. Write the updated index back to the file once.
	if err := WriteIndex(indexEntries); err != nil {
		return fmt.Errorf("error writing index file: %w", err)
	}

	return nil
}

func worker(pathsChan <-chan string, indexEntries map[string]string, wg *sync.WaitGroup) {
	defer wg.Done()
	for path := range pathsChan {
		if err := processFile(path, indexEntries); err != nil {
			log.Printf("Error processing file %s: %v\n", path, err)
		}
	}
}

func discoverFiles(pathsChan chan string, ignorePatterns []string, cmdfilePath string) error {
	err := filepath.WalkDir(cmdfilePath, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			// Handle error from WalkDir (e.g., permission denied)
			log.Printf("Error walking path %s: %v\n", path, err)
			return err
		}

		// Ignore the .gogit directory
		if info.IsDir() && info.Name() == ".gogit" {
			return filepath.SkipDir
		}
		if info.IsDir() && info.Name() == ".git" {
			return filepath.SkipDir
		}

		// Check against .gogitignore patterns
		ignored, err := isIgnored(path, ignorePatterns)
		if err != nil {
			return fmt.Errorf("error checking ignore patterns for %s: %w", path, err)
		}
		if ignored {
			if info.IsDir() {
				return filepath.SkipDir // Skip directory and its contents
			}
			return nil // Skip file
		}

		// Ignore other directories (that are not explicitly ignored by .gogitignore)
		if info.IsDir() {
			return nil
		}

		// Normalize path for consistent checks
		normalizedPath := filepath.ToSlash(path)
		if strings.HasPrefix(normalizedPath, ".gogit/") {
			return nil
		}

		if !info.IsDir() {
			// Send file path to the channel
			pathsChan <- path
		}
		return nil
	})
	if err != nil {
		log.Printf("Error during filepath.WalkDir: %v\n", err)
	}
	close(pathsChan) // Close the channel to signal workers that no more paths are coming

	return nil
}

// processFile handles hashing a single file and adding it to the in-memory index map.
func processFile(filePath string, indexEntries map[string]string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file %s: %w", filePath, err)
	}

	blobHash, buffer, err := HashObject(content)
	if err != nil {
		return fmt.Errorf("error hashing file: %w", err)
	}

	firstTwo := blobHash[:2]
	rest := blobHash[2:]
	objectPath := filepath.Join(ObjectsPath, firstTwo, rest)

	// Check if object exists. If not, create it.
	if _, err := os.Stat(objectPath); os.IsNotExist(err) {
		if err := os.MkdirAll(filepath.Join(ObjectsPath, firstTwo), 0755); err != nil {
			return fmt.Errorf("error creating directory %s: %w", filepath.Join(ObjectsPath, firstTwo), err)
		}
		if err := os.WriteFile(objectPath, buffer.Bytes(), 0644); err != nil {
			return fmt.Errorf("error writing blob object to %s: %w", objectPath, err)
		}
	} else if err != nil {
		return fmt.Errorf("error checking object existence at %s: %w", objectPath, err)
	}

	// Update the in-memory map
	indexEntries[filePath] = blobHash
	return nil
}

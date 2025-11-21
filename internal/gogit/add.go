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

// Add stages files into the index (equivalent to `git add`).
// It walks the given path, respects .gogitignore, computes SHA-1 hashes,
// writes new blob objects when needed, and updates the index only for changed files.
func Add(path string) error {
	// Load ignore rules
	ignorePatterns, err := readGogitignore()
	if err != nil {
		return fmt.Errorf("reading .gogitignore: %w", err)
	}

	// Load current index into memory
	indexEntries, err := ReadIndex()
	if err != nil {
		return fmt.Errorf("reading index: %w", err)
	}

	// Channels
	pathsChan := make(chan string, 100)       // Files discovered by walker
	resultsChan := make(chan FileResult, 100) // Results from workers

	// Worker pool
	var wgWorkers sync.WaitGroup
	numWorkers := 4 // Adjust based on CPU cores / typical workload
	for range numWorkers {
		wgWorkers.Add(1)
		go worker(pathsChan, resultsChan, &wgWorkers)
	}

	// Single collector goroutine — the ONLY one that writes to indexEntries
	var wgCollector sync.WaitGroup
	wgCollector.Add(1)
	go func() {
		defer wgCollector.Done()
		for result := range resultsChan {
			if result.Err != nil {
				log.Printf("Failed to stage %s: %v", result.Path, result.Err)
				continue
			}
			// Only update index if hash changed
			if oldHash, exists := indexEntries[result.Path]; !exists || oldHash != result.Hash {
				indexEntries[result.Path] = result.Hash
			}
		}
	}()

	// Start directory walker (closes pathsChan when done)
	go func() {
		_ = discoverFiles(pathsChan, ignorePatterns, path)
	}()

	// Wait for workers to finish processing all paths
	wgWorkers.Wait()
	close(resultsChan) // No more results → collector exits
	wgCollector.Wait() // Wait for collector to finish updating the map

	// Persist updated index to disk
	if err := WriteIndex(indexEntries); err != nil {
		return fmt.Errorf("writing index: %w", err)
	}

	return nil
}

// FileResult is the structure sent from workers to the collector
type FileResult struct {
	Path string
	Hash string
	Err  error
}

// worker reads files, computes hashes, writes blob objects if needed.
// It never touches the shared index map directly.
func worker(pathsChan <-chan string, resultsChan chan<- FileResult, wg *sync.WaitGroup) {
	defer wg.Done()
	for filePath := range pathsChan {
		content, err := os.ReadFile(filePath)
		if err != nil {
			resultsChan <- FileResult{Path: filePath, Err: fmt.Errorf("read: %w", err)}
			continue
		}

		blobHash, buffer, err := HashObject(content)
		if err != nil {
			resultsChan <- FileResult{Path: filePath, Err: fmt.Errorf("hash: %w", err)}
			continue
		}

		// Ensure the blob object exists in .gogit/objects/
		firstTwo := blobHash[:2]
		rest := blobHash[2:]
		objectPath := filepath.Join(ObjectsPath, firstTwo, rest)

		if _, err := os.Stat(objectPath); os.IsNotExist(err) {
			objDir := filepath.Join(ObjectsPath, firstTwo)
			if err := os.MkdirAll(objDir, 0755); err != nil {
				resultsChan <- FileResult{Path: filePath, Err: fmt.Errorf("mkdir: %w", err)}
				continue
			}
			if err := os.WriteFile(objectPath, buffer.Bytes(), 0644); err != nil {
				resultsChan <- FileResult{Path: filePath, Err: fmt.Errorf("write object: %w", err)}
				continue
			}
		} else if err != nil {
			resultsChan <- FileResult{Path: filePath, Err: fmt.Errorf("stat object: %w", err)}
			continue
		}

		// Success — send result
		resultsChan <- FileResult{
			Path: filePath,
			Hash: blobHash,
			Err:  nil,
		}
	}
}

// discoverFiles walks the directory tree and sends regular file paths to pathsChan.
// It respects .gogitignore and never stages anything inside .gogit.
func discoverFiles(pathsChan chan<- string, ignorePatterns []string, rootPath string) error {
	defer close(pathsChan)

	return filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Printf("Access error %s: %v", path, err)
			return nil // Continue walking
		}

		// Always skip the .gogit directory completely
		if d.IsDir() && (d.Name() == ".gogit" || d.Name() == ".git") {
			return filepath.SkipDir
		}

		// Apply .gogitignore rules
		if ignored, _ := isIgnored(path, ignorePatterns); ignored {
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		// Only stage regular files
		if d.IsDir() {
			return nil
		}

		// Extra safety: never stage files inside .gogit
		if strings.HasPrefix(filepath.ToSlash(path), ".gogit/") {
			return nil
		}

		pathsChan <- path
		return nil
	})
}

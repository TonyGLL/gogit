package gogit

import (
	"fmt"
	"os"
	"path/filepath"
)

// InitRepo contains the logic to initialize the repository directory structure.
// It receives the path where the repository will be created.
func InitRepo(path string) error {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return fmt.Errorf("%s", path)
	}

	// Check if it already exists
	if _, err := os.Stat(RepoPath); !os.IsNotExist(err) {
		return fmt.Errorf("gogit repository already exists in %s", path)
	}

	// Create necessary directories
	dirs := []string{OBJECTS, REF_HEADS}
	for _, dir := range dirs {
		if err := os.MkdirAll(filepath.Join(RepoPath, dir), 0755); err != nil {
			return fmt.Errorf("error creating directory %s: %w", dir, err)
		}
	}

	// Create initial files like index
	indexContent := []byte("")
	if err := os.WriteFile(IndexPath, indexContent, 0644); err != nil {
		return fmt.Errorf("error creating HEAD file: %w", err)
	}

	// Create initial files like HEAD
	// By default, HEAD points to the 'main' branch (or 'master')
	content := []byte("ref: refs/heads/main\n")
	if err := os.WriteFile(HeadPath, content, 0644); err != nil {
		return fmt.Errorf("error creating HEAD file: %w", err)
	}

	// Create initial files like HEAD
	// By default, HEAD points to the 'main' branch (or 'master')
	mainContent := []byte("")
	if err := os.WriteFile(RefHeadsMainPath, mainContent, 0644); err != nil {
		return fmt.Errorf("error creating main file: %w", err)
	}

	// Create .gogitignore file
	gogitignoreContent := []byte(".gogit\n.git\nmaind")
	if err := os.WriteFile(IgnorePath, gogitignoreContent, 0644); err != nil {
		return fmt.Errorf("error creating .gogitignore file: %w", err)
	}

	// Create .gogitignore file
	gogitCContent := "[credential]\n\thelper = store\n[init]\n\tdefaultBranch = master\n"
	gogitconfigContent := []byte(gogitCContent)
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("cannot get user home directory: %w", err)
	}
	configPath := filepath.Join(home, GLOBAL_CONFIG)
	if err := os.WriteFile(configPath, gogitconfigContent, 0644); err != nil {
		return fmt.Errorf("error creating .gogitconfig file: %w", err)
	}

	fmt.Printf("Initializing empty GoGit repository in %s/%s\n", wd, RepoPath)
	return nil
}

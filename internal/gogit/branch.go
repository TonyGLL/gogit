package gogit

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ListBranches() error {
	branches := []string{}
	headRef, err := GetHeadRef()
	if err != nil {
		return err
	}
	currentBranch := strings.TrimPrefix(headRef["ref:"], "refs/heads/")
	err = filepath.WalkDir(RefHeadsPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Printf("Access error %s: %v", path, err)
			return nil // Continue walking
		}

		// Only stage regular files
		if d.IsDir() {
			return nil
		}

		branchName, err := filepath.Rel(RefHeadsPath, path)
		if err != nil {
			return fmt.Errorf("error getting branch name: %w", err)
		}
		branchName = filepath.ToSlash(branchName)
		branches = append(branches, branchName)
		return nil
	})
	if err != nil {
		return fmt.Errorf("error listing branches: %w", err)
	}

	PrintBranches(branches, currentBranch)
	return nil
}

func CreateBranch(name string) error {
	currentHash, err := GetBranchHash()
	if err != nil {
		return err
	}

	branchRefPath := filepath.Join(RefHeadsPath, name)
	if err := os.MkdirAll(filepath.Dir(branchRefPath), 0755); err != nil {
		return fmt.Errorf("error creating branch directories: %w", err)
	}

	_, err = os.Stat(branchRefPath)
	if err == nil {
		return fmt.Errorf("fatal: a branch named '%s' already exists", name)
	}
	if !os.IsNotExist(err) {
		return fmt.Errorf("error checking if branch exists: %w", err)
	}

	branchRefFile, err := os.Create(branchRefPath)
	if err != nil {
		return fmt.Errorf("error creating branch ref file: %w", err)
	}
	defer branchRefFile.Close()

	_, err = branchRefFile.WriteString(currentHash + "\n")
	if err != nil {
		return fmt.Errorf("error writing to branch ref file: %w", err)
	}

	fmt.Printf("branch '%s' created at %s\n", name, currentHash)

	return nil
}

func DeleteBranch(name string) error {
	headRef, err := GetHeadRef()
	if err != nil {
		return err
	}
	currentBranch := strings.TrimPrefix(headRef["ref:"], "refs/heads/")
	if name == currentBranch {
		return fmt.Errorf("error: cannot delete branch '%s' used by worktree at '%s'", name, RepoPath)
	}

	err = os.Remove(filepath.Join(RefHeadsPath, name))
	if err != nil {
		return fmt.Errorf("error: branch '%s' not found", name)
	}

	fmt.Printf("branch '%s' deleted successfully\n", name)

	return nil
}

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
	branchMap := make(map[string]bool)
	headRef, err := GetHeadRef()
	if err != nil {
		return err
	}
	headRefSplitted := strings.Split(headRef["ref:"], "/")
	currentBranch := headRefSplitted[len(headRefSplitted)-1]
	err = filepath.WalkDir(RefHeadsPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Printf("Access error %s: %v", path, err)
			return nil // Continue walking
		}

		// Only stage regular files
		if d.IsDir() {
			return nil
		}

		branchName := strings.TrimPrefix(path, RefHeadsPath+"/")
		if branchName == currentBranch {
			branchMap[branchName] = true
		} else {
			branchMap[branchName] = false
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("error listing branches: %w", err)
	}

	PrintBranches(branchMap)
	return nil
}

func CreateBranch(name string) error {

	branchRefPath := filepath.Join(RefHeadsPath, name)
	_, err := os.Stat(branchRefPath)
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

	return nil
}

func DeleteBranch(name string) error {
	headRef, err := GetHeadRef()
	if err != nil {
		return err
	}
	headRefSplitted := strings.Split(headRef["ref:"], "/")
	currentBranch := headRefSplitted[len(headRefSplitted)-1]
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

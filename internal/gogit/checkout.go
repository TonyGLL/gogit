package gogit

import (
	"fmt"
)

func CheckoutBranch(branchName string, createBranch bool) error {
	// If createBranch is true, create the branch if it does not exist.
	if createBranch {
		err := CreateBranch(branchName)
		if err != nil {
			return err
		}
	}

	// Check if branch exists
	existBranch, err := CheckIfBranchExists(branchName)
	if err != nil {
		return err
	}
	if !existBranch {
		return fmt.Errorf("error: branch '%s' does not exist", branchName)
	}

	// Load current branch tree
	var currentTreeMap map[string]string
	currentHash, err := GetBranchHash()
	if err != nil {
		return err
	}
	if currentHash != "" {
		lastCommit, err := ReadCommit(currentHash)
		if err != nil {
			return err
		}

		lastTreeHash := lastCommit.Tree
		currentTreeMap, err = ReadTree(lastTreeHash)
		if err != nil {
			return err
		}
	}

	// Load target branch tree
	var targetTreeMap map[string]string
	targetBranchHash, err := GetTargetBranchHash(branchName)
	if err != nil {
		return err
	}
	if targetBranchHash != "" {
		targetLastCommit, err := ReadCommit(targetBranchHash)
		if err != nil {
			return err
		}

		targetLastTreeHash := targetLastCommit.Tree
		targetTreeMap, err = ReadTree(targetLastTreeHash)
		if err != nil {
			return err
		}
	}

	// Safety check for uncommitted changes
	workdirMap, err := BuildWorkdirMap()
	if err != nil {
		return fmt.Errorf("could not build the working directory map: %w", err)
	}

	ignorePatterns, err := readGogitignore()
	if err != nil {
		return fmt.Errorf("error reading .gogitignore: %w", err)
	}

	filteredWorkdirMap := make(map[string]string)
	for path, hash := range workdirMap {
		ignored, err := isIgnored(path, ignorePatterns)
		if err != nil {
			return fmt.Errorf("error checking ignore patterns for %s: %w", path, err)
		}
		if !ignored {
			filteredWorkdirMap[path] = hash
		}
	}

	for path, workdirHash := range filteredWorkdirMap {
		currentHash, inCurrent := currentTreeMap[path]
		targetHash, inTarget := targetTreeMap[path]

		if inCurrent && (!inTarget || currentHash != targetHash) && workdirHash != currentHash {
			return fmt.Errorf("error: your local changes to the file '%s' would be overwritten by checkout", path)
		}
	}

	// The DIFF
	if err := ApplyDiffCheckout(currentTreeMap, targetTreeMap); err != nil {
		return fmt.Errorf("error applying diff: %w", err)
	}

	// Update HEAD to point to the new branch
	if err := UpdateHeadRef(branchName); err != nil {
		return fmt.Errorf("error updating HEAD ref: %w", err)
	}

	fmt.Printf("Switched to branch '%s'\n", branchName)

	return nil
}

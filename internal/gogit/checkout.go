package gogit

import "fmt"

func CheckoutBranch(branchName string) error {
	// Check if branch exists
	existBranch, err := CheckIfBranchExists(branchName)
	if err != nil {
		return err
	}
	if !existBranch {
		return fmt.Errorf("error: branch '%s' does not exist", branchName)
	}

	return nil
}

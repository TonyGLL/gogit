package gogit

func SetName(name string) error {
	if err := SetGoGitStyleConfig("name", name); err != nil {
		return err
	}
	return nil
}

func SetEmail(email string) error {
	if err := SetGoGitStyleConfig("email", email); err != nil {
		return err
	}
	return nil
}

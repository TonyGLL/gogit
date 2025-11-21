package gogit

import (
	"fmt"
	"log"
)

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

func GetConfig(configType string) error {
	goGitUserConfig, err := GetGoGitStyleConfig(configType)
	if err != nil {
		log.Printf("No gogit user config")
		return nil
	}

	fmt.Printf("user.name = %s\nuser.email = %s\n", goGitUserConfig.Name, goGitUserConfig.Email)

	return nil
}

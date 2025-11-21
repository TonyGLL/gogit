package gogit

import (
	"os"

	"gopkg.in/ini.v1"
)

type GitUserConfig struct {
	Name  string
	Email string
}

func SetGoGitStyleConfig(stringType, value string) error {
	cfg, err := ini.LoadSources(ini.LoadOptions{
		AllowBooleanKeys:    true,
		IgnoreInlineComment: true,
	}, os.ExpandEnv("$HOME/.gogitconfig")) // tu archivo global

	if err != nil && !os.IsNotExist(err) {
		return err
	}
	if cfg == nil {
		cfg = ini.Empty()
	}

	// [user]
	userSection, _ := cfg.NewSection("user")
	userSection.NewKey(stringType, value)

	// Guarda con el formato exacto de Git (comentarios, espacios, etc)
	return cfg.SaveTo(os.ExpandEnv("$HOME/.gogitconfig"))
}

func GetGoGitStyleConfig(section string) (GitUserConfig, error) {
	cfg, err := ini.LoadSources(ini.LoadOptions{
		AllowBooleanKeys:    true,
		IgnoreInlineComment: true,
	}, os.ExpandEnv("$HOME/.gogitconfig"))

	if err != nil && !os.IsNotExist(err) {
		return GitUserConfig{}, err
	}
	if cfg == nil {
		cfg = ini.Empty()
	}

	switch section {
	case "user":
		sec, err := cfg.GetSection(section)
		if err != nil {
			return GitUserConfig{}, err
		}

		nameKey, err := sec.GetKey("name")
		if err != nil {
			return GitUserConfig{}, err
		}

		emailKey, err := sec.GetKey("email")
		if err != nil {
			return GitUserConfig{}, err
		}

		return GitUserConfig{
			Name:  nameKey.String(),
			Email: emailKey.String(),
		}, nil
	}

	return GitUserConfig{}, nil
}

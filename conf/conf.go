package conf

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

type Configuration struct {
	Filename        string
	Key             string
	Value           string
	GithubDeployKey string
	GithubOrg       string
	GithubRepo      string
	Username        string
	Email           string
}

func InitConfigs() *Configuration {
	configs := &Configuration{
		Filename:        os.Getenv("FILE_NAME"),
		Key:             os.Getenv("KEY"),
		Value:           os.Getenv("VALUE"),
		GithubDeployKey: os.Getenv("GITHUB_DEPLOY_KEY"),
		GithubRepo:      os.Getenv("GITHUB_REPO_NAME"),
		GithubOrg:       os.Getenv("GITHUB_ORG_NAME"),
		Username:        os.Getenv("USERNAME"),
		Email:           os.Getenv("EMAIL"),
	}

	if err := configs.validate(); err != nil {
		panic(fmt.Sprintf("config validation failed - err: %v", err.Error()))
	}

	return configs
}

func (configs *Configuration) validate() error {
	// set default values or panic if required
	message := "missing env variable:"
	if configs.Filename == "" {
		return errors.Wrap(errors.New(message), "FILE_NAME")
	}
	if configs.Key == "" {
		return errors.Wrap(errors.New(message), "KEY")
	}
	if configs.Value == "" {
		return errors.Wrap(errors.New(message), "VALUE")
	}
	if configs.GithubDeployKey == "" {
		return errors.Wrap(errors.New(message), "GITHUB_DEPLOY_KEY")
	}
	if configs.GithubRepo == "" {
		return errors.Wrap(errors.New(message), "GITHUB_REPO_NAME")
	}
	if configs.GithubOrg == "" {
		return errors.Wrap(errors.New(message), "GITHUB_ORG_NAME")
	}
	if configs.Username == "" {
		return errors.Wrap(errors.New(message), "USERNAME")
	}
	if configs.Email == "" {
		return errors.Wrap(errors.New(message), "EMAIL")
	}

	return nil
}

package conf

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
)

type Configuration struct {
	Filename        string
	KeyValue        string
	GithubDeployKey string
	GithubOrg       string
	GithubRepo      string
	Username        string
	Email           string
}

func InitConfigs() *Configuration {
	configs := &Configuration{
		KeyValue:        os.Args[1],
		Filename:        os.Args[2],
		GithubRepo:      os.Args[3],
		GithubOrg:       os.Args[4],
		GithubDeployKey: os.Getenv("GITHUB_DEPLOY_KEY"),
		Username:        os.Getenv("USERNAME"),
		Email:           os.Getenv("EMAIL"),
	}

	if err := configs.validate(); err != nil {
		panic(fmt.Sprintf("config validation failed - err: %v", err.Error()))
	}

	return configs
}

func (configs *Configuration) validate() error {
	fmt.Println(configs)
	// set default values or panic if required
	message := "missing env variable:"
	if configs.Filename == "" {
		return errors.Wrap(errors.New(message), "FILE_NAME")
	}
	if configs.KeyValue == "" {
		return errors.Wrap(errors.New(message), "KeyValue")
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

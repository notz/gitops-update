package conf

import (
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
)

type Configuration struct {
	Filename        string
	Key             []string
	Value           []string
	GithubDeployKey string
	GithubOrg       string
	GithubRepo      string
	Username        string
	Email           string
}

func InitConfigs() *Configuration {
	configs := &Configuration{
		Filename:        os.Getenv("FILE_NAME"),
		Key:             getList("KEY"),
		Value:           getList("VALUE"),
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
	if len(configs.Key) == 0 {
		return errors.Wrap(errors.New(message), "KEY")
	}
	if len(configs.Value) == 0 {
		return errors.Wrap(errors.New(message), "VALUE")
	}
	if len(configs.Key) != len(configs.Value) {
		return errors.New("length of key and values does not match")
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

func getList(variable string) []string {
	return strings.Split(strings.Replace(os.Getenv(variable), " ", "", -1), ",")
}

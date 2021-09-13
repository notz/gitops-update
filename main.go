package main

import (
	"os"

	"github.com/simplycubed/gitops-update/conf"
	"github.com/simplycubed/gitops-update/runner"
)

func main() {
	config := conf.InitConfigs()

	// get home directory of current user
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	r := runner.NewRunner(homeDir)

	// set git config with user's name and email
	if err := r.Pusher.SetGitConfig(config.Username, config.Email); err != nil {
		panic(err)
	}

	// place ssh key in .ssh directory and set config on it
	if err := r.Pusher.SetSshkey(config.GithubDeployKey, homeDir); err != nil {
		panic(err)
	}

	// clone the repository that contains the yaml file
	if err := r.Pusher.CloneRepo(config.GithubOrg, config.GithubRepo); err != nil {
		panic(err)
	}

	// install yq tool to replace key value in yaml file
	if err := r.Updater.InstallUpdater(); err != nil {
		panic(err)
	}

	// replace key in yaml file
	if err := r.Updater.UpdateFile(
		config.Key, config.Value, config.Filename, config.GithubRepo); err != nil {
		panic(err)
	}

	// push the changes into the repository with appropriate commit message
	if err := r.Pusher.PushChanges(config.Key, config.Filename, config.GithubRepo); err != nil {
		panic(err)
	}
}

package runner

import (
	"github.com/simplycubed/gitops-update/conf"
	"github.com/simplycubed/gitops-update/pusher"
	"github.com/simplycubed/gitops-update/updater"
)

type Runner struct {
	updater    updater.Updater
	pusher     pusher.Pusher
	config     *conf.Configuration
	homeDir    string
	keyValPair map[string]string
}

func NewRunner(homeDir string, conf *conf.Configuration, keyValPair map[string]string) Runner {
	return Runner{homeDir: homeDir,
		config:     conf,
		keyValPair: keyValPair,
	}
}

func (r Runner) Run() error {
	// set git config with user's name and email
	if err := r.pusher.SetGitConfig(r.config.Username, r.config.Email); err != nil {
		return err
	}

	// place ssh key in .ssh directory and set config on it
	if err := r.pusher.SetSshkey(r.config.GithubDeployKey, r.homeDir); err != nil {
		return err
	}

	// clone the repository that contains the yaml file
	if err := r.pusher.CloneRepo(r.config.GithubOrg, r.config.GithubRepo); err != nil {
		return err
	}

	// install yq tool to replace key value in yaml file
	if err := r.updater.InstallUpdater(); err != nil {
		return err
	}

	// replace key in yaml file
	if err := r.updater.UpdateFile(r.config.Filename, r.config.GithubRepo, r.keyValPair); err != nil {
		return err
	}

	// push the changes into the repository with appropriate commit message
	if err := r.pusher.PushChanges(r.config.Filename, r.config.GithubRepo); err != nil {
		return err
	}

	return nil
}

package runner

import (
	"github.com/simplycubed/gitops-update/pusher"
	"github.com/simplycubed/gitops-update/updater"
)

type Runner struct {
	Updater updater.Updater
	Pusher  pusher.Pusher
	HomeDir string
}

func NewRunner(homeDir string) Runner {
	return Runner{HomeDir: homeDir}
}

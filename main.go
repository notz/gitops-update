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

	r := runner.NewRunner(homeDir, config)

	if err := r.Run(); err != nil {
		panic(err)
	}
}

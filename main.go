package main

import (
	"errors"
	"os"
	"strings"

	"github.com/simplycubed/gitops-update/conf"
	"github.com/simplycubed/gitops-update/runner"
)

func main() {
	config := conf.InitConfigs()

	keyValPair, err := extractKeyValuePair(config.KeyValue)
	if err != nil {
		panic(err)
	}

	// get home directory of current user
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	r := runner.NewRunner(homeDir, config, keyValPair)

	if err := r.Run(); err != nil {
		panic(err)
	}
}

func extractKeyValuePair(kvStr string) (map[string]string, error) {
	pair := map[string]string{}

	str := strings.Split(kvStr, ",")
	if len(str) == 0 {
		return nil, errors.New("invalid key value pairs")
	}

	for i := range str {
		s := strings.Split(str[i], ":")
		if len(s) == 2 {
			pair[strings.TrimSpace(s[0])] = strings.TrimSpace(s[1])
		}
	}

	return pair, nil
}

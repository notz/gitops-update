package updater

import (
	"fmt"
	"os/exec"
)

type Updater struct{}

func (u Updater) InstallUpdater() error {
	if err := exec.Command("wget", "https://github.com/mikefarah/yq/releases/download/v4.2.0/yq_linux_amd64", "-O", "yq").Run(); err != nil {
		return err
	}

	if err := exec.Command("chmod", "+x", "yq").Run(); err != nil {
		return err
	}

	return nil
}

func (u Updater) UpdateFile(filename, repo string, keys, values []string) error {
	for i, key := range keys {
		value := values[i]
		err := exec.Command("./yq", "e", "-i", fmt.Sprintf(".%s=%s", key, value), fmt.Sprintf("%s/%s", repo, filename)).Run()
		if err != nil {
			return err
		}
	}
	return nil
}

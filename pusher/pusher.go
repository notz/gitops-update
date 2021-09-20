package pusher

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/simplycubed/gitops-update/metadata"
	"os"
	"os/exec"
)

var config = []byte(
	`Hostname github.com
IdentityFile ~/.ssh/id_rsa
`,
)

type Pusher struct{}

func (p Pusher) CloneRepo(repoOrg, repoName string) error {
	if err := os.RemoveAll(repoName); err != nil {
		return err
	}

	if err := exec.Command("git", "clone", fmt.Sprintf("git@github.com:%s/%s.git", repoOrg, repoName)).Run(); err != nil {
		return err
	}

	return nil
}

func (p Pusher) SetGitConfig(username, email string) error {
	commands := [][]string{
		{"git", "config", "--global", "user.email", fmt.Sprintf("%s", email)},
		{"git", "config", "--global", "user.name", fmt.Sprintf("%s", username)},
	}

	for _, command := range commands {
		if err := exec.Command(command[0], command[1:]...).Run(); err != nil {
			return err
		}
	}

	return nil
}

func (p Pusher) SetSshkey(sshkey string, homeDir string) error {
	err := os.Mkdir(fmt.Sprintf("%s/.ssh", homeDir), 0777)
	if err != nil && !errors.Is(err, os.ErrExist) {
		return err
	}

	output, err := exec.Command("ssh-keyscan", "-t", "rsa", "github.com").Output()
	if err != nil {
		return err
	}

	if err := metadata.WriteToFile(fmt.Sprintf("%s/.ssh/known_hosts", homeDir), output); err != nil {
		return err
	}

	if err := metadata.WriteToFile(fmt.Sprintf("%s/.ssh/config", homeDir), config); err != nil {
		return err
	}

	if err := metadata.WriteToFile(fmt.Sprintf("%s/.ssh/id_rsa", homeDir), []byte(sshkey)); err != nil {
		return err
	}

	commands := [][]string{
		{"chmod", "600", fmt.Sprintf("%s/.ssh/id_rsa", homeDir)},
	}

	for _, command := range commands {
		if err := exec.Command(command[0], command[1:]...).Run(); err != nil {
			return err
		}
	}

	return nil
}

func (p Pusher) PushChanges(filname, repoName string) error {
	if err := os.Chdir(repoName); err != nil {
		return err
	}
	commands := [][]string{
		{"git", "add", "."},
		{"git", "commit", "-m", fmt.Sprintf("Release of keys in %s", filname)},
		{"git", "push"},
	}
	for _, command := range commands {
		output, err := exec.Command(command[0], command[1:]...).CombinedOutput()
		if err != nil {
			return err
		}

		fmt.Println(output)
	}

	return nil
}

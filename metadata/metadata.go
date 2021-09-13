package metadata

import (
	"io/ioutil"
	"strings"
)

func WriteToFile(filename string, data []byte) error {
	return ioutil.WriteFile(filename, data, 0644)
}

func GetRepoName(name string) string {
	str := strings.Split(name, "/")
	if len(str) < 2 {
		return ""
	}

	return str[1]
}

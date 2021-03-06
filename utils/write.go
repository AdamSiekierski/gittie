package utils

import (
	"fmt"
	"io/ioutil"
)

// WriteGitignore : Write given template to .gitignore file
func WriteGitignore(template map[string]string) error {
	contents := fmt.Sprintf("# Gitignore generated by Gittie with help of GitHub API\n\n%s", template["source"])

	err := ioutil.WriteFile(".gitignore", []byte(contents), 0644)

	if err != nil {
		return err
	}

	return nil
}

package utils

import (
	"encoding/json"
	"net/http"
)

// AllTemplates : Get all .gitignore templates
func AllTemplates() ([]string, error) {
	resp, err := http.Get("https://api.github.com/gitignore/templates")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var data []string

	json.NewDecoder(resp.Body).Decode(&data)

	return data, nil
}

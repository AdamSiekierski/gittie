package utils

import (
	"encoding/json"
	"net/http"
)

// GetTemplates : Get all .gitignore templates
func GetTemplates() ([]string, error) {
	resp, err := http.Get("https://api.github.com/gitignore/templates")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var data []string

	json.NewDecoder(resp.Body).Decode(&data)

	return data, nil
}

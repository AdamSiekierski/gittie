package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetTemplate : Get a .gitignore template
func GetTemplate(name string) (map[string]string, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/gitignore/templates/%s", name))

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var data map[string]string

	json.NewDecoder(resp.Body).Decode(&data)

	return data, nil
}

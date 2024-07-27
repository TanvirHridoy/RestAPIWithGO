package config

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	ConnectionString string `json:"ConnectionString"`
}

func LoadConfig() (*Config, error) {
	file, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file content
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// Remove BOM if present
	content = bytes.TrimPrefix(content, []byte("\xef\xbb\xbf"))

	var config Config
	err = json.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

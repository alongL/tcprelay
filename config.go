package main

import (
	"encoding/json"
	"os"
)

// Config for client
type Config struct {
	ListenTcp    string `json:"listentcp"`
	TargetTcp    string `json:"targettcp"`
}

func parseJSONConfig(config *Config, path string) error {
	file, err := os.Open(path) // For read access.
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(config)
}


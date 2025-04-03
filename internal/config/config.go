package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func Read() (Config, error) {
	jsonPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, fmt.Errorf("Cannot get config path")
	}
	jsonFile, err := os.ReadFile(jsonPath)
	if err != nil {
		return Config{}, fmt.Errorf("Cannot read file")
	}

	var config Config

	if err = json.Unmarshal(jsonFile, &config); err != nil {
		return Config{}, fmt.Errorf("Cannot unmarshal")
	}

	return config, nil
}

func SetUser(c Config) error {
	jsonBytes, _ := json.Marshal(c)
	jsonPath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	err = os.WriteFile(jsonPath, jsonBytes, 0644)
	return err
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Could not get home directory")
	}
	config_file_path := home + "/.gatorconfig.json"
	return config_file_path, nil
}

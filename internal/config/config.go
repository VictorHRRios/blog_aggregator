package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (cfg Config) SetUser(username string) error {
	cfg.CurrentUserName = username
	if err := write(cfg); err != nil {
		return err
	}
	return nil
}

func Read() (Config, error) {
	jsonPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, fmt.Errorf("Cannot get config path: %v", err)
	}
	jsonFile, err := os.OpenFile(jsonPath, os.O_RDONLY, 0644)
	if err != nil {
		return Config{}, fmt.Errorf("Cannot open file")
	}

	defer jsonFile.Close()

	var config Config

	decoder := json.NewDecoder(jsonFile)

	if err = decoder.Decode(&config); err != nil {
		return Config{}, fmt.Errorf("Cannot decode to string")
	}

	return config, nil
}

func write(cfg Config) error {
	jsonPath, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("Cannot get config path: %v", err)
	}
	jsonFile, err := os.OpenFile(jsonPath, os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("Cannot open file for write")
	}
	defer jsonFile.Close()

	encoder := json.NewEncoder(jsonFile)

	if err = encoder.Encode(cfg); err != nil {
		return fmt.Errorf("Cannot encode to file")
	}
	return nil
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Could not get home directory")
	}
	return path.Join(home, configFileName), nil
}

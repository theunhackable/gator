package config

import (
	"encoding/json"
	"os"
	"os/user"
	"path/filepath"
)

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	configFilePath := filepath.Join(homeDir, ".gatorconfig.json")

	return configFilePath, nil
}

func write(cfg *Config) error {

	configFilePath, err := getConfigFilePath()

	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(cfg, "", " ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(configFilePath, data, 0644); err != nil {
		return err
	}
	return nil
}

func Read() (Config, error) {
	configFilePath, err := getConfigFilePath()

	if err != nil {
		return Config{}, err
	}

	dataBytes, err := os.ReadFile(configFilePath)

	if err != nil {
		return Config{}, err
	}

	var config Config

	if err := json.Unmarshal(dataBytes, &config); err != nil {
		return Config{}, err
	}
	return config, nil
}

func (cfg *Config) SetUser() error {
	currentUser, err := user.Current()
	if err != nil {
		return err
	}

	cfg.CurrentUserName = currentUser.Username

	if err := write(cfg); err != nil {
		return err
	}
	return nil
}

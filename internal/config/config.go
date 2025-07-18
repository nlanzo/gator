package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const CONFIG_FILE_NAME = ".gatorconfig.json"


type Config struct {
	DBURL string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}


func get_config_file_path() (string, error) {
	home_dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home_dir, CONFIG_FILE_NAME), nil
}


func Read() (Config, error) {
	json_file_path, err := get_config_file_path()
	if err != nil {
		return Config{}, err
	}

	json_file, err := os.Open(json_file_path)
	if err != nil {
		return Config{}, err
	}
	defer json_file.Close()

	var config Config
	err = json.NewDecoder(json_file).Decode(&config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}

func write(cfg *Config) error {
	json_file, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	json_file_path, err := get_config_file_path()
	if err != nil {
		return err
	}
	os.WriteFile(json_file_path, json_file, 0644)
	return nil
}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username
	return write(c)
}

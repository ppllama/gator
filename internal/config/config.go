package config

import (
	"os"
	"fmt"
	"encoding/json"
)

const configFileName = "/.gatorconfig.json"

type Config struct {
	Db_url	string `json:"db_url"`
	Current_user_name string `json:"current_user_name"`
}

func Read() (Config, error) {

	config := Config{}

	ConfigFilePath, err := getConfigFilePath()
	if err != nil {
		return config, fmt.Errorf("error getting config file path: %s", err)
	}
	data, err := os.ReadFile(ConfigFilePath)
	if err != nil {
		return config, fmt.Errorf("error reading config file: %s", err)
	}
	
	if err := json.Unmarshal(data, &config); err != nil {
		return config, fmt.Errorf("error decoding config file: %s", err)
	}

	return config, nil
}

func (conf Config) SetUser(name string) error {
	conf.Current_user_name = name
	if err := write(conf); err != nil {
		return fmt.Errorf("error writing config file: %s", err)
	}
	return nil
}

func write(cfg Config) error {
	data, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("error encoding config file: %s", err)
	}

	config_path, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("error getting config file path: %s", err)
	}
	if err := os.WriteFile(config_path, data,0666); err != nil {
		return fmt.Errorf("error writing config file: %s", err)
	}

	return nil
}

func getConfigFilePath() (string, error) {
	home_dir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error getting home directory: %s", err)
	}

	config_path := home_dir + configFileName
	return config_path, nil
}
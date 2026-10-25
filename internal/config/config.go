package config
import (
	"os"
	"encoding/json"
	"fmt"
)

const configFileName = "/.gatorconfig.json"

type Config struct {
	DbURL 		string `json:"db_url"`
	CurUserName string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error getting home directory: %w", err)
	}

	filepath := homeDir + configFileName
	return filepath, nil
}

func Read() (Config, error) {
	filepath, err := getConfigFilePath()
	if err != nil {
		return Config{}, fmt.Errorf("error getting config filepath: %w", err)
	}

	data, err := os.ReadFile(filepath)
	if err != nil {
		return Config{}, fmt.Errorf("error reading config file: %w", err)
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return Config{}, fmt.Errorf("error parsing config file: %w", err)
	}

	return config, nil
}

func write(cfg *Config) error {
	filepath, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("error getting config filepath: %w", err)
	}

	data, err := json.MarshalIndent(cfg, "", "	")
	if err != nil {
		return fmt.Errorf("error encoding config to json: %w", err)
	}

	if err := os.WriteFile(filepath, data, 0600); err != nil {
		return fmt.Errorf("error writing to config file: %w", err)
	}

	return nil
}

func (c *Config) SetUser(username string) error {
	c.CurUserName = username

	return write(c)
}
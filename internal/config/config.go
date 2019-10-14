package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Config is the representation of the data contained in the config file.
type Config struct {
	GitHubAPIKey string `json:"github_api_key"`
}

func configFilepath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, ".contribs"), nil
}

// Get gets the config data.
func Get() (*Config, error) {
	cFilepath, err := configFilepath()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(cFilepath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}

		return nil, err
	}
	defer file.Close()

	config := new(Config)

	jsonDec := json.NewDecoder(file)
	err = jsonDec.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// Write writes a config to a file.
func Write(c *Config) error {
	cFilepath, err := configFilepath()
	if err != nil {
		return err
	}

	file, err := os.Create(cFilepath)
	if err != nil {
		return err
	}

	jsonEnc := json.NewEncoder(file)
	err = jsonEnc.Encode(c)
	if err != nil {
		return err
	}

	return nil
}

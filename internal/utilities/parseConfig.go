package utilities

import (
	"dlubac/photo-portfolio-generator/internal/structs"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func ParseConfig(configFilePath string) (*structs.Config, error) {
	log.Printf("Parsing config file: %v", configFilePath)
	configFile, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	config := &structs.Config{}
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

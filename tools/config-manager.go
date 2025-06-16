package tools

import (
"fmt"
"gopkg.in/yaml.v3"
"os"
)

type serverSetup struct {
	Port int `yaml:"port"`
}

type databaseCredentials struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type databaseSetup struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	Name        string `yaml:"name"`
	Schema      string `yaml:"schema"`
	Credentials databaseCredentials `yaml:"credentials"`
	Engine      string `yaml:"engine"`
	//  mongodb
	//    additional mongo db info
}

type Config struct {
	Server   serverSetup   `yaml:"server"`
	Database databaseSetup `yaml:"database"`
}

func getConfigFileName() string {
	fileTitle := "config.yaml"
	filename := GetLocalFilePath(fileTitle, GetWorkingDir(), true)

	return filename
}

func ValidateSetup(setup Config)error {
	return nil
}

func LoadConfig() (*Config, error) {
	configFile := getConfigFileName()

	if configFile == "" {
		return nil, fmt.Errorf("Couldn't find config.yaml")
	}

	fileContent,fileErr := os.ReadFile(configFile)
	if (fileErr!=nil){
		return nil,fileErr
	}

	var parsedSetup map[string]Config
	yamlErr := yaml.Unmarshal([]byte(fileContent),&parsedSetup)
	if (yamlErr!=nil){
		return nil,yamlErr
	}

	var setup  Config = parsedSetup["config"]
	setupErr := ValidateSetup(setup)
	if (setupErr!=nil){
		return nil,setupErr
	}

	return &setup,nil;
}
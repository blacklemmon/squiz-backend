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

var AllowedDBs = []string{"mongodb"}

func ValidateSetup(setup Config)error {

	//Check database engine
	if (ArrayFind(AllowedDBs, setup.Database.Engine)<0){
		return fmt.Errorf("Database engine not supported")
	}

	//Check port is valid
	if (setup.Server.Port<=0) {
		return fmt.Errorf("Invalid port: %d", setup.Server.Port)
	}

	//We shouldn't use these ports
	var ports = []int{80,443,81,82}
	if (ArrayFind(ports,setup.Server.Port)>=0) {
		fmt.Printf("!!! Please don't use the %d port...\n",setup.Server.Port)
	}

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
package config

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/ko/cms-db/cms-es/config/configuration"
	"github.com/sanchitlohia2711/go-config/config"
)

const (
	EnvConfigurationFilePath = "CASHBACKWORKER_CONFIGURATION_FILE_PATH"

	DefaultConfigurationFilePath = "./config.example.toml"

	ErrFailedToDecodeConfigurationFile = "Failed to decode configuration file: %v"

	InfoUsingDefaultConfigurationFilePath = "Using default configuration file path: %s"
)

// Configurationer is a common interface that config structs must adhere to
type Configurationer interface {
	Process()
}

//Configuration : configuration to be used
type Configuration struct {
	ES *configuration.ES `json:"es"`
}

//Process : process
func (c *Configuration) Process() {
	//Do nothing
}

var (
	//ConfigurationData var to store configuration information
	ConfigurationData     = &Configuration{}
	configurationFilePath string
)

func init() {
	var ok bool
	goEnvironment, ok := os.LookupEnv("GOENV")
	if !ok {
		goEnvironment = "development"
	}
	configurationFilePath, ok = os.LookupEnv(EnvConfigurationFilePath)
	absPath := configurationFilePath
	if !ok {
		_, filename, _, _ := runtime.Caller(0)
		fmt.Println(filename)
		configurationFilePath = path.Join(path.Dir(filename), "./json/"+goEnvironment+"/settings.json")
		//absPath, _ = filepath.Abs(configurationFilePath)
		absPath = configurationFilePath
		fmt.Printf(InfoUsingDefaultConfigurationFilePath, configurationFilePath)
	}

	err := config.Initialize(absPath, &ConfigurationData)
	if err != nil {
		fmt.Println(ErrFailedToDecodeConfigurationFile, err)
		panic(ErrFailedToDecodeConfigurationFile)
	}
}

//New Get the new configuration
func New() *Configuration {
	return ConfigurationData
}

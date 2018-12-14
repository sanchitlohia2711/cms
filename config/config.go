package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ko/cms-db/config/configuration"
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
	Global  *configuration.Global  `json:"global"`
	Log     *configuration.Log     `json:"log"`
	Kafka   *configuration.Kafka   `json:"kafka"`
	Worker  *configuration.Worker  `json:"worker"`
	MySQL   *configuration.MySQL   `json:"mysql"`
	SQS     *configuration.SQS     `json:"sqs"`
	Loyalty *configuration.Loyalty `json:"loyalty"`
	Cron    *configuration.Cron    `json:"cron"`
	Vault   *configuration.Vault   `json:"vault"`
}

//Process : process
func (c *Configuration) Process() {
	c.Global.Process()
	c.Log.Process()
	c.Loyalty.Process()
}

var (
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
		configurationFilePath = "./config/json/" + goEnvironment + "/settings.json"
		absPath, _ = filepath.Abs(configurationFilePath)
		fmt.Println(InfoUsingDefaultConfigurationFilePath, configurationFilePath)
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

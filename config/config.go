package config

import (
	"fmt"
	"github.com/tkanos/gonfig"
)

type StorageConfig struct {
	Host                string `json:"Host"`
	Port                int    `json:"Port"`
	Database            string `json:"Database"`
	Username            string `json:"Username"`
	Password            string `json:"Password"`
	FlightStopMaxNumber int    `json:"FlightStopMaxNumber"`
}

func GetConfig(params ...string) StorageConfig {

	configuration := StorageConfig{}
	env := "prod"
	if len(params) > 0 {
		env = params[0]
	}
	fileName := fmt.Sprintf("./config/%sConfig.json", env)
	err := gonfig.GetConf(fileName, &configuration)
	if err != nil {
		return configuration
	}
	return configuration

}

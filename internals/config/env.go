package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type EnvStruct struct {
	DBHost           string `mapstructure:"HOST"`
	DBPort           string `mapstructure:"PORT"`
	DBUser           string `mapstructure:"USER"`
	DBPassword       string `mapstructure:"PASSWORD"`
	Database         string `mapstructure:"DATABASE"`
	EmailApiPassword string `mapstructure:"EMAIL_API_PASSWORD"`
	ToEmailAddress   string `mapstructure:"TO_EMAIL_ADDRESS"`
	SmtpHost         string `mapstructure:"SMTP_HOST"`
	SmtpPort         string `mapstructure:"SMTP_PORT"`
	FromEmailAddress string `mapstructure:"FROM_EMAIL_ADDRESS"`
}

func loadEnv(env Env) (*EnvStruct, error) {
	envData := &EnvStruct{}

	log.Println("Comming ENv is ", env)

	viper.AutomaticEnv()
	//envConfigFileName := fmt.Sprintf(".env.%s", env)

	// log.Println("EnvConfigFileName is ",envConfigFileName)
	// viper.SetConfigFile(envConfigFileName)
	// viper.SetConfigType("env")
	// viper.AddConfigPath("./.secrets")

	envConfigFileName := fmt.Sprintf("./.secrets/.env.%s", env)
	log.Println("EnvConfigFileName is ", envConfigFileName)

	viper.SetConfigFile(envConfigFileName)
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("Config file not found.Using environment variables")
		} else {
			return nil, fmt.Errorf("failed to read config file: %w", err)
		}
	}
	err = viper.Unmarshal(envData)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}
	return envData, nil

}

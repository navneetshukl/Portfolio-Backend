package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func LoadConfig(configPath string) (conf *Config, err error) {
	env := os.Getenv("ENV")
	if env == "" {
		env = EnvLocalhost.ToString()
	}

	validEnv := VerifyEnv(env)
	if !validEnv {
		log.Println("Not a valid environment")
		return nil, fmt.Errorf("not a valid environment")
	}
	log.Println("Environment is ", env)
	conf = &Config{}

	// Environment variable for prod server
	if env == EnvProd.ToString() {
		// Production path — use only env variables

		conf.DBConfig.DATABASE = os.Getenv("DB_DATABASE")
		conf.DBConfig.DB_PORT = os.Getenv("DB_PORT")
		conf.DBConfig.DB_USER = os.Getenv("DB_USER")
		conf.DBConfig.DB_PASSWORD = os.Getenv("DB_PASSWORD")
		conf.DBConfig.DB_HOST = os.Getenv("DB_HOST")
		conf.EmailConfig.EMAIL_API_PASSWORD = os.Getenv("EMAIL_API_PASSWORD")
		conf.EmailConfig.TO_EMAIL_ADDRESS = os.Getenv("TO_EMAIL_ADDRESS")
		conf.EmailConfig.FROM_EMAIL_ADDRESS = os.Getenv("FROM_EMAIL_ADDRESS")
		conf.EmailConfig.SMTP_HOST = os.Getenv("SMTP_HOST")
		conf.EmailConfig.SMTP_PORT = os.Getenv("SMTP_PORT")

		// Validate
		if conf.DBConfig.DATABASE == "" || conf.DBConfig.DB_HOST == "" {
			return nil, fmt.Errorf("missing database env variables")
		}

		return conf, nil
	} else if env == EnvDev.ToString() { // Environment variable for dev server
		err := godotenv.Load()
		if err != nil {
			log.Println("error in loading the prod env ", err)
			return nil, err
		}

		conf.DBConfig.DATABASE = os.Getenv("DB_DATABASE")
		conf.DBConfig.DB_PORT = os.Getenv("DB_PORT")
		conf.DBConfig.DB_USER = os.Getenv("DB_USER")
		conf.DBConfig.DB_PASSWORD = os.Getenv("DB_PASSWORD")
		conf.DBConfig.DB_HOST = os.Getenv("DB_HOST")
		conf.EmailConfig.EMAIL_API_PASSWORD = os.Getenv("EMAIL_API_PASSWORD")
		conf.EmailConfig.TO_EMAIL_ADDRESS = os.Getenv("TO_EMAIL_ADDRESS")
		conf.EmailConfig.FROM_EMAIL_ADDRESS = os.Getenv("FROM_EMAIL_ADDRESS")
		conf.EmailConfig.SMTP_HOST = os.Getenv("SMTP_HOST")
		conf.EmailConfig.SMTP_PORT = os.Getenv("SMTP_PORT")

		// Validate
		if conf.DBConfig.DATABASE == "" || conf.DBConfig.DB_HOST == "" {
			return nil, fmt.Errorf("missing database env variables")
		}

		return conf, nil
	}

	viper.AutomaticEnv()

	envConfigFileName := fmt.Sprintf("config.%s", env)

	viper.SetConfigName(envConfigFileName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)

	err = viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	err = viper.Unmarshal(&conf)
	if err != nil {
		return nil, fmt.Errorf("unable to decode into config struct, %v", err)
	}

	if conf.ServerConfig.Environment == "" {
		return nil, fmt.Errorf("unable to find environment in config file")
	}

	if conf.ServerConfig.Environment == EnvLocalhost {
		envData, err := loadEnv(conf.ServerConfig.Environment)
		if err != nil {
			log.Panic("Env file is not loaded ", err)
		}

		conf.DBConfig.DATABASE = envData.Database
		conf.DBConfig.DB_PORT = envData.DBPort
		conf.DBConfig.DB_USER = envData.DBUser
		conf.DBConfig.DB_PASSWORD = envData.DBPassword
		conf.DBConfig.DB_HOST = envData.DBHost
		conf.EmailConfig.EMAIL_API_PASSWORD = envData.EmailApiPassword
		conf.EmailConfig.TO_EMAIL_ADDRESS = envData.ToEmailAddress
		conf.EmailConfig.FROM_EMAIL_ADDRESS = envData.FromEmailAddress
		conf.EmailConfig.SMTP_HOST = envData.SmtpHost
		conf.EmailConfig.SMTP_PORT = envData.SmtpPort
		return conf, nil

	}

	return nil, err
}

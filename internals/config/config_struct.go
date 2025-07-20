package config

type DBConfig struct {
	DB_HOST     string `json:"db_host"`
	DB_PORT     string `json:"db_port"`
	DB_USER     string `json:"db_user"`
	DB_PASSWORD string `json:"db_password"`
	DATABASE    string `json:"database"`
}

type ServerConfig struct {
	Environment Env `json:"environment"`
}
type EmailConfig struct {
	EMAIL_API_PASSWORD string `json:"email_api_password"`
	TO_EMAIL_ADDRESS   string `json:"to_email_address"`
	SMTP_HOST          string `json:"smtp_host"`
	SMTP_PORT          string `json:"smtp_port"`
	FROM_EMAIL_ADDRESS string `json:"from_email_address"`
}

type Config struct {
	DBConfig     DBConfig     `json:"db_config"`
	ServerConfig ServerConfig `json:"server_config"`
	EmailConfig  EmailConfig  `json:"email_config"`
}

type Env string

const (
	EnvLocalhost Env = "localhost"
	EnvDev       Env = "dev"
	EnvProd      Env = "prod"
)

func (e Env) ToString() string {
	return string(e)
}

// VerifyEnv verify that env variables should be of only 3 types
func VerifyEnv(env string) bool {
	if env == EnvDev.ToString() || env == EnvLocalhost.ToString() || env == EnvProd.ToString() {
		return true
	}
	return false
}

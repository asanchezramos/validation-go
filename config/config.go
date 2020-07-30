package config

const (
	configFile           = ".env"
	applicationName      = "VALIDATION"
	applicationEnvPrefix = "VALIDATION"
)

// Configuration struct for config
type Configuration struct {
	AppEnv   string `env:"VALIDATION_APP_ENV" envDefault:"test"`
	Secret   string `env:"VALIDATION_API_SECRET"`
	Database Database
}

// Database struct for database settings
type Database struct {
	Name   string `env:"VALIDATION_DB_NAME"`
	User   string `env:"VALIDATION_DB_USER"`
	Pass   string `env:"VALIDATION_DB_PASSWORD"`
	Host   string `env:"VALIDATION_DB_HOST"`
	Port   string `env:"VALIDATION_DB_PORT" envDefault:"3306"`
	Driver string `env:"VALIDATION_DB_DRIVER"`
}

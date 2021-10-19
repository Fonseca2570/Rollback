package globals

type Config struct {
	AppPort          int    `toml:"APP_PORT"`
	AppName          string `toml:"APP_NAME"`
	DatabaseUser     string `toml:"DATABASE_USER"`
	DatabasePassword string `toml:"DATABASE_PASSWORD"`
	DatabaseName     string `toml:"DATABASE_NAME"`
	DatabasePort     string `toml:"DATABASE_PORT"`
	Key              string `toml:"KEY"`
	RabbitUser       string `toml:"RABBITMQ_DEFAULT_USER"`
	RabbitPass       string `toml:"RABBITMQ_DEFAULT_PASS"`
	RabbitPort       string `toml:"RABBITMQ_PORT"`
	Auth             string `toml:"AUTH"`
}

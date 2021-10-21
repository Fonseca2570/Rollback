package globals

type Config struct {
	AppPort          int    `toml:"APP_PORT"`
	AppName          string `toml:"APP_NAME"`
	DatabaseUser     string `toml:"DATABASE_USER"`
	DatabasePassword string `toml:"DATABASE_PASSWORD"`
	DatabaseName     string `toml:"DATABASE_NAME"`
	DatabasePort     string `toml:"DATABASE_PORT"`
}

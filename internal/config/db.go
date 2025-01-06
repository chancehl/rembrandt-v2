package config

type DBConfig struct {
	URL string `env:"DATABASE_URL"`
}

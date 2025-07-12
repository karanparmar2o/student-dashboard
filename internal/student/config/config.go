package config

type Config struct {
	DBUrl string
}

func Load() *Config {
	return &Config{
		DBUrl: "localhost:5432", // example
	}
}

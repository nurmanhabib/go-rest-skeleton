package config

type Config struct {
	App
}

func New() *Config {
	return &Config{
		App: App{
			Env:      getEnv("APP_ENV", "production"),
			Timezone: getEnv("APP_TZ", "UTC"),
			Port:     getEnv("APP_PORT", "8080"),
			Debug:    getEnvAsBool("APP_DEBUG", false),
		},
	}
}

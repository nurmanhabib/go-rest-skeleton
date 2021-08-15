package config

type App struct {
	Env      string
	Timezone string
	Port     string
	Debug    bool
}

func (a App) IsProduction() bool {
	return a.Env == "production"
}

func (a App) IsDebugMode() bool {
	return a.IsProduction() == false && a.Debug
}

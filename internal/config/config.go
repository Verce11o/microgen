package config

type Config struct {
	App App
}

type App struct {
	Module string
}

func NewConfig(app App) *Config {
	return &Config{App: app}
}

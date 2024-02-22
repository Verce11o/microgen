package config

type Config struct {
	App    App
	Folder string
}

type App struct {
	Module string
}

func NewConfig(app App, folder string) *Config {
	return &Config{App: app, Folder: folder}
}

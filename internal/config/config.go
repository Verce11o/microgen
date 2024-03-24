package config

type Config struct {
	App         App
	ServiceName string
}

type App struct {
	Module   string
	Storages []Storage
}

type Storage string

func NewConfig(app App, serviceName string) *Config {
	return &Config{App: app, ServiceName: serviceName}
}

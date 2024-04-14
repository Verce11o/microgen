package config

type Config struct {
	App          App
	ServiceName  string
	Storage      []string
	StorageChunk []StorageChunk
}

type App struct {
	Module string
}

type StorageChunk struct {
	Name                  string
	ClientImports         []string
	ImplementationImports []string
	ClientTmpl            string
}

func NewConfig(module, serviceName string, storages []string) *Config {
	return &Config{App: App{Module: module}, ServiceName: serviceName, Storage: storages}
}

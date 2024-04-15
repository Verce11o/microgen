package config

type Chunk interface {
	GetConfigStruct() string
	GetConfigFields() string
}

type Config struct {
	App          App
	ServiceName  string
	Storage      []string
	StorageChunk []Chunk
}

type App struct {
	Module string
}

type StorageChunk struct {
	Name                  string
	ClientImports         []string
	ConfigFields          string
	ImplementationImports []string
	ClientTmpl            string
	ConfigTmpl            string
}

func NewConfig(module, serviceName string, storages []string) *Config {
	return &Config{App: App{Module: module}, ServiceName: serviceName, Storage: storages}
}

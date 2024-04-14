package steps

import (
	"github.com/Verce11o/microgen/internal/config"
	storage "github.com/Verce11o/microgen/internal/generator/storages/postgres"
	"io/fs"
)

var chunks = map[string]config.StorageChunk{
	"postgres": storage.NewPostgresChunk(),
}

type LoadChunks struct{}

func NewLoadChunks() *LoadChunks {
	return &LoadChunks{}
}

func (l *LoadChunks) Invoke(_ fs.FS, data *config.Config) error {
	for _, storageName := range data.Storage {
		chunk := chunks[storageName]
		data.StorageChunk = append(data.StorageChunk, chunk)
	}
	return nil
}

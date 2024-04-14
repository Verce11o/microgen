package storage

import (
	_ "embed"
	"github.com/Verce11o/microgen/internal/config"
)

//go:embed postgres.go.tmpl
var clientTmpl string

func NewPostgresChunk() config.StorageChunk {
	return config.StorageChunk{
		Name:                  "postgres",
		ClientImports:         []string{"\"github.com/jackc/pgx/v5/pgxpool\""},
		ImplementationImports: []string{"\"github.com/jackc/pgx/v5/pgxpool\"", "\"github.com/jackc/pgx/v5\""},
		ClientTmpl:            clientTmpl,
	}
}

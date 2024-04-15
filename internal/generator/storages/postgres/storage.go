package storage

import (
	"bytes"
	_ "embed"
	"github.com/Verce11o/microgen/internal/config"
	"text/template"
)

//go:embed postgres.go.tmpl
var clientTmpl string

//go:embed config.tmpl
var configTmpl string

type PostgresChunk struct {
	config config.StorageChunk
}

func NewPostgresChunk() *PostgresChunk {
	return &PostgresChunk{config: config.StorageChunk{
		Name:                  "postgres",
		ClientImports:         []string{"\"github.com/jackc/pgx/v5/pgxpool\""},
		ConfigFields:          "Postgres Postgres `yaml:\"postgres\"`",
		ImplementationImports: []string{"\"github.com/jackc/pgx/v5/pgxpool\"", "\"github.com/jackc/pgx/v5\""},
		ClientTmpl:            clientTmpl,
		ConfigTmpl:            configTmpl,
	}}
}

func (p *PostgresChunk) GetConfigFields() string {
	return p.config.ConfigFields
}

func (p *PostgresChunk) GetConfigStruct() string {

	buf := bytes.NewBuffer(nil)
	t := template.New("postgresConfig")
	tmpl := template.Must(t.Parse(p.config.ConfigTmpl))

	err := tmpl.Execute(buf, nil)
	if err != nil {
		return ""
	}

	return buf.String()
}

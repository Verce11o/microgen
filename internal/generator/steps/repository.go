package steps

import (
	"fmt"
	"github.com/Verce11o/microgen/internal/config"
	"io/fs"
	"os"
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) Invoke(genFS fs.FS, data *config.Config) error {
	repoDir := fmt.Sprintf("%s/internal/repository/", genFS)

	for _, repo := range data.App.Storages {
		err := os.MkdirAll(repoDir+string(repo), os.FileMode(0750))
		if err != nil {
			return err
		}
	}

	return nil
}

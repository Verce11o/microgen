package steps

import (
	"github.com/Verce11o/microgen/internal/config"
	"io/fs"
)

type Makefile struct{}

func NewMakefile() *Makefile {
	return &Makefile{}
}

func (m *Makefile) Invoke(genFS fs.FS, data *config.Config) error {

	return nil
}

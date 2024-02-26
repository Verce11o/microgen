package steps

import (
	"fmt"
	"github.com/Verce11o/microgen/internal/config"
	"io/fs"
	"os"
	"strings"
)

type RenameTmpl struct {
}

func NewRenameTmpl() *RenameTmpl {
	return &RenameTmpl{}
}

func (r *RenameTmpl) Invoke(genFS fs.FS, data *config.Config) error {
	err := fs.WalkDir(genFS, ".", func(path string, f fs.DirEntry, err error) error {
		if f.IsDir() {
			return nil
		}

		fullPath := fmt.Sprintf("%s/%s", genFS, path)

		if strings.Contains(f.Name(), ".tmpl") {
			if _, err = fs.Stat(genFS, path); !os.IsNotExist(err) {
				err = os.Rename(fullPath, strings.Replace(fullPath, ".tmpl", "", 1))
				if err != nil {
					return err
				}
			}
		}

		return nil
	})

	return err
}

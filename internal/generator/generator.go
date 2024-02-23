package generator

import (
	"fmt"
	"github.com/Verce11o/microgen/internal/config"
	"io/fs"
	"os"
)

type Generator struct {
	config   *config.Config
	skeleton *Skeleton
}

func NewGenerator(config *config.Config, skeleton *Skeleton) *Generator {
	return &Generator{config: config, skeleton: skeleton}
}

func (g *Generator) GenerateFS() error {
	err := fs.WalkDir(g.skeleton.template, g.skeleton.root, func(path string, d fs.DirEntry, _ error) error {
		savePath := path[len(g.skeleton.root):]
		if savePath == "" {
			return nil
		}
		fullPath := g.config.Folder + savePath

		if d.IsDir() {
			err := os.MkdirAll(fullPath, os.FileMode(0750))
			if err != nil {
				return fmt.Errorf("cannot create folder %s: %w", savePath, err)
			}
			return nil
		}

		content, err := g.skeleton.template.ReadFile(path)
		if err != nil {
			return err
		}

		err = os.WriteFile(fullPath, content, os.FileMode(0775))
		if err != nil {
			return err
		}

		return nil
	})

	return err

}

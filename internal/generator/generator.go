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
		fmt.Println(savePath)

		if d.IsDir() {
			err := os.MkdirAll(g.config.Folder+savePath, os.FileMode(0750))
			if err != nil {
				return fmt.Errorf("cannot create folder %s: %w", savePath, err)
			}
			return nil
		}

		return nil
	})

	return err

}

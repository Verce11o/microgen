package generator

import (
	"fmt"
	"github.com/Verce11o/microgen/internal/config"
	"io/fs"
	"os"
)

type Step interface {
	Invoke(fs fs.FS, data *config.Config) error
}

type Generator struct {
	config   *config.Config
	skeleton *Skeleton
	Steps    []Step
}

func NewGenerator(config *config.Config, skeleton *Skeleton) *Generator {
	return &Generator{config: config, skeleton: skeleton}
}

func (g *Generator) AddStep(st Step) {
	g.Steps = append(g.Steps, st)
}

func (g *Generator) generateFS() (fs.FS, error) {
	err := fs.WalkDir(g.skeleton.template, g.skeleton.root, func(path string, d fs.DirEntry, _ error) error {
		savePath := path[len(g.skeleton.root):]
		if savePath == "" {
			return nil
		}

		if _, err := os.Stat(g.config.ServiceName); os.IsNotExist(err) {
			err = os.Mkdir(g.config.ServiceName, os.FileMode(0750))
			if err != nil {
				return fmt.Errorf("cannot create root folder %s: %w", savePath, err)
			}
		}

		fullPath := g.config.ServiceName + savePath

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

	genFS := os.DirFS(g.config.ServiceName)

	return genFS, err
}

func (g *Generator) Generate() error {
	genFS, err := g.generateFS()
	if err != nil {
		return err
	}

	for _, st := range g.Steps {
		if err = st.Invoke(genFS, g.config); err != nil {
			fmt.Println(err.Error())
			return err
		}
	}

	return nil
}

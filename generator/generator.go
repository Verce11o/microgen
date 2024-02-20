package generator

import (
	"fmt"
	"github.com/Verce11o/microgen/internal/config"
	"io/fs"
)

type Generator struct {
	config   *config.Config
	skeleton *Skeleton
}

func (g *Generator) Generate() {
	err := fs.WalkDir(g.skeleton.template, g.skeleton.root, func(path string, d fs.DirEntry, err error) error {
		savePath := path[len(g.skeleton.root):]
		fmt.Println(savePath)

		return nil
	})

	if err != nil {
		panic(err)
	}
}

func NewGenerator(config *config.Config) *Generator {
	return &Generator{config: config, skeleton: NewSkeleton()}
}

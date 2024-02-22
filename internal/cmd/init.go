package cmd

import (
	"github.com/Verce11o/microgen/internal/config"
	"github.com/Verce11o/microgen/internal/generator"
)

func InitApp(config *config.Config) error {
	g := generator.NewGenerator(config, generator.NewSkeleton())
	err := g.GenerateFS()
	return err
}

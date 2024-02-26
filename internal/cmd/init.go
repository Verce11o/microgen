package cmd

import (
	"github.com/Verce11o/microgen/internal/config"
	"github.com/Verce11o/microgen/internal/generator"
	"github.com/Verce11o/microgen/internal/generator/steps"
)

func InitApp(config *config.Config) error {
	g := generator.NewGenerator(config, generator.NewSkeleton())
	g.AddStep(steps.NewCodeGen())
	g.AddStep(steps.NewRenameTmpl())

	err := g.Generate()
	return err
}

package cmd

import (
	"github.com/Verce11o/microgen/generator"
	"github.com/Verce11o/microgen/internal/config"
)

func InitApp(config *config.Config) {
	g := generator.NewGenerator(config)
	g.Generate()
}

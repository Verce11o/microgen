package generator

import (
	"github.com/Verce11o/microgen/internal/config"
)

type Generator struct {
	config config.Config
}

func (g *Generator) Generate() {

}

func NewGenerator(config config.Config) *Generator {
	return &Generator{config: config}
}

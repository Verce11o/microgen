package steps

import (
	"fmt"
	"github.com/Verce11o/microgen/internal/config"
	"io/fs"
	"os/exec"
)

type GoModInit struct {
}

func NewGoModInit() *GoModInit {
	return &GoModInit{}
}

func (g *GoModInit) Invoke(genFS fs.FS, data *config.Config) error {
	cmd := exec.Command("go", "mod", "init", data.App.Module)
	cmd.Dir = fmt.Sprintf("%s", genFS)

	return cmd.Run()
}

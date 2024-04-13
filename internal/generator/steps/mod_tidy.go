package steps

import (
	"errors"
	"fmt"
	"github.com/Verce11o/microgen/internal/config"
	"io/fs"
	"os/exec"
)

type GoModTidy struct {
}

func NewGoModTidy() *GoModTidy {
	return &GoModTidy{}
}

func (g *GoModTidy) Invoke(genFS fs.FS, _ *config.Config) error {
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = fmt.Sprintf("%s", genFS)

	output, err := cmd.CombinedOutput()
	out := string(output)
	if err != nil && out != "" {
		return errors.New(out)
	}

	return nil

}

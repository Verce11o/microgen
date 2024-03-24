package steps

import (
	"fmt"
	"github.com/Verce11o/microgen/internal/config"
	"io/fs"
	"os/exec"
)

type ProtoCompiler struct{}

func NewProtoCompiler() *ProtoCompiler {
	return &ProtoCompiler{}
}

func (p *ProtoCompiler) Invoke(genFS fs.FS, data *config.Config) error {
	e := exec.Command("make", "gen-proto")
	e.Dir = fmt.Sprintf("%s", genFS)
	return e.Run()
}

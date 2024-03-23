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
	//protoPath := fmt.Sprintf("%s/protos/proto/service", genFS)
	//newProtoPath := strings.Replace(protoPath, "service", strings.ToLower(data.ServiceName), 1)
	//err := os.Rename(protoPath, newProtoPath)
	//if err != nil && !errors.Is(err, os.ErrExist) {
	//	return err
	//}

	e := exec.Command("make", "gen-proto")
	e.Dir = fmt.Sprintf("%s", genFS)
	return e.Run()
}

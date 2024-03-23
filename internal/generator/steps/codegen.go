package steps

import (
	"bytes"
	"fmt"
	"github.com/Verce11o/microgen/internal/config"
	"github.com/Verce11o/microgen/internal/generator"
	"go/format"
	"io/fs"
	"os"
	"strings"
	"text/template"
)

type CodeGen struct {
}

func NewCodeGen() *CodeGen {
	return &CodeGen{}
}

func (c *CodeGen) parseTemplate(name string, content []byte, config any) ([]byte, error) {
	fMap := generator.FuncMap

	buf := bytes.NewBuffer(nil)
	t := template.New(name).Funcs(fMap)
	tmpl := template.Must(t.Parse(string(content)))

	err := tmpl.Execute(buf, config)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil

}

func (c *CodeGen) Invoke(genFS fs.FS, config *config.Config) error {
	err := fs.WalkDir(genFS, ".", func(path string, f fs.DirEntry, err error) error {

		if f.IsDir() {
			return nil
		}

		fullPath := fmt.Sprintf("%s/%s", genFS, path)

		content, err := os.ReadFile(fullPath)
		if err != nil {
			return fmt.Errorf("cannot read file: %w", err)
		}

		generatedCode, err := c.parseTemplate(f.Name(), content, config)
		if err != nil {
			return fmt.Errorf("cannot parse template: %w", err)
		}

		if strings.Contains(f.Name(), "go.tmpl") {
			generatedCode, err = format.Source(generatedCode)
			if err != nil {
				return fmt.Errorf("cannot format file: %w", err)
			}
		}

		err = os.WriteFile(fullPath, generatedCode, 0644)

		return err
	})

	return err

}

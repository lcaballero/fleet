package generate

import (
	"fmt"

	"os"
	"path/filepath"

	cmd "github.com/codegangsta/cli"
)

func Generate(ctx *cmd.Context) {
	fmt.Println("generate")
}

type Generator struct {
	config *Config
}

func NewGenerator(conf *Config) *Generator {
	return &Generator{
		config: conf,
	}
}

func (g *Generator) Run() error {
	src := &Source{
		Source: filepath.Join(g.config.Source, "names.html"),
		Dest: filepath.Join(g.config.Dest, "names.html"),
		Name: "source-names.html",
		Data: map[string]interface{}{
			"names": []string{
				"Superman",
				"Batman",
				"The Flash",
			},
		},
	}
	return g.GenSource(src)
}

func (g *Generator) GenSource(src *Source) error {
	t, err := src.ToTemplate()
	if err != nil {
		return err
	}

	f, err := os.OpenFile(src.Dest, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = t.Execute(f, src.Data)
	if err != nil {
		return err
	}
	err = f.Sync()
	if err != nil {
		return err
	}
	return nil
}

package create

import (
	"fmt"
	"os"
	"path/filepath"

	cmd "github.com/codegangsta/cli"
	"github.com/lcaballero/fleet/shared/errors"
)

var DefaultDirs = []string{
	"content",
	"static",
	"layouts",
	"data",
	"themes",
	"types",
	"libs",
}

func New(ctx *cmd.Context) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println(dir)
}

type Site struct {
	conf *Config
}

func NewSite(c *Config) *Site {
	return &Site{
		conf: c,
	}
}
func (t *Site) Run() error {
	errs := errors.Errors{
		t.CreateDirs(),
	}
	return errs
}

func (t *Site) CreateDirs() error {
	fq := func(dir string) string {
		return filepath.Join(t.conf.Root, dir)
	}
	mkdir := func(dir string) error {
		return os.Mkdir(fq(dir), 0777)
	}
	errs := errors.Errors{}
	for _, dir := range DefaultDirs {
		errs = errs.Add(mkdir(dir))
	}
	return errs
}

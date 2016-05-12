package create

import (
	"testing"

	. "github.com/lcaballero/fleet/shared/testing"
)

func TestNew(t *testing.T) {
	cf := &Config{Root: ".test/t1"}
	defer RemoveAll(cf.Root, DefaultDirs...)
	s := NewSite(cf)
	s.Run()

	Exists(t, cf.Root, "content")
}

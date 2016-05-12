package generate

import (
	"testing"

	. "github.com/lcaballero/fleet/shared/testing"
)

func TestGenerateRun(t *testing.T) {
	cf := &Config{
		Source: ".test/s1",
		Dest:   ".test/t1",
	}
	g := NewGenerator(cf)
	err := g.Run()
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	Exists(t, cf.Dest, "names.html")
}

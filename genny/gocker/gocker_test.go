package gocker

import (
	"testing"

	"github.com/gobuffalo/genny/gentest"
	"github.com/stretchr/testify/require"
)

func Test_New(t *testing.T) {
	r := require.New(t)

	g, err := New(&Options{
		Args: []string{"get", "-v"},
		Keep: true,
	})
	r.NoError(err)

	run := gentest.NewRunner()
	run.With(g)

	r.NoError(run.Run())

	res := run.Results()

	cmds := []string{
		"docker build . -f Gockerfile",
		"say gocker has finished",
	}
	r.NoError(gentest.CompareCommands(cmds, res.Commands))

	files := []string{"Gockerfile"}
	r.NoError(gentest.CompareFiles(files, res.Files))
}

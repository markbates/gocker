package gocker

import (
	"os/exec"

	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/plush"
	"github.com/gobuffalo/plushgen"
	"github.com/pkg/errors"
)

func New(opts *Options) (*genny.Generator, error) {
	g := genny.New()

	if err := opts.Validate(); err != nil {
		return g, errors.WithStack(err)
	}

	if err := g.Box(packr.New("../gocker/templates", "../gocker/templates")); err != nil {
		return g, errors.WithStack(err)
	}
	ctx := plush.NewContext()
	ctx.Set("opts", opts)
	g.Transformer(plushgen.Transformer(ctx))

	g.Command(exec.Command("docker", "build", ".", "-f", "Gockerfile"))

	if !opts.Keep {
		g.RunFn(func(r *genny.Runner) error {
			r.Delete("Gockerfile")
			return nil
		})
	}

	g.RunFn(func(r *genny.Runner) error {
		if _, err := r.LookPath("say"); err != nil {
			return nil
		}
		return r.Exec(exec.Command("say", "gocker has finished"))
	})

	return g, nil
}

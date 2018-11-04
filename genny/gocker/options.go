package gocker

import (
	"errors"
	"go/build"
	"os"
	"path"
	"strings"
)

type Options struct {
	GoVersion string
	GoMods    string
	WithDep   bool
	Args      []string
	Path      string
}

// Validate that options are usuable
func (opts *Options) Validate() error {
	if opts == nil {
		opts = &Options{}
	}

	if len(opts.Args) == 0 {
		return errors.New("you must pass arguments")
	}

	if len(opts.GoVersion) == 0 {
		opts.GoVersion = "latest"
	}

	if len(opts.GoMods) == 0 {
		opts.GoMods = "on"
	}

	if len(opts.Path) == 0 {
		pwd, err := os.Getwd()
		if err != nil {
			return err
		}
		opts.Path = pwd
		c := build.Default
		for _, src := range c.SrcDirs() {
			opts.Path = strings.TrimPrefix(opts.Path, src)
		}
		opts.Path = path.Join("$GOPATH/src", opts.Path)
	}

	return nil
}

func (opts *Options) GoCmd() string {
	return strings.Join(append([]string{"go"}, opts.Args...), " ")
}

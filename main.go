package main

import (
	"context"
	"flag"
	"log"

	"github.com/gobuffalo/genny"
	"github.com/markbates/gocker/genny/gocker"
	"github.com/pkg/errors"
)

var options = struct {
	*gocker.Options
	DryRun bool
}{
	Options: &gocker.Options{},
}

func main() {
	flag.BoolVar(&options.DryRun, "dry-run", false, "dry run")
	flag.BoolVar(&options.WithDep, "with-dep", false, "with dep management")
	flag.BoolVar(&options.SkipGet, "skip-get", false, "skips go get/dej")
	flag.BoolVar(&options.Keep, "keep", false, "keeps the generated Gockerfile")
	flag.StringVar(&options.GoVersion, "go-version", "latest", "go version to run")
	flag.StringVar(&options.GoMods, "go-mods", "on", "turn on/off Go Modules")
	flag.StringVar(&options.Path, "path", "", "path to use in Docker for WORKDIR")
	flag.Parse()

	if err := exec(flag.Args()); err != nil {
		log.Fatal(err)
	}
}

func exec(args []string) error {
	ctx := context.Background()
	run := genny.WetRunner(ctx)
	if options.DryRun {
		run = genny.DryRunner(ctx)
	}

	opts := options.Options
	opts.Args = args

	if err := run.WithNew(gocker.New(opts)); err != nil {
		return errors.WithStack(err)
	}

	return run.Run()
}

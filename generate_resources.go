//+build ignore

package main

import (
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/pkgmain/smooth"
	"github.com/pkgmain/smooth/internal/smooth/templates"
)

func main() {
	opts := smooth.DefaultVFSGenOptions
	opts.PackageName = "templates"
	opts.Filename = "internal/smooth/templates/resources_vfs_gen.go"

	err := smooth.GenerateResources(templates.VFS, opts)
	if err != nil {
		log.Fatal().Msgf("%+v", errors.WithStack(err))
	}
}

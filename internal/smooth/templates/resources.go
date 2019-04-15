//+build dev

package templates

import (
	"net/http"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

var VFS http.FileSystem

func init() {
	abs, err := filepath.Abs("internal/smooth/templates")
	if err != nil {
		log.Fatal().Msgf("%+v", errors.WithStack(err))
	}

	VFS = http.Dir(abs)
}

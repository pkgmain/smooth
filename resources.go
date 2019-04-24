package smooth

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gobuffalo/packr/v2"
	"github.com/pkg/errors"
	"github.com/shurcooL/httpfs/filter"
	"github.com/shurcooL/vfsgen"
)

var (
	config     = packr.New("config", "./config")
	migrations = packr.New("migrations", "./migrations")
	templates  = packr.New("templates", "./templates")
)

func SetConfigbox(box *packr.Box) {
	config = box
}

func SetMigrationsBox(box *packr.Box) {
	migrations = box
}

func SetTemplateBox(box *packr.Box) {
	templates = box
}

var DefaultVFSGenOptions = vfsgen.Options{
	PackageName:  "resources",
	Filename:     "resources/resources_vfs_gen.go",
	BuildTags:    "!dev",
	VariableName: "VFS",
}

// GenerateResources uses uses vfsgen to generate a virtual filesytem. Files
// with the .go extension will be excluded.
func GenerateResources(fs http.FileSystem, opts vfsgen.Options) error {
	filtered := filter.Skip(fs, func(path string, fi os.FileInfo) bool {
		return filepath.Ext(path) == ".go"
	})

	err := vfsgen.Generate(filtered, opts)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

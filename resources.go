package smooth

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gobuffalo/packd"
	"github.com/pkg/errors"
	"github.com/shurcooL/httpfs/filter"
	"github.com/shurcooL/httpfs/vfsutil"
	"github.com/shurcooL/vfsgen"
)

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

type PackdWalkable struct {
	fs http.FileSystem
}

func NewPackdWalkable(fs http.FileSystem) *PackdWalkable {
	return &PackdWalkable{fs: fs}
}

func (p PackdWalkable) Walk(wf packd.WalkFunc) error {
	return vfsutil.Walk(p.fs, "/", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		f, err := p.fs.Open(path)
		if err != nil {
			return errors.WithStack(err)
		}

		file, err := packd.NewFile(info.Name(), f)
		if err != nil {
			return err
		}

		err = wf(path, file)
		if err != nil {
			return errors.WithStack(err)
		}

		return nil
	})
}

func (p PackdWalkable) WalkPrefix(prefix string, wf packd.WalkFunc) error {
	return p.Walk(func(path string, file packd.File) error {
		if strings.HasPrefix(filepath.ToSlash(path), filepath.ToSlash(prefix)) {
			err := wf(path, file)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

package smooth

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gobuffalo/packd"
	"github.com/pkg/errors"
	"github.com/shurcooL/httpfs/vfsutil"
)

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

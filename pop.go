package smooth

import (
	"net/http"

	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
)

func PopConnect(fs http.FileSystem, env string) (*pop.Connection, error) {
	cfg, err := fs.Open("database.yml")
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = pop.LoadFrom(cfg)
	if err != nil {
		return nil, err
	}

	return pop.Connect(env)
}

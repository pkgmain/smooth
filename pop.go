package smooth

import (
	"net/http"

	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
)

type PopMigrator struct {
	box pop.MigrationBox
}

func NewMigrator(conn *pop.Connection, fs http.FileSystem) (m PopMigrator, _ error) {
	walkable := NewPackdWalkable(fs)
	box, err := pop.NewMigrationBox(walkable, conn)
	if err != nil {
		return m, err
	}

	return PopMigrator{box}, nil
}

func (m PopMigrator) Up() error {
	err := m.box.Up()
	if err != nil {
		return err
	}

	err = m.box.Status()
	if err != nil {
		return err
	}

	return nil
}

func (m PopMigrator) Down() error {
	err := m.box.Down(1)
	if err != nil {
		return errors.WithStack(err)
	}

	err = m.box.Status()
	if err != nil {
		return err
	}

	return nil
}

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

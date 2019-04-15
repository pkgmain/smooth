package templates

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/shurcooL/httpfs/vfsutil"
)

type AppModel struct {
	App    string
	Module string
}

type DomainModel struct {
	Capitalized       string
	CapitalizedPlural string
	Singular          string
	Plural            string
	FirstLetter       string
}

func GenerateAppFiles(model AppModel) error {
	return vfsutil.Walk(VFS, "/app", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		content, err := vfsutil.ReadFile(VFS, path)
		if err != nil {
			return errors.WithStack(err)
		}

		path = strings.Replace(path, "/app", model.App, 1)
		path = strings.Replace(path, filepath.Ext(path), "", 1)
		return Write(string(content), path, model)
	})
}

func GenerateDomainFiles(model DomainModel) error {
	return vfsutil.Walk(VFS, "/domain", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		content, err := vfsutil.ReadFile(VFS, path)
		if err != nil {
			return errors.WithStack(err)
		}

		path = strings.Replace(path, "/domain", "", 1)
		path = strings.Replace(path, filepath.Ext(path), "", 1)
		path = strings.Replace(path, "domain", model.Singular, 1)
		return Write(string(content), filepath.Join(model.Singular, path), model)
	})
}

func GenerateMigrationFiles(model DomainModel) error {
	return vfsutil.Walk(VFS, "/migration", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		content, err := vfsutil.ReadFile(VFS, path)
		if err != nil {
			return errors.WithStack(err)
		}

		path = strings.Replace(path, "/migration/", "", 1)
		path = strings.Replace(path, filepath.Ext(path), "", 1)
		now := time.Now()
		timestamp := now.Format("20060102150405")
		path = fmt.Sprintf("%s_%s.%s", timestamp, model.Singular, path)
		path = filepath.Join("resources", "migrations", path)
		return Write(string(content), path, model)
	})
}

func Write(content, path string, model interface{}) error {
	log.Info().Msgf("writing: %s", path)

	tmpl, err := template.New(path).Parse(content)
	if err != nil {
		return errors.WithStack(err)
	}

	data := new(bytes.Buffer)
	err = tmpl.Execute(data, model)
	if err != nil {
		return errors.WithStack(err)
	}

	dir := filepath.Dir(path)
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return errors.WithStack(err)
	}

	err = ioutil.WriteFile(path, data.Bytes(), 0600)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

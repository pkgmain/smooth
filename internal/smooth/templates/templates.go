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

	"github.com/gobuffalo/packd"
	"github.com/gobuffalo/packr/v2"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

var appTemplates *packr.Box
var domainTemplates *packr.Box
var migrationTemplates *packr.Box

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

func init() {
	appTemplates = packr.New("appTemplates", "./app")
	domainTemplates = packr.New("domainTemplates", "./domain")
	migrationTemplates = packr.New("migrationTemplates", "./migration")
}

func GenerateAppFiles(model AppModel) error {
	return appTemplates.Walk(func(path string, file packd.File) error {
		path = strings.Replace(path, filepath.Ext(path), "", 1)
		return Write(file.String(), filepath.Join(model.App, path), model)
	})
}

func GenerateDomainFiles(model DomainModel) error {
	return domainTemplates.Walk(func(path string, file packd.File) error {
		path = strings.Replace(path, filepath.Ext(path), "", 1)
		path = strings.Replace(path, "domain", model.Singular, 1)
		return Write(file.String(), filepath.Join(model.Singular, path), model)
	})
}

func GenerateMigrationFiles(model DomainModel) error {
	return migrationTemplates.Walk(func(path string, file packd.File) error {
		path = strings.Replace(path, filepath.Ext(path), "", 1)
		now := time.Now()
		timestamp := now.Format("20060102150405")
		path = fmt.Sprintf("%s_%s.%s", timestamp, model.Singular, path)
		path = filepath.Join("migrations", path)
		return Write(file.String(), path, model)
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

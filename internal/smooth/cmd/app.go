package cmd

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/pkgmain/smooth/internal/smooth/templates"
)

var infoWriter = InfoWriter{}

var appCmd = &cobra.Command{
	Use:   "app [module name]",
	Short: "Scaffold a new application",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Info().Msg("Scaffolding application...")

		module := getModule(args)
		app := getApp(module)

		// write files and folders
		m := templates.AppModel{App: app, Module: module}
		err := writeFilesFromTemplates(m)
		if err != nil {
			return err
		}

		err = os.Mkdir(filepath.Join(app, "resources", "migrations"), os.ModePerm)
		if err != nil {
			return errors.WithStack(err)
		}

		// run go mod init
		log.Info().Msgf("running: go mod init %s", module)
		err = execCommand(app, "go", "mod", "init", module)
		if err != nil {
			return err
		}

		// run go get
		log.Info().Msg("running: go get -tags=dev")
		err = execCommand(app, "go", "get", "-tags=dev")
		if err != nil {
			return err
		}

		// run go generate
		log.Info().Msg("running: go generate")
		err = execCommand(app, "go", "generate")
		if err != nil {
			return err
		}

		absPath, err := filepath.Abs(app)
		if err != nil {
			return errors.WithStack(err)
		}

		log.Info().Msgf("Finished scaffolding! Your application is now available in %s", absPath)

		return nil
	},
}

func init() {
	newCmd.AddCommand(appCmd)
}

func writeFilesFromTemplates(m templates.AppModel) error {
	err := templates.GenerateAppFiles(m)

	if err != nil {
		return err
	}

	return nil
}

func getModule(args []string) string {
	return args[0]
}

func getApp(module string) string {
	tokens := strings.Split(module, "/")
	app := tokens[len(tokens)-1]

	return app
}

func execCommand(workingDir, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Dir = workingDir
	cmd.Stdout = infoWriter
	cmd.Stderr = infoWriter

	err := cmd.Run()
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

type InfoWriter struct{}

func (InfoWriter) Write(p []byte) (n int, err error) {
	log.Info().Msg(string(p))
	return len(p), nil
}

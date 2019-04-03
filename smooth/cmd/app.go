package cmd

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/pkgmain/smooth/smooth/templates"
)

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
		writeFilesFromTemplates(m)
		err := os.Mkdir(filepath.Join(app, "migrations"), os.ModePerm)
		if err != nil {
			return err
		}

		// run go mod init
		err = execCommand(app, "go", "mod", "init", module)
		if err != nil {
			return err
		}

		// run go get
		err = execCommand(app, "go", "get")
		if err != nil {
			return err
		}

		log.Info().Msg("Finished scaffolding application!")

		return nil
	},
}

func init() {
	newCmd.AddCommand(appCmd)
}

func writeFilesFromTemplates(m templates.AppModel) {
	err := templates.GenerateAppFiles(m)

	if err != nil {
		log.Fatal().Msgf("%+v", errors.WithStack(err))
	}
}

func getModule(args []string) string {
	m := args[0]
	log.Info().Msgf("Module: %s", m)

	return m
}

func getApp(module string) string {
	tokens := strings.Split(module, "/")
	app := tokens[len(tokens)-1]
	log.Info().Msgf("App: %s", app)

	return app
}

func execCommand(workingDir, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Dir = workingDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

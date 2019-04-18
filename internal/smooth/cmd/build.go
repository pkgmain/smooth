package cmd

import (
	"os"

	"github.com/gobuffalo/packr/v2/jam"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:                "build",
	Short:              "build a binary including static assets such as config files, migrations, and templates",
	DisableFlagParsing: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := jam.Pack(jam.PackOptions{})
		defer func() {
			err := jam.Clean()
			if err != nil {
				log.Error().Msg(errors.Cause(err).Error())
			}
		}()
		if err != nil {
			return err
		}

		wd, err := os.Getwd()
		if err != nil {
			return err
		}

		cmdArgs := append([]string{"build"}, args...)
		err = execCommand(wd, "go", cmdArgs...)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}

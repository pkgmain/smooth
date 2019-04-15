package cmd

import (
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var debug bool

var rootCmd = &cobra.Command{
	Use:   "smooth",
	Short: "A toolkit to create applications based on Chi and Pop",
}

func init() {
	log.Logger = log.Output(zerolog.NewConsoleWriter())

	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "debug logs")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {

		if debug {
			log.Error().Msgf("%+v", err)
			return
		}

		log.Error().Msg(errors.Cause(err).Error())
	}
}

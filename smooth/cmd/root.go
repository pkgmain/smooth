package cmd

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "smooth",
	Short: "A toolkit to create applications based on Chi and Pop",
}

func init() {
	log.Logger = log.Output(zerolog.NewConsoleWriter())
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Error().Msgf("%+v", err)
	}
}

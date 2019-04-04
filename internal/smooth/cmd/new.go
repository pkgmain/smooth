package cmd

import (
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:       "new [type]",
	Short:     "Scaffold a new app or domain",
	ValidArgs: []string{"app", "domain"},
	Args:      cobra.OnlyValidArgs,
}

func init() {
	rootCmd.AddCommand(newCmd)
}

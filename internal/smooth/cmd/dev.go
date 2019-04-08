package cmd

import (
	"github.com/markbates/refresh/refresh"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "hot reload web server on code change",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := &refresh.Configuration{}
		err := cfg.Load("refresh.yml")
		if err != nil {
			return errors.WithStack(err)
		}

		mgr := refresh.New(cfg)
		err = mgr.Start()
		if err != nil {
			return errors.WithStack(err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(devCmd)
}

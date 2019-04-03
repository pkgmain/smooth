package cmd

import (
	"io/ioutil"
	"strings"

	"github.com/markbates/inflect"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/pkgmain/smooth/smooth/templates"
)

var domainCmd = &cobra.Command{
	Use:   "domain [name]",
	Short: "Scaffold a new domain",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := ioutil.ReadFile("go.mod")
		if err != nil {
			return errors.Wrap(err, "command must be run from the root of a go module")
		}

		return createDomain(args[0])
	},
}

func createDomain(domain string) error {
	domain = strings.ToLower(domain)
	m := templates.DomainModel{
		Singular:    domain,
		Plural:      inflect.Pluralize(domain),
		Capitalized: strings.Title(domain),
		FirstLetter: domain[0:1],
	}
	err := templates.GenerateDomainFiles(m)
	if err != nil {
		log.Fatal().Msgf("%+v", err)
	}

	err = templates.GenerateMigrationFiles(m)
	if err != nil {
		log.Fatal().Msgf("%+v", err)
	}

	return nil
}

func init() {
	newCmd.AddCommand(domainCmd)
}

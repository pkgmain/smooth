package cmd

import (
	"io/ioutil"
	"strings"

	"github.com/markbates/inflect"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/pkgmain/smooth/internal/smooth/templates"
)

var domainCmd = &cobra.Command{
	Use:   "domain [name]",
	Short: "Scaffold a new domain. This includes code for the domain, migrations, CRUD http handlers and repos.",
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
	log.Info().Msg("Scaffolding new domain...")

	domain = strings.ToLower(domain)
	plural := inflect.Pluralize(domain)
	m := templates.DomainModel{
		Singular:          domain,
		Plural:            plural,
		Capitalized:       strings.Title(domain),
		CapitalizedPlural: strings.Title(plural),
		FirstLetter:       domain[0:1],
	}
	err := templates.GenerateDomainFiles(m)
	if err != nil {
		return err
	}

	err = templates.GenerateMigrationFiles(m)
	if err != nil {
		return err
	}

	// run go generate
	log.Info().Msg("running: go generate")
	err = execCommand(".", "go", "generate")
	if err != nil {
		return err
	}

	log.Info().Msg("Finished scaffolding domain!")
	return nil
}

func init() {
	newCmd.AddCommand(domainCmd)
}

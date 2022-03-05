package utils

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/endorama/devid/internal/persona"
)

var errNoActivePersona = errors.New("no active persona found")
var errNoActivePersonaAndNoFlag = errors.New("no active persona found and no --persona flag specified")

func LoadPersona(cmd *cobra.Command) (persona.Persona, error) {
	currentPersona := viper.GetString("active_persona")

	var err error
	if currentPersona == "" && cmd.Flag("persona") != nil {
		currentPersona, err = cmd.Flags().GetString("persona")

		if err != nil {
			return persona.Persona{}, fmt.Errorf("cannot access flag current persona: %w", err)
		}
	}

	if currentPersona == "" {
		if cmd.Flag("persona") == nil {
			return persona.Persona{}, errNoActivePersona
		}

		return persona.Persona{}, errNoActivePersonaAndNoFlag
	}

	p, err := persona.Load(currentPersona)
	if err != nil {
		return persona.Persona{}, fmt.Errorf("cannot load persona: %w", err)
	}

	return p, nil
}

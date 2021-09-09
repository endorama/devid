package ssh

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/endorama/devid/cmd/ui"
	"github.com/endorama/devid/cmd/utils"
	"github.com/spf13/cobra"
)

var createKeyCmd = &cobra.Command{ //nolint:gochecknoglobals // required by cobra
	Use:   "create-key",
	Short: "create a SSH key",
	Long: `Perform secure SSH key generation.
`,
	Run: func(cmd *cobra.Command, args []string) {
		p, err := utils.LoadPersona(cmd)
		if err != nil {
			ui.Fatal(err, 1)
		}

		passphrase := utils.GeneratePassphrase()
		ui.Info(fmt.Sprintf("Generated passphrase: %s", passphrase))

		genCmd := "ssh-keygen"
		genArgs := []string{
			"-t", "ed25519",
			"-a", "100",
			"-N", passphrase,
			"-C", p.Config.GetString("identity.email"),
			"-f", path.Join(p.Location(), pluginName, "id_rsa"),
		}

		c := exec.Command(genCmd, genArgs...)

		out, err := c.Output()
		if err != nil {
			ui.Error(err)
			if e, ok := err.(*exec.ExitError); ok {
				ui.Error(errors.New(string(e.Stderr)))
			}
			os.Exit(1)
		}
		ui.Output(string(out))
	},
}

func (p Plugin) Commands() []*cobra.Command {
	return []*cobra.Command{createKeyCmd}
}

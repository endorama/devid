package cmds

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/spf13/cobra"

	"github.com/endorama/devid/cmd/ui"
	"github.com/endorama/devid/cmd/utils"
)

var CreateKey = &cobra.Command{ //nolint:gochecknoglobals // required by cobra
	Use:   "create-key",
	Short: "create a SSH key",
	Long: `Perform secure SSH key generation.
`,
	Run: func(cmd *cobra.Command, _ []string) {
		p, err := utils.LoadPersona(cmd)
		if err != nil {
			ui.Fatal(err, 1)
		}

		passphrase := utils.GeneratePassphrase()
		ui.Infof(fmt.Sprintf("Generated passphrase: %s", passphrase))

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

			var eerr *exec.ExitError
			if errors.As(err, &eerr) {
				ui.Error(err)
			}
			os.Exit(1)
		}

		ui.Outputf(string(out))
	},
}

package cmds

import (
	"os"
	"path"

	"github.com/spf13/cobra"

	"github.com/endorama/devid/cmd/ui"
	"github.com/endorama/devid/cmd/utils"
)

var PrintPubKey = &cobra.Command{ //nolint:gochecknoglobals // required by cobra
	Use:   "print-pubkey",
	Short: "Print public key",
	Long: `Print SSH public key to be reused or copied outside of devid.
`,
	Run: func(cmd *cobra.Command, _ []string) {
		p, err := utils.LoadPersona(cmd)
		if err != nil {
			ui.Fatal(err, 1)
		}

		pubKeyLocatin := path.Join(p.Location(), pluginName, "id_rsa.pub")
		dat, err := os.ReadFile(pubKeyLocatin)
		if err != nil {
			ui.Error(err)
		}

		ui.Outputf(string(dat))
	},
}

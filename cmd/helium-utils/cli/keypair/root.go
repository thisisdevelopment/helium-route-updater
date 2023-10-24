package keypair

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "keypair",
	Short: "",
}

func init() {
	//RootCmd.AddCommand(check.RootCmd)
}

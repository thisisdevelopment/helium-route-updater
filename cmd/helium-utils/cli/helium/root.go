package helium

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "helium",
	Short: "",
}

func init() {
	//RootCmd.AddCommand(check.RootCmd)
}

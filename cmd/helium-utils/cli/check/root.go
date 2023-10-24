package check

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "check",
	Short: "",
}

func init() {
	//RootCmd.AddCommand(check.RootCmd)
}

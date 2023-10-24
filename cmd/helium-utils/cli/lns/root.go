package lns

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "lns",
	Short: "",
}

func init() {
	//RootCmd.AddCommand(check.RootCmd)
}

package keypair

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func NewAllCmd(commands... *cobra.Command) *cobra.Command {
	return &cobra.Command{
		Use: "all",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			for _, command := range commands {
				if command == cmd {
					continue
				}
				fmt.Println(command.Use)
				fmt.Println(strings.Repeat("-", len(command.Use)))
				cmd.Root().SetArgs([]string{cmd.Parent().Use, command.Use})
				cmd.Root().Execute()
				fmt.Println()
			}

			return nil
		},
	}
}

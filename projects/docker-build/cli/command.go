package cli

import "github.com/spf13/cobra"

type Command interface {
	Execute(cmd *cobra.Command, args []string)
	Initialize(cmd *cobra.Command)
}

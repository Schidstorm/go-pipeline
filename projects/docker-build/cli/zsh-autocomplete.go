package cli

import (
	"os"

	"github.com/spf13/cobra"
)

//Backup is the interface between the lib.Backup and cobra
type ZshAutocomplete struct {
}

//Execute runs lib.Backup
func (b ZshAutocomplete) Execute(cmd *cobra.Command, args []string) {
	cmd.Root().GenZshCompletion(os.Stdout)
}

//Initialize sets up die cobra command arguments
func (b ZshAutocomplete) Initialize(cmd *cobra.Command) {
	cmd.Use = "zsh_autocomplete"
}

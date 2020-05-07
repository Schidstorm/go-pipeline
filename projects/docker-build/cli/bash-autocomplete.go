package cli

import (
	"os"

	"github.com/spf13/cobra"
)

//Backup is the interface between the lib.Backup and cobra
type BashAutocomplete struct {
}

//Execute runs lib.Backup
func (b BashAutocomplete) Execute(cmd *cobra.Command, args []string) {
	cmd.Root().GenBashCompletion(os.Stdout)
}

//Initialize sets up die cobra command arguments
func (b BashAutocomplete) Initialize(cmd *cobra.Command) {
	cmd.Use = "bash_autocomplete"
}

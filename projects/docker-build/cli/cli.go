package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var commands []Command = []Command{
	&Count{},
	&BashAutocomplete{},
	&ZshAutocomplete{},
}

var rootCmd = &cobra.Command{
	Short: "kurz",
	Long:  `lang`,
}

func Execute() {
	for _, cmd := range commands {
		var subCommand = &cobra.Command{
			Run: cmd.Execute,
		}
		cmd.Initialize(subCommand)
		rootCmd.AddCommand(subCommand)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

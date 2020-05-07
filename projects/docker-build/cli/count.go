package cli

import (
	"strconv"

	"github.com/schidstorm/imap-backup/lib"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type Count struct {
}

func (b Count) Execute(cmd *cobra.Command, args []string) {
	port, err := strconv.ParseUint(cmd.PersistentFlags().Lookup("port").Value.String(), 10, 16)

	if err != nil {
		log.Fatalln(err)
	}

	backup := &lib.Backup{lib.BackupArguments{
		User:      cmd.PersistentFlags().Lookup("user").Value.String(),
		Password:  cmd.PersistentFlags().Lookup("password").Value.String(),
		Host:      cmd.PersistentFlags().Lookup("host").Value.String(),
		Port:      uint16(port),
		ExportDir: "mock",
	}}

	err = backup.DoCount()
	if err != nil {
		log.Fatalln(err)
	}
}

//Initialize sets up die cobra command arguments
func (b Count) Initialize(cmd *cobra.Command) {
	cmd.Use = "count"
	cmd.PersistentFlags().String("user", "", "The username (email) of your imap account")
	cmd.PersistentFlags().String("password", "", "The password of your imap account")
	cmd.PersistentFlags().String("host", "", "The hostname of the imap server")
	cmd.PersistentFlags().Uint16("port", 0, "The port of the imap server")

}

package cmd

import (
	hermes "github.com/f-ewald/hermes/pkg"
	"github.com/spf13/cobra"
	"log"
)

// checkCommand represents the conversation command
var checkCommand = &cobra.Command{
	Use:   "check",
	Short: "Validate the environment",
	Long:  `Validate that the environment works and that the iMessage database can be accessed.`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := hermes.MessageDBFilename()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Trying to access iMessage database at: \"%s\"\n", db)

		readAccess := hermes.HasReadAccess(db)
		if readAccess {
			log.Println("Success! Read access to iMessage database exists.")
		} else {
			log.Printf(`Cannot open the database at "%s"`, db)
			log.Println(`This is likely due to no read permissions for the terminal.
There are two options to solve this.

1. Give your terminal full disk access in Settings > Security > Full Disk Access
2. (Recommended) Copy your iMessage database into your user folder and use the --db flag when starting hermes.`)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(checkCommand)
}

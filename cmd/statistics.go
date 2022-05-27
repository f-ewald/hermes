package cmd

import (
	"context"
	hermes "github.com/f-ewald/hermes/pkg"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(statisticsCmd)
}

var statisticsCmd = &cobra.Command{
	Use:     "statistics",
	Aliases: []string{"stats"},
	Short:   "Display message statistics",
	Long: `Analyze all messages that have been sent and received and display the statistics.
Supports text, JSON and YAML output format.`,
	Run: func(cmd *cobra.Command, args []string) {
		messageDB, err := hermes.MessageDBFilename()
		if err != nil {
			log.Fatal(err)
		}
		db := hermes.NewDatabase(messageDB)
		err = db.Connect()
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			err = db.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()

		stats, err := db.Statistics(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
		formatted, err := cfg.Formatter.Format(stats, "statistics.tpl")
		if err != nil {
			log.Fatal(err)
		}
		var printer hermes.Printer
		printer = &hermes.StdoutPrinter{}
		printer.Print(formatted)
	},
}

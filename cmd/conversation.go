package cmd

import (
	"context"
	hermes "github.com/f-ewald/hermes/pkg"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(conversationCmd)
	conversationCmd.AddCommand(conversationListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// conversationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// conversationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// conversationCmd represents the conversation command.
var conversationCmd = &cobra.Command{
	Use:   "conversation",
	Short: "Show conversations, find participants",
	//Long:  `TODO`,
	//Run: func(cmd *cobra.Command, args []string) {
	//	fmt.Println("conversation called")
	//},
}

var conversationListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all conversations",
	Long:  "List all conversations with ",
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

		conversations, err := db.ListConversations(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		var printer hermes.Printer
		printer = &hermes.StdoutPrinter{}
		for _, conversation := range conversations {
			printer.Print([]byte(conversation))
		}
	},
}

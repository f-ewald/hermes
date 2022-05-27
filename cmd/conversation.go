package cmd

import (
	"context"
	hermes "github.com/f-ewald/hermes/pkg"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

func init() {
	rootCmd.AddCommand(conversationCmd)
	conversationCmd.AddCommand(conversationGetCmd)
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
	Use:     "conversation",
	Short:   "Show retrieveConversations, find participants",
	Aliases: []string{"retrieveConversations"},
	//Long:  `TODO`,
	//Run: func(cmd *cobra.Command, args []string) {
	//	fmt.Println("conversation called")
	//},
}

func retrieveConversations() []*hermes.Chat {
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

	handles, err := db.ListConversations(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return handles
}

var conversationListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all retrieveConversations",
	Long:  "List all retrieveConversations",
	Run: func(cmd *cobra.Command, args []string) {
		chats := retrieveConversations()

		output, err := cfg.Formatter.Format(chats, "conversation-list.tpl")
		if err != nil {
			panic(err)
		}
		var printer hermes.Printer
		printer = &hermes.StdoutPrinter{}
		printer.Print(output)
	},
}

var conversationGetCmd = &cobra.Command{
	Use:       "get NUMBER ... ",
	Short:     "Get conversation",
	Long:      "Get conversation",
	Args:      cobra.MinimumNArgs(1),
	ValidArgs: []string{"NUMBER"},
	Run: func(cmd *cobra.Command, args []string) {
		var printer hermes.Printer
		printer = &hermes.StdoutPrinter{}

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

		for _, arg := range args {
			chatID, err := strconv.Atoi(arg)
			if err != nil {
				panic(err)
			}
			conversations, err := db.Conversation(context.Background(), chatID)
			if err != nil {
				panic(err)
			}
			b, err := cfg.Formatter.Format(conversations, "conversation.tpl")
			if err != nil {
				panic(err)
			}
			printer.Print(b)
		}
	},
}

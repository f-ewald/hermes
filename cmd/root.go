package cmd

import (
	"fmt"
	hermes "github.com/f-ewald/hermes/pkg"
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

var output string

// Config bundles the configuration.
type Config struct {
	// Output defines the output format. If nothing is defined, normal text will be written to
	// STDOUT.
	Output string

	Formatter hermes.Formatter

	// ChatDB contains the full path to the chat database.
	ChatDB string
}

// cfg contains the configuration.
var cfg Config

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hermes",
	Short: "Command-line interface for iMessage databases",
	Long: `Hermes is a command-line interface for iMessage databases.
You can use it to analyze and display retrieveConversations and view statistics.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		switch output {
		case "text":
			cfg.Formatter = &hermes.TextFormatter{}
		case "json":
			cfg.Formatter = &hermes.JsonFormatter{}
		case "yaml":
			cfg.Formatter = &hermes.YamlFormatter{}
		default:
			panic("output must be one of json, text or yaml")
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.hermes.yaml)")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "text", "The output format. Can be either json, yaml or text")
	rootCmd.PersistentFlags().StringVarP(&cfg.ChatDB, "database", "d", "", "Full path to the chat database if it is different than the default path.")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".hermes" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".hermes")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		_, _ = fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

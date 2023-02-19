package cmd

import (
	"github.com/event_bright/internal/config"
	"github.com/spf13/cobra"
	"log"
)

var (
	// cfgFile store the configuration file name
	//cfgFile                 string
	//verbose, prettyPrintLog bool
	rootCmd = &cobra.Command{
		Use:   "event_bright",
		Short: "Event Bright Backend Server",
		Long:  `Event Bright Backend Server`,
	}
)

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	log.Println("Loading configurations")
	config.Init()
	log.Println("Configurations loaded successfully!")
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

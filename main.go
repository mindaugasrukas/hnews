package main

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hnews",
	Short: "Download HackerNews Post as JSON",
}

func main() {
	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	// Configure default logger
	// TODO: make this configurable
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
}

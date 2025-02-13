package main

import (
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"g"},
	Short:   "",
	Run: func(cmd *cobra.Command, args []string) {
		id := cmd.Flag("id").Value.String()
		if id == "" {
			cmd.Help()
			return
		}

	},
}

func init() {
	getCmd.Flags().StringP("id", "i", "", "HackerNews Post ID")
	rootCmd.AddCommand(getCmd)
}

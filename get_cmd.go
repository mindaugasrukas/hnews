package main

import (
	"encoding/json"
	"fmt"
	"hnews/hn"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"g"},
	Short:   "",
	Run: func(cmd *cobra.Command, args []string) {
		sid := cmd.Flag("id").Value.String()
		if sid == "" {
			cmd.Help()
			return
		}

		id, err := strconv.Atoi(sid)
		if err != nil {
			log.Fatal(err, "Invalid ID: ", sid)
		}

		api := hn.NewAPI()
		item, err := api.GetPost(id)
		if err != nil {
			log.Fatal(err)
		}

		// print the post as JSON
		data, err := json.MarshalIndent(item, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(data))
	},
}

func init() {
	getCmd.Flags().StringP("id", "i", "", "HackerNews Post ID")
	rootCmd.AddCommand(getCmd)
}

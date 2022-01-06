/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	controller2 "github.com/LuVlk/dadjoke/controller"
	"github.com/spf13/cobra"
)

var (
	// used for flags
	limit int

	// searchCmd represents the search command
	searchCmd = &cobra.Command{
		Use:   "search",
		Short: "Search for your favorite dad joke",
		Long:  `Use a term to search for a specific dad joke`,
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				args = append(args, "")
			}

			controller := controller2.MakeJokeController()
			jokes, err := controller.GetJokesByTerm(args[0], limit)
			if err != nil {
				fmt.Printf("could not get jokes - %s", err.Error())
			}

			for _, joke := range jokes {
				fmt.Println(joke.Joke)
				fmt.Println("----")
			}
		},
	}
)

func init() {
	searchCmd.PersistentFlags().IntVarP(&limit, "limit", "l", 10, "the maximum number of jokes to return (max. 30)")
	rootCmd.AddCommand(searchCmd)
}

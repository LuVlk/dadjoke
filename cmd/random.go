/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/LuVlk/dadjoke/controller"
	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get a random dad joke",
	Long:  `Get a random dad joke from https://icanhazdadjoke.com/`,
	Run: func(cmd *cobra.Command, args []string) {
		controller := controller.MakeJokeController()
		joke, err := controller.GetRandomJoke()
		if err != nil {
			fmt.Printf("could not get random joke - %s", err.Error())
		}

		fmt.Println(joke.Joke)
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
}

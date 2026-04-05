/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get a random dad joke",
	Long:  `This command fetches a random dad joke from the icanhazdadjoke API and displays it in the terminal.`,
	Run: func(cmd *cobra.Command, args []string) {
		getRandomJoke()
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
}

type Joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func getRandomJoke() {
	url:= "https://icanhazdadjoke.com/"

	responseBytes := getJokeData(url)
	joke:= Joke{}

	if err := json.Unmarshal(responseBytes, &joke); err != nil {
		fmt.Printf("Could not parse joke data - %v\n", err)
		return
	}

	fmt.Println(joke.Joke) // Access the joke text using joke.Joke after unmarshal

	
}

func getJokeData(baseAPI string) []byte {
	request, err := http.NewRequest(
		http.MethodGet,
		baseAPI,
		nil,
	)
	if err != nil {
		fmt.Printf("Could not request a dadjoke - %v\n", err)
		return nil
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "Dadjoke CLI (https://github.com/Afra997/dadjoke)")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Printf("Could not get a dadjoke - %v\n", err)
		return nil
	}

	defer response.Body.Close()

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Could not read response body - %v\n", err)
		return nil
	}

	return responseBytes
}

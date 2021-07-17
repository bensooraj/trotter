package repl

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bensooraj/trotter/graph"
	"github.com/c-bata/go-prompt"
)

func autocomplete(cityMap *graph.CityMap) func(prompt.Document) []prompt.Suggest {

	s := []prompt.Suggest{}
	for _, c := range *cityMap {
		s = append(s, prompt.Suggest{Text: c.ID, Description: c.Name})
	}

	return func(d prompt.Document) []prompt.Suggest {
		return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
	}
}

func Run(cityMap *graph.CityMap) {

	// Blocking main and waiting for shutdown of the daemon.
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)

	ac := autocomplete(cityMap)
	for {
		fmt.Println("Please type the city you want to begin globe-trotting from? (type `exit` if you want to stop the REPL) ")
		input := prompt.Input("> ", ac)

		switch input {
		case "exit":
			fmt.Println("Exiting...")
			os.Stdout.Close()
			return
		default:
			city, ok := (*cityMap)[strings.ToUpper(input)]
			if !ok {
				fmt.Printf("%s not found\n", input)
			} else {
				fmt.Println("Computing the shortest path if you were start and end at ")
				outMsg := graph.FindSingleSourceShortestPath(city)
				fmt.Println(outMsg)
			}
		}
	}
}

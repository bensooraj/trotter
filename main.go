package main

import (
	"time"

	"github.com/bensooraj/trotter/graph"
	"github.com/bensooraj/trotter/repl"
	"github.com/bensooraj/trotter/util"
)

func main() {
	cityMap := graph.LoadContinentToCitiesMap()
	cityMap.PopulateDistanceDataLinear()

	util.LogRuntimeMemStats()

	repl.Run(cityMap)

	time.Sleep(1 * time.Second)
}

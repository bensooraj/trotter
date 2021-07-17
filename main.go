package main

import (
	"fmt"
	"time"

	"github.com/bensooraj/trotter/graph"
	"github.com/bensooraj/trotter/util"
)

func main() {
	util.LogRuntimeMemStats()

	cts := graph.LoadContinentToCitiesMap()
	fmt.Println(len(*cts))

	util.LogRuntimeMemStats()

	cts.PopulateDistanceDataLinear()

	util.LogRuntimeMemStats()
	time.Sleep(5 * time.Second)
}

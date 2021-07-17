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

	bom, ok := (*cts)["BOM"]
	if !ok {
		fmt.Println("BOM not found")
	} else {
		fmt.Println("City: ", bom.Name)
		fmt.Println("Africa: ", bom.AfricaD.City.Name, bom.AfricaD.Distance)
		fmt.Println("Asia: ", bom.AsiaD.City.Name, bom.AsiaD.Distance)
		fmt.Println("Europe: ", bom.EuropeD.City.Name, bom.EuropeD.Distance)
		fmt.Println("NorthAmerica: ", bom.NorthAmericaD.City.Name, bom.NorthAmericaD.Distance)
		fmt.Println("Oceania: ", bom.OceaniaD.City.Name, bom.OceaniaD.Distance)
		fmt.Println("SouthAmerica: ", bom.SouthAmericaD.City.Name, bom.SouthAmericaD.Distance)
	}

	time.Sleep(5 * time.Second)
}

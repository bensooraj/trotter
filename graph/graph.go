package graph

import (
	"fmt"

	"github.com/RyanCarrier/dijkstra"
	"github.com/bensooraj/trotter/util"
)

func FindSingleSourceShortestPath(originCity *City) {

	// https://pkg.go.dev/gonum.org/v1/gonum/graph/path#DijkstraFrom

	d := dijkstra.NewGraph()

	cityStops := []*City{originCity}
	d.AddVertex(0)

	for i, city := range []*City{
		originCity.AfricaD.City, originCity.AsiaD.City,
		originCity.EuropeD.City, originCity.NorthAmericaD.City,
		originCity.OceaniaD.City, originCity.SouthAmericaD.City,
	} {
		if city.ContID != originCity.ContID {
			cityStops = append(cityStops, city)
			// Add the cities as vertices
			d.AddVertex(i + 1)
		}
	}

	// Create edges and calculate the weights between them
	for i, sourceCity := range cityStops {
		for j, destinationCity := range cityStops {
			if sourceCity.ID != destinationCity.ID {
				d.AddArc(i, j, int64(util.GetDistanceFromLatLonInKm(
					sourceCity.Lat, sourceCity.Lon,
					destinationCity.Lat, destinationCity.Lon,
				)))
			}
		}
	}

	d.ExportToFile("d.txt")

	// Find the shortest path starting from the origin back to the origin
	best, err := d.Shortest(0, 5)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}

	fmt.Println(best.Path, best.Distance)
}
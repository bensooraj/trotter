package graph

import (
	"fmt"
	"math"

	"github.com/bensooraj/trotter/util"
)

func FindSingleSourceShortestPath(originCity *City) {

	cityStops := []*City{originCity}

	for _, city := range []*City{
		originCity.AfricaD.City, originCity.AsiaD.City,
		originCity.EuropeD.City, originCity.NorthAmericaD.City,
		originCity.OceaniaD.City, originCity.SouthAmericaD.City,
	} {
		if city.ContID != originCity.ContID {
			cityStops = append(cityStops, city)
		}
	}

	permutations := util.GetPermutations(6, 6)
	fmt.Println("permutations: ", len(permutations))

	// Since I am going the brute force way, I might as well save on a few
	// distance calculations
	distCache := make(map[string]float64, 0)

	leastDistance := math.MaxFloat64
	bestPath := []int{}

	for i := 0; i < len(permutations); i++ {
		p := permutations[i]
		var totalDistance float64 = 0

		for j := 0; j < 6; j++ {
			s, d := cityStops[p[j]], cityStops[p[(j+1)%5]]

			var distance float64
			comboID := fmt.Sprintf("%s-%s", s.ID, d.ID)

			if v, found := distCache[comboID]; found {
				distance = v
			} else {
				distance = util.GetDistanceFromLatLonInKm(s.Lat, s.Lon, d.Lat, d.Lon)
				distCache[comboID] = distance
			}

			totalDistance += distance
		}

		if totalDistance < leastDistance {
			leastDistance = totalDistance
			bestPath = p
		}
	}

	fmt.Println("bestPath: ", bestPath, leastDistance, "that is, ")
	for _, seq := range bestPath {
		fmt.Printf("[%s] => ", cityStops[seq].Name)
	}
	fmt.Println()
	fmt.Println()
}

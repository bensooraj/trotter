package graph

import "github.com/bensooraj/trotter/store"

func LoadContinentToCitiesMap() *CityMap {
	var cityMap CityMap = make(CityMap, 6)

	cities, err := store.GetCityListFromJSON()
	if err != nil {
		panic(err)
	}

	// Form the initial cities map
	for _, city := range *cities {
		cityMap[city.ID] = &City{
			ID:     city.ID,
			Name:   city.Name,
			ContID: string(city.ContID),

			Lat: city.Location.Lat,
			Lon: city.Location.Lon,

			AfricaD:       make(map[*City]float64),
			AsiaD:         make(map[*City]float64),
			EuropeD:       make(map[*City]float64),
			NorthAmericaD: make(map[*City]float64),
			OceaniaD:      make(map[*City]float64),
			SouthAmericaD: make(map[*City]float64),
		}

	}

	//

	return &cityMap
}

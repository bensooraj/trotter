package graph

import (
	"github.com/bensooraj/trotter/store"
)

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

			AfricaD:       nil,
			AsiaD:         nil,
			EuropeD:       nil,
			NorthAmericaD: nil,
			OceaniaD:      nil,
			SouthAmericaD: nil,
		}

	}

	//

	return &cityMap
}

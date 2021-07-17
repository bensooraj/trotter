package graph

import (
	"github.com/bensooraj/trotter/util"
)

const (
	Africa       = "africa"
	Asia         = "asia"
	Europe       = "europe"
	NorthAmerica = "north-america"
	Oceania      = "oceania"
	SouthAmerica = "south-america"
)

type Nearest struct {
	City     *City
	Distance float64
}

type City struct {
	ID     string
	Name   string
	ContID string

	Lat float64
	Lon float64

	// Distance maps for each continent
	AfricaD       *Nearest
	AsiaD         *Nearest
	EuropeD       *Nearest
	NorthAmericaD *Nearest
	OceaniaD      *Nearest
	SouthAmericaD *Nearest
}

type CityMap map[string]*City

func (cityMap *CityMap) PopulateDistanceDataLinear() {
	for _, sourceCity := range *cityMap {
		for _, destinationCity := range *cityMap {
			if sourceCity.ContID != destinationCity.ContID {

				distance := util.GetDistanceFromLatLonInKm(
					sourceCity.Lat, sourceCity.Lon,
					destinationCity.Lat, destinationCity.Lon,
				)

				switch destinationCity.ContID {
				case Africa:
					if sourceCity.AfricaD == nil {
						sourceCity.AfricaD = &Nearest{destinationCity, distance}
					} else if distance < sourceCity.AfricaD.Distance {
						sourceCity.AfricaD.City = destinationCity
						sourceCity.AfricaD.Distance = distance
					}
				case Asia:
					if sourceCity.AsiaD == nil {
						sourceCity.AsiaD = &Nearest{destinationCity, distance}
					} else if distance < sourceCity.AsiaD.Distance {
						sourceCity.AsiaD.City = destinationCity
						sourceCity.AsiaD.Distance = distance
					}
				case Europe:
					if sourceCity.EuropeD == nil {
						sourceCity.EuropeD = &Nearest{destinationCity, distance}
					} else if distance < sourceCity.EuropeD.Distance {
						sourceCity.EuropeD.City = destinationCity
						sourceCity.EuropeD.Distance = distance
					}
				case NorthAmerica:
					if sourceCity.NorthAmericaD == nil {
						sourceCity.NorthAmericaD = &Nearest{destinationCity, distance}
					} else if distance < sourceCity.NorthAmericaD.Distance {
						sourceCity.NorthAmericaD.City = destinationCity
						sourceCity.NorthAmericaD.Distance = distance
					}
				case Oceania:
					if sourceCity.OceaniaD == nil {
						sourceCity.OceaniaD = &Nearest{destinationCity, distance}
					} else if distance < sourceCity.OceaniaD.Distance {
						sourceCity.OceaniaD.City = destinationCity
						sourceCity.OceaniaD.Distance = distance
					}
				case SouthAmerica:
					if sourceCity.SouthAmericaD == nil {
						sourceCity.SouthAmericaD = &Nearest{destinationCity, distance}
					} else if distance < sourceCity.SouthAmericaD.Distance {
						sourceCity.SouthAmericaD.City = destinationCity
						sourceCity.SouthAmericaD.Distance = distance
					}
				}
			}
		}
	}
}

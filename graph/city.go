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

type City struct {
	ID     string
	Name   string
	ContID string

	Lat float64
	Lon float64

	// Distance maps for each continent
	AfricaD       map[*City]float64
	AsiaD         map[*City]float64
	EuropeD       map[*City]float64
	NorthAmericaD map[*City]float64
	OceaniaD      map[*City]float64
	SouthAmericaD map[*City]float64
}

type CityMap map[string]*City

func (cityMap *CityMap) PopulateDistanceDataLinear() {
	for _, sourceCity := range *cityMap {
		for _, destinationCity := range *cityMap {
			if sourceCity.ContID != destinationCity.ContID {
				switch destinationCity.ContID {
				case Africa:
					sourceCity.AfricaD[destinationCity] = util.GetDistanceFromLatLonInKm(
						sourceCity.Lat, sourceCity.Lon,
						destinationCity.Lat, destinationCity.Lon,
					)
				case Asia:
					sourceCity.AsiaD[destinationCity] = util.GetDistanceFromLatLonInKm(
						sourceCity.Lat, sourceCity.Lon,
						destinationCity.Lat, destinationCity.Lon,
					)
				case Europe:
					sourceCity.EuropeD[destinationCity] = util.GetDistanceFromLatLonInKm(
						sourceCity.Lat, sourceCity.Lon,
						destinationCity.Lat, destinationCity.Lon,
					)
				case NorthAmerica:
					sourceCity.NorthAmericaD[destinationCity] = util.GetDistanceFromLatLonInKm(
						sourceCity.Lat, sourceCity.Lon,
						destinationCity.Lat, destinationCity.Lon,
					)
				case Oceania:
					sourceCity.OceaniaD[destinationCity] = util.GetDistanceFromLatLonInKm(
						sourceCity.Lat, sourceCity.Lon,
						destinationCity.Lat, destinationCity.Lon,
					)
				case SouthAmerica:
					sourceCity.SouthAmericaD[destinationCity] = util.GetDistanceFromLatLonInKm(
						sourceCity.Lat, sourceCity.Lon,
						destinationCity.Lat, destinationCity.Lon,
					)
				}
			}
		}
	}
}

package store

type Cities map[string]City

type City struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Location    Location `json:"location"`
	CountryName string   `json:"countryName"`
	Iata        string   `json:"iata"`
	Rank        int64    `json:"rank"`
	CountryID   string   `json:"countryId"`
	Dest        string   `json:"dest"`
	Airports    []string `json:"airports"`
	Images      []string `json:"images"`
	Popularity  float64  `json:"popularity"`
	RegID       RegID    `json:"regId"`
	ContID      ContID   `json:"contId"`
	SubID       *string  `json:"subId"`
	TerID       *string  `json:"terId"`
	Con         int64    `json:"con"`
}

type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type ContID string

const (
	Africa       ContID = "africa"
	Asia         ContID = "asia"
	Europe       ContID = "europe"
	NorthAmerica ContID = "north-america"
	Oceania      ContID = "oceania"
	SouthAmerica ContID = "south-america"
)

type RegID string

const (
	Australasia      RegID = "australasia"
	Caribbean        RegID = "caribbean"
	CentralAfrica    RegID = "central-africa"
	CentralAmerica   RegID = "central-america"
	CentralAsia      RegID = "central-asia"
	CentralEurope    RegID = "central-europe"
	EasternAfrica    RegID = "eastern-africa"
	EasternAsia      RegID = "eastern-asia"
	EasternEurope    RegID = "eastern-europe"
	Melanesia        RegID = "melanesia"
	Micronesia       RegID = "micronesia"
	NorthernAfrica   RegID = "northern-africa"
	NorthernAmerica  RegID = "northern-america"
	NorthernEurope   RegID = "northern-europe"
	Polynesia        RegID = "polynesia"
	SouthEasternAsia RegID = "south-eastern-asia"
	SouthernAfrica   RegID = "southern-africa"
	SouthernAmerica  RegID = "southern-america"
	SouthernAsia     RegID = "southern-asia"
	SouthernEurope   RegID = "southern-europe"
	WesternAfrica    RegID = "western-africa"
	WesternAsia      RegID = "western-asia"
	WesternEurope    RegID = "western-europe"
)

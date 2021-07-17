package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func GetCityListFromJSON() (*Cities, error) {
	jsonFile, err := os.Open("store/dump/cities.json")
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("ErrorReadingJSONFile")
	}
	defer jsonFile.Close()

	var cities Cities

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &cities)

	return &cities, nil
}

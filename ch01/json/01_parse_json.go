package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const citiBikeURL = "https://gbfs.citibikenyc.com/gbfs/en/station_status.json"

type stationData struct {
	LastUpdated int `json:"lat_updated"`
	TTL         int `json:"ttl"`
	Data        struct {
		Stations []station `json:"stations"`
	} `json:"data"`
}

type station struct {
	ID                string `json:"station_id"`
	NumBikesAvailable int    `json:"num_bikes_available"`
	NumBikesDisabled  int    `json:"num_bikes_disabled"`
	NumDocksAvailable int    `json:"num_docks_available"`
	NumDocksDisabled  int    `json:"num_docks_disabled"`
	IsInstalled       int    `json:"is_installed"`
	IsRenting         int    `json:"is_renting"`
	IsReturning       int    `json:"is_returning"`
	LastReported      int    `json:"last_reported"`
	HasVaailableKeys  bool   `json:"eightd_has_available_keys"`
}

func main() {
	response, err := http.Get(citiBikeURL)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// resp의 body를 []byte로 읽음
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var sd stationData

	if err := json.Unmarshal(body, &sd); err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("%+v\n\n", sd.Data.Stations[0])
}

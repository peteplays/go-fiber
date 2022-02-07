package services

import (
	"encoding/json"
	"go-fiber-api/utils"
	"log"
)

type personResponse struct {
	Name      string   `json:"name"`
	Homeworld string   `json:"homeWorld"`
	Vehicles  []string `json:"vehicles,omitempty"`
	Starships []string `json:"starships,omitempty"`
}

type peopleResponse struct {
	Count    int              `json:"count"`
	Next     string           `json:"next"`
	Previous string           `json:"previous"`
	Results  []personResponse `json:"results"`
}

type personResult struct {
	Name      string             `json:"name"`
	Homeworld homeWordResponse   `json:"homeWorld"`
	Vehicles  []vehicleResponse  `json:"vehicles,omitempty"`
	Starships []starShipResponse `json:"starships,omitempty"`
}

type homeWordResponse struct {
	Name       string `json:"name"`
	Climate    string `json:"climate"`
	Terrain    string `json:"terrain"`
	Population string `json:"population"`
}

type vehicleResponse struct {
	Name                 string `json:"name"`
	Model                string `json:"model"`
	Manufacturer         string `json:"manufacturer"`
	MaxAtmospheringSpeed string `json:"max_atmosphering_speed"`
}

type starShipResponse struct {
	vehicleResponse
	HyperdriveRating string `json:"hyperdrive_rating"`
}

func SayHello() string {
	return "Hello from the service"
}

func GetPeople() peopleResponse {
	res := utils.HttpGet("people")

	var peopleRes peopleResponse
	if err := json.NewDecoder(res.Body).Decode(&peopleRes); err != nil {
		log.Fatal(err)
	}

	return peopleRes
}

func GetPerson(id string) personResult {
	res := utils.HttpGet("people/" + id)

	var personRes personResponse
	if err := json.NewDecoder(res.Body).Decode(&personRes); err != nil {
		log.Fatal(err)
	}

	return personResult{
		Name:      personRes.Name,
		Homeworld: GetHomeWord(personRes.Homeworld),
		Vehicles:  GetVehicles(personRes.Vehicles),
		Starships: GetStarShips(personRes.Starships),
	}
}

func GetHomeWord(url string) homeWordResponse {
	res := utils.HttpGet(url)

	var homeWordRes homeWordResponse
	if err := json.NewDecoder(res.Body).Decode(&homeWordRes); err != nil {
		log.Fatal(err)
	}

	return homeWordRes
}

func GetVehicles(urls []string) (vehicles []vehicleResponse) {
	for _, url := range urls {
		res := utils.HttpGet(url)

		var vehicleRes vehicleResponse
		if err := json.NewDecoder(res.Body).Decode(&vehicleRes); err != nil {
			log.Fatal(err)
		}

		vehicles = append(vehicles, vehicleRes)
	}

	return vehicles
}

func GetStarShips(urls []string) (starShips []starShipResponse) {
	for _, url := range urls {
		res := utils.HttpGet(url)

		var starShipRes starShipResponse
		if err := json.NewDecoder(res.Body).Decode(&starShipRes); err != nil {
			log.Fatal(err)
		}

		starShips = append(starShips, starShipRes)
	}

	return starShips
}

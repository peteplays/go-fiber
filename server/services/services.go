package services

import (
	"encoding/json"
	"go-fiber-api/utils"
	"io"
	"log"
	"net/http"
)

type personResponse struct {
	Name      string
	Homeworld string
	Vehicles  []string
	Starships []string
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

func GetPeople() string {
	res, err := http.Get("people")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(body)
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

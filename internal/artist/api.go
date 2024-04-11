package artist

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	ApiURL       = "https://groupietrackers.herokuapp.com/api"
	ApiArtists   = "https://groupietrackers.herokuapp.com/api/artists"
	ApiLocations = "https://groupietrackers.herokuapp.com/api/locations"
	ApiDates     = "https://groupietrackers.herokuapp.com/api/dates"
	ApiRelation  = "https://groupietrackers.herokuapp.com/api/realation"
)

func GetResponseData(path string) ([]byte, error) {
	resp, err := http.Get(path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetArtists() ([]Artists, error) {
	var artists []Artists

	data, err := GetResponseData(ApiArtists)
	if err != nil {
		log.Println(err.Error())
		return artists, err
	}
	if err := json.Unmarshal(data, &artists); err != nil {
		log.Println(err.Error())
		return artists, err
	}
	return artists, nil
}

func getLocationById(id string) (Location, error) {
	var result Location

	data, err := GetResponseData(ApiLocations + "/" + id)
	if err != nil {
		log.Println(err.Error())
		return result, err
	}

	if err := json.Unmarshal(data, &result); err != nil {
		log.Println(err.Error())
		return result, err
	}
	return result, nil
}

func getConcertDateById(id string) (Date, error) {
	var result Date

	data, err := GetResponseData(ApiDates + "/" + id)
	if err != nil {
		log.Println(err.Error())
		return result, err
	}

	if err := json.Unmarshal(data, &result); err != nil {
		log.Println(err.Error())
		return result, err
	}
	return result, nil
}

func GetArtistById(id string) (Artist, error) {
	var result Artist
	data, err := GetResponseData(ApiArtists + "/" + id)
	if err != nil {
		log.Println(err.Error())
		return result, err
	}

	if err := json.Unmarshal(data, &result); err != nil {
		log.Println(err.Error())
		return result, err
	}
	if result.Id == 0 {
		return result, errors.New("404")
	}
	var location Location

	var date Date
	location, err = getLocationById(id)
	if err != nil {
		log.Println(err.Error())
		return result, err
	}

	result.locations = location.Locations

	date, err = getConcertDateById(id)
	if err != nil {
		log.Println(err.Error())
		return result, err
	}

	result.concertDates = date.Dates
	for i, loc := range result.locations {
		rel := Relation{Location: loc, Date: result.concertDates[i]}
		result.Rels = append(result.Rels, rel)
	}
	log.Printf("%v", result)
	return result, nil
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

)

const Url = "https://groupietrackers.herokuapp.com/api/"

type Artist struct {
	ID             int    
	Image          string  ` json:"image"`
	Name           string   `json:"name"`
	Members        []string `json:"members"`
	CreationDate   int      `json:"creationDate"`
	FirstAlbum     string   `json:"firstAlbum"`
}

type locations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}
type locationResponse struct {
	Index []locations `json:"index"`
}

type Dates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type DatesResponse struct {
	Index []Dates `json:"index"`
}


type Relation struct {
	ID             int                ` json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type RelationResponse struct {
	Index []Relation `json:"index"`
}
// FetchArtists fetches the list of artists from the external API.
func FetchArtists()([]Artist,error) {
	resp, err := http.Get(Url+"artists")
	if err != nil {
		return nil,fmt.Errorf("failed to fetch artist %#v", err)
	}
	defer resp.Body.Close()
	var artists []Artist
	// Decode the JSON response into the artists slice
	err = json.NewDecoder(resp.Body).Decode(&artists)
	if err != nil {
		return nil,fmt.Errorf("failed to decode artists %#v", err)
	}

	return artists,nil
}

func FetchLocation()([]locations,error) {
	resp, err := http.Get(Url+"locations")
	if err != nil {
		return nil,fmt.Errorf("failed to fetch artist %#v", err)
	}
	defer resp.Body.Close()
	var location locationResponse
	// Decode the JSON response into the artists slice
	err = json.NewDecoder(resp.Body).Decode(&location)
	if err != nil {
		return nil,fmt.Errorf("failed to decode artists %#v", err)
	}
	return location.Index,nil
}

func fetchDates()([]Dates,error){

	resp, err := http.Get(Url+"dates")
	if err != nil{
		return nil,fmt.Errorf("failed to fetch artist %#v", err)
	}
	defer resp.Body.Close()
	var dates DatesResponse
	// Decode the JSON response into the artists slice
	err = json.NewDecoder(resp.Body).Decode(&dates)
	if err != nil{
		return nil,fmt.Errorf("failed to decode artists %#v", err)
	}
	return dates.Index,nil
}


func FetchRelation()([]Relation,error) {
	resp, err := http.Get(Url + "relation")
	if err != nil {
		return nil,fmt.Errorf("failed to fetch artist %#v", err)
	}
	defer resp.Body.Close()
	var relation RelationResponse
	// Decode the JSON response into the artists slice
	err = json.NewDecoder(resp.Body).Decode(&relation)
	if err != nil {
		return nil,fmt.Errorf("failed to decode artists %#v", err)
	}
	return relation.Index,nil
}

func main(){
	// fmt.Println(FetchRelation())
	// fmt.Println(fetchDates())
	fmt.Println(FetchArtists())
	// fmt.Println(FetchLocation())
}
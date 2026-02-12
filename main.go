package main

import (
	"encoding/json"
	"fmt"
	"net/http"

)

// const Url = "https://groupietrackers.herokuapp.com/api/artists"
// const Urldt = "https://groupietrackers.herokuapp.com/api/dates"
type Artist struct {
	ID             int      `json:"id"`
	Image          string  ` json:"image"`
	Name           string   `json:"name"`
	Members        []string `json:"members"`
	CreationDate   int      `json:"creationDate"`
	FirstAlbum     string   `json:"firstAlbum"`
	// locations      locations
	// Dates          Dates
	// Relation 	  Relation
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

// // FetchArtists fetches the list of artists from the external API.
func FetchArtists(url string)([]Artist,error) {
	resp, err := http.Get(url)
	if err != nil {
		// return nil,fmt.Errorf("failed to fetch artist %#v", err)
		fmt.Println(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Body)
	var artists []Artist
	// Decode the JSON response into the artists slice
	err = json.NewDecoder(resp.Body).Decode(&artists)
	if err != nil {
		// return nil,fmt.Errorf("failed to decode artists %#v", err)
		// fmt.Errorf("failed to decode artists %#v", err)
		return nil,err
	}
	return artists,nil
}

func FetchLocation(url string)([]locations,error) {
	resp, err := http.Get(url)
	if err != nil {
		// return nil,fmt.Errorf("failed to fetch artist %#v", err)
		fmt.Println(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Body)
	var location locationResponse
	// Decode the JSON response into the artists slice
	err = json.NewDecoder(resp.Body).Decode(&location)
	if err != nil {
		// return nil,fmt.Errorf("failed to decode artists %#v", err)
		// fmt.Errorf("failed to decode artists %#v", err)
		return nil,err
	}
	return location.Index,nil
}

// FetchArtists fetches the list of artists from the external API.// 	resp, err := http.Get(Url)

func fetchDates(url string)[]Dates{

	resp, err := http.Get(url)
	if err != nil{
		// return nil,fmt.Errorf("failed to fetch artist %#v", err)
	}
	defer resp.Body.Close()
	var dates DatesResponse
	// Decode the JSON response into the artists slice
	err = json.NewDecoder(resp.Body).Decode(&dates)
	// if err != nil{
	// 	// return nil,fmt.Errorf("failed to decode artists %#v", err)
	// 	// fmt.Println(err)23
	// 	return err
	// }
	// return artists,nil
	return dates.Index
}

func main(){
	const Urldt = "https://groupietrackers.herokuapp.com/api/dates"
	const Urllc = "https://groupietrackers.herokuapp.com/api/locations"
	const Url = "https://groupietrackers.herokuapp.com/api/artists"
	fmt.Println(FetchArtists(Url))
	fmt.Println(fetchDates(Urldt))
	fmt.Println(FetchLocation(Urllc))
}
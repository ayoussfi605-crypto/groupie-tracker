package handler

import (
	"html/template"
	"net/http"
	"strconv"
)

type ArtistDetails struct {
	Artist
	Locations []string
	Dates     []string
	Relations map[string][]string
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idstr := r.URL.Query().Get("id")
	if idstr == "" {
		http.Error(w, "Missing id", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	var selected Artist
	found := false
	artists, err := FetchArtists()
	if err != nil {
		http.Error(w, "Error fetching artists", http.StatusInternalServerError)
		return
	}

	for _, a := range artists {
		if a.ID == id {
			selected = a
			found = true
			break
		}
	}

	if !found {
		http.Error(w, "Artist not found", http.StatusNotFound)
		return
	}

	locations, _ := FetchLocations()
	dates, _ := FetchDates()
	relation, _ := FetchRelation()

	var loc []string
	var dat []string
	var rel map[string][]string
	for _, l := range locations {
		if l.ID == id {
			loc = l.Locations
			break
		}
	}

	for _, d := range dates {
		if d.ID == id {
			dat = d.Dates
			break
		}
	}

	for _, r := range relation {
		if r.ID == id {
			rel = r.DatesLocations
			break
		}
	}

	details := ArtistDetails{
		Artist: selected,
		Locations: loc,
		Dates: dat,
		Relations: rel,
	}
	tmpl, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, details)
}

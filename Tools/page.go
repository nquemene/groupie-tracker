package tools

import (
	"fmt"
)

type Page struct {
	Artist
	Location []string
	Date     []string
}

func FetchPage(url string) ([]Page, error) {

	artists, err := FetchArtist(url + "/artists")

	if err != nil {
		return nil, fmt.Errorf("artists : %w", err)
	}

	dates, err := FetchDate(url + "/dates")

	if err != nil {
		return nil, fmt.Errorf("dates : %w", err)
	}

	locations, err := FetchLocation(url + "/locations")

	if err != nil {
		return nil, fmt.Errorf("locations : %w", err)
	}

	datesByID := make(map[int][]string, len(dates.Index))

	for _, d := range dates.Index {
		datesByID[d.ID] = d.Dates
	}

	locationsByID := make(map[int][]string, len(locations.Index))

	for _, l := range locations.Index {
		locationsByID[l.ID] = l.Locations
	}

	result := make([]Page, 0, len(artists))
	for _, a := range artists {
		result = append(result, Page{
			Artist:   a,
			Date:     datesByID[a.ID],
			Location: locationsByID[a.ID],
		})
	}

	return result, nil
}

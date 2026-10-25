package tools

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

func FetchArtist(url string) ([]Artist, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get(url)

	if err != nil {
		return nil, fmt.Errorf("Echec de la requête : %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Statut HTTP inattendu : %s", resp.Status)
	}

	var artist []Artist

	if err := json.NewDecoder(resp.Body).Decode(&artist); err != nil {
		return nil, err
	}

	return artist, nil
}

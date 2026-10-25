package tools

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

type EnvLocation struct {
	Index []Location `json:"index"`
}

func FetchLocation(url string) (*EnvLocation, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get(url)

	if err != nil {
		return nil, fmt.Errorf("Echec de la requête : %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Statut HTTP inattendu : %s", resp.Status)
	}

	var location EnvLocation

	if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
		return nil, err
	}

	return &location, nil
}

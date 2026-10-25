package tools

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type EnvDate struct {
	Index []Date `json:"index"`
}

func FetchDate(url string) (*EnvDate, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get(url)

	if err != nil {
		return nil, fmt.Errorf("Echec de la requête : %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Statut HTTP inattendu : %s", resp.Status)
	}

	var date EnvDate

	if err := json.NewDecoder(resp.Body).Decode(&date); err != nil {
		return nil, err
	}

	return &date, nil
}

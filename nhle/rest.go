// Package nhle implements a PlayerService that uses api-web.nhle.com as a datasource
package nhle

import (
	"encoding/json"
	"fmt"
	"github.com/StephenGriese/hello-api/roster"
	"net/http"
	"time"
)

const (
	BaseURLV1 = "https://api-web.nhle.com/v1/roster/PHI/20232024"
)

func NewPlayerService() roster.PlayerService {
	return playerService{
		BaseURL: BaseURLV1,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

type playerService struct {
	BaseURL    string
	HTTPClient *http.Client
}

var _ roster.PlayerService = playerService{}

func (ps playerService) Players() ([]roster.Player, error) {
	req, err := http.NewRequest("GET", BaseURLV1, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	res, err := ps.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	type Name struct {
		Default string `json:"default"`
	}

	type Player struct {
		ID            int  `json:"id"`
		FirstName     Name `json:"firstName"`
		LastName      Name `json:"lastName"`
		SweaterNumber int  `json:"sweaterNumber"`
	}

	toRosterPlayer := func(p Player, position roster.Position) roster.Player {
		return roster.Player{
			ID:            p.ID,
			FirstName:     p.FirstName.Default,
			LastName:      p.LastName.Default,
			SweaterNumber: p.SweaterNumber,
			Position:      position,
		}
	}

	apiResp := struct {
		Forwards   []Player `json:"forwards"`
		Defensemen []Player `json:"defensemen"`
		Goalies    []Player `json:"goalies"`
	}{}

	if err = json.NewDecoder(res.Body).Decode(&apiResp); err != nil {
		return nil, err
	}
	var result []roster.Player
	for _, p := range apiResp.Forwards {
		result = append(result, toRosterPlayer(p, roster.Forward))
	}
	for _, p := range apiResp.Defensemen {
		result = append(result, toRosterPlayer(p, roster.Defense))
	}
	for _, p := range apiResp.Goalies {
		result = append(result, toRosterPlayer(p, roster.Goalie))
	}
	return result, nil
}

/*
func (ps playerService) sendRequest(req *http.Request, v interface{}) error {
	res, err := ps.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	type Name struct {
		Default string `json:"default"`
	}

	type Player struct {
		ID            int  `json:"id"`
		FirstName     Name `json:"firstName"`
		LastName      Name `json:"lastName"`
		SweaterNumber int  `json:"sweaterNumber"`
	}

	apiResp := struct {
		Forwards   []Player `json:"forwards"`
		Defensemen []Player `json:"defensemen"`
		Goalies    []Player `json:"goalies"`
	}{}

	if err = json.NewDecoder(res.Body).Decode(&apiResp); err != nil {
		return err
	}

	return nil
}
*/

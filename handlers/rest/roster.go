package rest

import (
	"fmt"
	"github.com/StephenGriese/hello-api/nhle"
	"net/http"
	"sort"
)

func RosterHandler(w http.ResponseWriter, r *http.Request) {
	ps := nhle.NewPlayerService()
	players, err := ps.Players()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	sort.Slice(players, func(i, j int) bool {
		return players[i].SweaterNumber < players[j].SweaterNumber
	})
	for _, b := range players {
		fmt.Fprintf(w, "%2d   %-25s %s\n", b.SweaterNumber, b.FirstName+" "+b.LastName, b.Position)
	}
}

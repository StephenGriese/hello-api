package roster

type Position uint

const (
	Forward Position = iota
	Defense
	Goalie
)

func (p Position) String() string {
	switch p {
	case Forward:
		return "forward"
	case Defense:
		return "defense"
	case Goalie:
		return "goalie"
	default:
		return "unknown"
	}
}

type Player struct {
	ID            int      `json:"id"`
	FirstName     string   `json:"firstName"`
	LastName      string   `json:"lastName"`
	SweaterNumber int      `json:"sweaterNumber"`
	Position      Position `json:"position"`
}

type PlayerService interface {
	Players() ([]Player, error)
}

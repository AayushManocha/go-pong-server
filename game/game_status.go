package game

import "encoding/json"

type GameStatus int

const (
	CREATED GameStatus = iota
	READY
	IN_PLAY
	PAUSED
	FINISHED
	UNKNOWN
)

func (s GameStatus) String() string {
	switch s {
	case CREATED:
		return "CREATED"
	case READY:
		return "READY"
	case IN_PLAY:
		return "IN_PLAY"
	case PAUSED:
		return "PAUSED"
	case FINISHED:
		return "FINISHED"
	default:
		return "UNKNOWN"
	}
}

func ParseGameStatus(s string) GameStatus {
	switch s {
	case "CREATED":
		return CREATED
	case "READY":
		return READY
	case "IN_PLAY":
		return IN_PLAY
	case "PAUSED":
		return PAUSED
	case "FINISHED":
		return FINISHED
	default:
		return UNKNOWN
	}
}

func (s GameStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

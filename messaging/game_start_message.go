package messaging

type GameStartMessage struct {
	Type string `json:"type"`
}

func NewGameStartMessage() GameStartMessage {
	return GameStartMessage{
		Type: "GAME_START_MESSAGE",
	}
}

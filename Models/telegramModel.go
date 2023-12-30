package Models

type MessageData struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}

type SendPhoto struct {
	ChatID  string `json:"chat_id"`
	Photo   string `json:"photo"`
	Caption string `json:"caption"`
}

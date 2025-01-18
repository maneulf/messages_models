package models

type ApiGateWayMessage struct {
	Client    Client   `json:"client"`
	Metadata  Metadata `json:"metadata"`
	RequestID string   `json:"request_id"`
	Payload   Payload  `json:"payload"`
	Provider  Provider `json:"provider"`
}

type Client struct {
	UserID string `json:"user_id"`
}

type Metadata struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type Content struct {
	Text string `json:"text"`
}

type Payload struct {
	Content     Content `json:"content"`
	ContentType string  `json:"content_type"`
}

type Provider struct {
	Name string `json:"name"`
}

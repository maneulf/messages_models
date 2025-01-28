package models

type CsmlRequestMessage struct {
	Client struct {
		UserID string `json:"user_id"`
	} `json:"client"`
	Metadata struct {
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
	} `json:"metadata"`
	RequestID string `json:"request_id"`
	Payload   struct {
		Content struct {
			Text string `json:"text"`
		} `json:"content"`
		ContentType string `json:"content_type"`
	} `json:"payload"`
}

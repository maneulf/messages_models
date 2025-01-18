package models

type CsmlMessage struct {
	Client    CsmlClient         `json:"client"`
	Metadata  CsmlMetadata       `json:"metadata"`
	RequestID string             `json:"request_id"`
	Payload   CsmlMessagePayload `json:"payload"`
}

type CsmlClient struct {
	UserID string `json:"user_id"`
}

type CsmlMetadata struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type CsmlMessageContent struct {
	Text string `json:"text"`
}

type CsmlMessagePayload struct {
	Content     CsmlMessageContent `json:"content"`
	ContentType string             `json:"content_type"`
}

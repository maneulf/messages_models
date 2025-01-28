package models

import (
	"time"

	"gorm.io/gorm"
)

type CsmlResponseMessage struct {
	gorm.Model
	RequestID string `json:"request_id"`
	Client    struct {
		BotID     string `json:"bot_id"`
		UserID    string `json:"user_id"`
		ChannelID string `json:"channel_id"`
	} `json:"client"`
	ConversationEnd bool `json:"conversation_end"`
	Messages        []struct {
		Payload struct {
			ContentType string `json:"content_type"`
			Content     struct {
				Text    string `json:"text"`
				Title   string `json:"title"`
				Buttons []struct {
					Content struct {
						Title   string   `json:"title"`
						Payload string   `json:"payload"`
						Accepts []string `json:"accepts"`
						Accept  []string `json:"accept"`
					} `json:"content"`
					ContentType string `json:"content_type"`
				} `json:"buttons"`
			} `json:"content"`
		} `json:"payload"`
		InteractionOrder int    `json:"interaction_order"`
		ConversationID   string `json:"conversation_id"`
		Direction        string `json:"direction"`
	} `json:"messages"`
	ReceivedAt   time.Time `json:"received_at"`
	IsAuthorized bool      `json:"is_authorized"`
}

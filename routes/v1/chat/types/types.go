package types

type SubmitMessageRequest struct {
	User        int64  `json:"user"`
	ChatSession string `json:"chat_session"`
	Message     string `json:"message"`
}

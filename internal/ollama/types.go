package ollama

import "time"

type GenerateRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type GenerateResponse struct {
	Model     string    `json:"model"`
	CreatedAt time.Time `json:"created_at"`
	Response  string    `json:"response"`
	Done      bool      `json:"done"`
}

type Client	struct {
	baseURL		string
}

func NewClient() *Client {
	return &Client {
		baseURL: "http://localhost:11434",
	}
}

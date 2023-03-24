package api

type CompletionRequest struct {
	Prompt      string  `json:"prompt"`
	MaxTokens   int     `json:"max_tokens,omitempty"`
	Temperature float64 `json:"temperature,omitempty"`
	TopP        float64 `json:"top_p,omitempty"`
}

type CompletionResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int64    `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Text string `json:"text"`
}

type Engine struct {
	ID          string `json:"id"`
	Object      string `json:"object"`
	Owner       string `json:"owner"`
	Created     int64  `json:"created"`
	MaxTokens   int    `json:"max_tokens"`
	Ready       bool   `json:"ready"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type EnginesResponse struct {
	Object string   `json:"object"`
	Data   []Engine `json:"data"`
}

type FineTunedModel struct {
	ID        string `json:"id"`
	Object    string `json:"object"`
	Created   int64  `json:"created"`
	Updated   int64  `json:"updated"`
	Engine    string `json:"engine"`
	Ready     bool   `json:"ready"`
	AuthToken string `json:"auth_token"`
	Expires   int64  `json:"expires"`
}

type FineTunedModelsResponse struct {
	Object string           `json:"object"`
	Data   []FineTunedModel `json:"data"`
}

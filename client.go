package api

import (
	"net/http"
)

type Config struct {
	ApiKey        string
	BaseUrl       string
	DefaultEngine string
	HttpClient    *http.Client
}

type Client struct {
	Config
}

func NewClient(config Config) (*Client, error) {
	// Set default values if not provided
	if config.BaseUrl == "" {
		config.BaseUrl = "https://api.openai.com/v1/"
	}
	if config.DefaultEngine == "" {
		config.DefaultEngine = "davinci-codex"
	}
	if config.HttpClient == nil {
		config.HttpClient = &http.Client{}
	}

	return &Client{
		Config: config,
	}, nil
}

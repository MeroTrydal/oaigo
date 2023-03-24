package oaigo

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func (client *Client) sendRequest(method, url string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, client.BaseUrl+url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+client.ApiKey)

	resp, err := client.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, NewApiError(resp.StatusCode, string(respBody))
	}
	return respBody, nil
}

func (client *Client) CreateCompletion(prompt string) (*CompletionResponse, error) {
	requestData := &CompletionRequest{
		Prompt:      prompt,
		MaxTokens:   50,
		Temperature: 1,
	}

	requestJson, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}

	responseJson, err := client.sendRequest(http.MethodPost, client.CompletionsEndpoint(), bytes.NewBuffer(requestJson))
	if err != nil {
		return nil, err
	}

	responseData := &CompletionResponse{}
	err = json.Unmarshal(responseJson, responseData)
	if err != nil {
		return nil, err
	}
	return responseData, nil
}

func (client *Client) ListEngines() (*EnginesResponse, error) {
	responseJson, err := client.sendRequest(http.MethodGet, "engines", nil)
	if err != nil {
		return nil, err
	}

	responseData := &EnginesResponse{}
	err = json.Unmarshal(responseJson, responseData)
	if err != nil {
		return nil, err
	}
	return responseData, nil
}

func (client *Client) GetEngine(engineID string) (*Engine, error) {
	responseJson, err := client.sendRequest(http.MethodGet, "engines/"+engineID, nil)
	if err != nil {
		return nil, err
	}

	responseData := &Engine{}
	err = json.Unmarshal(responseJson, responseData)
	if err != nil {
		return nil, err
	}
	return responseData, nil
}

func (client *Client) ListFineTunedModels() (*FineTunedModelsResponse, error) {
	responseJson, err := client.sendRequest(http.MethodGet, "fine-tuned-models", nil)
	if err != nil {
		return nil, err
	}

	responseData := &FineTunedModelsResponse{}
	err = json.Unmarshal(responseJson, responseData)
	if err != nil {
		return nil, err
	}
	return responseData, nil
}

func (client *Client) GetFineTunedModel(modelID string) (*FineTunedModel, error) {
	responseJson, err := client.sendRequest(http.MethodGet, "fine-tuned-models/"+modelID, nil)
	if err != nil {
		return nil, err
	}

	responseData := &FineTunedModel{}
	err = json.Unmarshal(responseJson, responseData)
	if err != nil {
		return nil, err
	}
	return responseData, nil
}

func (client *Client) DeleteFineTunedModel(modelID string) error {
	_, err := client.sendRequest(http.MethodDelete, "fine-tuned-models/"+modelID, nil)
	return err
}

func (client *Client) CreateChat(messages []Message, maxTokens int, temperature, topP float64) (*ChatResponse, error) {
	requestData := &ChatRequest{
		Messages:    messages,
		MaxTokens:   maxTokens,
		Temperature: temperature,
		TopP:        topP,
	}

	requestJson, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}

	responseJson, err := client.sendRequest(http.MethodPost, client.ChatEndpoint(), bytes.NewBuffer(requestJson))
	if err != nil {
		return nil, err
	}

	responseData := &ChatResponse{}
	err = json.Unmarshal(responseJson, responseData)
	if err != nil {
		return nil, err
	}
	return responseData, nil
}

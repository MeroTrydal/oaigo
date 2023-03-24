package api

func (client *Client) CompletionsEndpoint() string {
	return "engines/" + client.DefaultEngine + "/completions"
}

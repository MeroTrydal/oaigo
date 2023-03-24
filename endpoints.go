package oaigo

func (client *Client) CompletionsEndpoint() string {
	return "engines/" + client.DefaultEngine + "/completions"
}

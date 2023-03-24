package oaigo

func (client *Client) CompletionsEndpoint() string {
	return "engines/" + client.DefaultEngine + "/completions"
}

func (client *Client) ChatEndpoint() string {
	return "engines/" + client.DefaultEngine + "/chat"
}

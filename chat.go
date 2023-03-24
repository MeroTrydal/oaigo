package oaigo

func (client *Client) ChatEndpoint() string {
	return "engines/" + client.DefaultEngine + "/chat"
}

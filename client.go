package plugnotas

type Client struct {
	token   string
	sandbox bool
}

func NewClient(token string, sandbox bool) *Client {
	return &Client{
		token:   token,
		sandbox: sandbox,
	}
}

func (c *Client) GetToken() string {
	return c.token
}

func (c *Client) GetEndpoint() string {

	if c.sandbox {
		return SandBoxUrl
	}

	return ProductionUrl

}

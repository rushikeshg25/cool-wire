package wire

import (
	"github.com/google/uuid"
)

type Client struct {
	clientId uuid.UUID
	port     int
	host     string
}

func NewClient(port int, host string) *Client {
	return &Client{
		clientId: uuid.New(),
		port:     port,
		host:     host,
	}
}

func (c *Client) SendQuery(q Query) {

}

func (c *Client) SendResponse(r Response) {

}

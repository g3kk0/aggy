package gdax

import (
	gdax "github.com/preichenberger/go-gdax"
)

type Client struct {
	Conn *gdax.Client
}

func NewClient(key, secret, passphrase string) *Client {
	var c Client
	c.Conn = gdax.NewClient(secret, key, passphrase)
	return &c
}

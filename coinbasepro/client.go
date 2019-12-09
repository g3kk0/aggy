package coinbasepro

import (
	coinbasepro "github.com/g3kk0/go-coinbasepro"
)

type Client struct {
	Conn *coinbasepro.Client
}

func NewClient(key, secret, passphrase string) *Client {
	var c Client
	c.Conn = coinbasepro.NewClient()

	c.Conn.UpdateConfig(&coinbasepro.ClientConfig{
		BaseURL:    "https://api.pro.coinbase.com",
		Key:        key,
		Passphrase: passphrase,
		Secret:     secret,
	})

	return &c
}

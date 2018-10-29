package binance

import (
	binance "github.com/adshao/go-binance"
)

type Client struct {
	Conn *binance.Client
}

func NewClient(key, secret string) *Client {
	var c Client
	c.Conn = binance.NewClient(key, secret)

	return &c
}

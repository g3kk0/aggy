package main

import (
	"context"
	"fmt"
	"log"
	"os"

	binance "github.com/adshao/go-binance"
)

type Client struct {
	Conn *binance.Client
}

func NewClient() *Client {
	c := &Client{}
	c.Conn = binance.NewClient(os.Getenv("BINANCE_KEY"), os.Getenv("BINANCE_SECRET"))
	return c
}

func (c *Client) GetBalances() []binance.Balance {
	res, err := c.Conn.NewGetAccountService().Do(context.Background())
	if err != nil {
		log.Println(err)
		return nil
	}

	return res.Balances
}

// pick up here
func (c *Client) ParseBalances(b []binance.Balance) map[string]map[string]string {
	x := make(map[string]map[string]string)
	return x
}

func main() {
	bc := NewClient()

	balancesRaw := bc.GetBalances()
	fmt.Printf("balancesRaw = %+v\n", balancesRaw)

	balances := bc.ParseBalances(balancesRaw)
	fmt.Printf("balances = %+v\n", balances)

	fmt.Println("done")
}

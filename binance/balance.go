package binance

import (
	"context"
	"fmt"
	"log"
)

func (c *Client) GetBalances() map[string]float64 {
	balances := make(map[string]float64)

	res, err := c.Conn.NewGetAccountService().Do(context.Background())
	if err != nil {
		log.Println(err)
		return nil
	}

	fmt.Printf("res = %+v\n", res)

	return balances
}

package gdax

import (
	"strconv"
)

type Asset struct {
	Symbol string
	Amount float64
}

func (c *Client) Assets(quoteCurrency string) ([]Asset, error) {
	var assets []Asset

	accounts, err := c.Conn.GetAccounts()
	if err != nil {
		return assets, err
	}

	// refactor (clear out zero balance accounts)
	for _, a := range accounts {
		var asset Asset
		asset.Symbol = a.Currency

		amount, err := strconv.ParseFloat(a.Available, 64)
		if err != nil {
			return assets, err
		}

		asset.Amount = amount

		assets = append(assets, asset)
	}

	return assets, nil
}

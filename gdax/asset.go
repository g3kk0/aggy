package gdax

import (
	"fmt"
	"strconv"

	forex "github.com/g3kk0/go-forex"
)

var FiatSymbols = []string{"EUR", "GBP", "USD"}

type Asset struct {
	Name          string
	Symbol        string
	Amount        float64
	QuoteCurrency string
	Value         float64
}

func (c *Client) Assets() ([]Asset, error) {
	var assets []Asset

	accounts, err := c.Conn.GetAccounts()
	if err != nil {
		return assets, err
	}

	for _, a := range accounts {
		var asset Asset
		asset.Symbol = a.Currency

		amount, err := strconv.ParseFloat(a.Available, 64)
		if err != nil {
			return assets, err
		}

		asset.Amount = amount

		asset.QuoteCurrency = "GBP"

		var fiatSymbol bool
		for _, v := range FiatSymbols {
			if asset.Symbol == v {
				fiatSymbol = true
				break
			}
		}

		if asset.Symbol == asset.QuoteCurrency {
			fmt.Printf("skipping = %+v\n", asset.Symbol)
			asset.Value = asset.Amount
			//continue
		} else if fiatSymbol {
			value, err := convertFiat(asset.Symbol, "GBP", asset.Amount)
			if err != nil {
				return assets, err
			}

			asset.Value = value

			fmt.Printf("forex convert = %+v\n", asset.Symbol)
		} else {
			//value := convertCrypto()
			//asset.Value = value

			fmt.Printf("crypto convert = %+v\n", asset.Symbol)
		}

		assets = append(assets, asset)
	}

	return assets, nil
}

func convertFiat(from, to string, amount float64) (float64, error) {
	fc := forex.NewClient()

	amountStr := fmt.Sprintf("%f", amount)

	params := map[string]string{"from": from, "to": to, "amount": amountStr}

	conversion, err := fc.Convert(params)
	if err != nil {
		return 0, err
	}

	return conversion.Result, nil
}

func convertCrypto() {}

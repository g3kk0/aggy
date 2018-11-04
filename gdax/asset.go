package gdax

import (
	"fmt"
	"strconv"
	"strings"

	coinmarketcap "github.com/g3kk0/go-coinmarketcap"
	forex "github.com/g3kk0/go-forex"
)

type Asset struct {
	Name          string
	Symbol        string
	Amount        float64
	QuoteCurrency string
	Value         float64
}

func (c *Client) Assets(quoteCurrency string) ([]Asset, error) {
	var assets []Asset
	var cryptos []string

	fiatSymbols := []string{"EUR", "GBP", "USD"}

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
		asset.QuoteCurrency = quoteCurrency

		var fiatSymbol bool
		for _, v := range fiatSymbols {
			if asset.Symbol == v {
				fiatSymbol = true
				break
			}
		}

		if asset.Symbol == quoteCurrency {
			fmt.Printf("skipping = %+v\n", asset.Symbol)
			asset.Value = asset.Amount
			//continue
		} else if fiatSymbol {
			value, err := fiatValue(asset.Symbol, quoteCurrency, asset.Amount)
			if err != nil {
				return assets, err
			}

			asset.Value = value

			fmt.Printf("forex convert = %+v\n", asset.Symbol)
		} else {
			cryptos = append(cryptos, asset.Symbol)

			//value := convertCrypto()
			//asset.Value = value

			fmt.Printf("crypto convert = %+v\n", asset.Symbol)
		}

		assets = append(assets, asset)
	}

	fmt.Printf("assets = %+v\n", assets)
	fmt.Printf("cryptos = %+v\n", cryptos)
	// lookup crypto value

	cValue := cryptoValue(assets, cryptos, quoteCurrency)

	fmt.Printf("cValue = %+v\n", cValue)

	return assets, nil
}

func fiatValue(from, to string, amount float64) (float64, error) {
	fc := forex.NewClient()

	amountStr := fmt.Sprintf("%f", amount)

	params := map[string]string{"from": from, "to": to, "amount": amountStr}

	conversion, err := fc.Convert(params)
	if err != nil {
		return 0, err
	}

	return conversion.Result, nil
}

func cryptoValue(assets []Asset, cryptos []string, quoteCurrency string) []Asset {
	cmc := coinmarketcap.NewClient()

	//params := map[string]string{"symbol": "bch,eth"}
	params := map[string]string{"symbol": strings.Join(cryptos, ","), "convert": quoteCurrency}
	quotes, err := cmc.QuotesLatest(params)
	if err != nil {
		panic(err)
	}

	for i, a := range assets {
		for _, c := range cryptos {
			if a.Symbol == c {

				price := quotes.Data[c].Quote[quoteCurrency].Price

				// get GBP value of a single sotoshi

				priceInPence := int64(price * 100)
				fmt.Printf("priceInPence = %+v\n", priceInPence)

				// below is wrong
				assets[i].Value = 100
			}
		}
	}

	//fmt.Printf("quotes = %+v\n", quotes)
	fmt.Printf("assets1 = %+v\n", assets)

	return assets
}

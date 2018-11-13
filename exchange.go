package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/g3kk0/aggy/gdax"
	coinmarketcap "github.com/g3kk0/go-coinmarketcap"
	forex "github.com/g3kk0/go-forex"
)

type Exchange struct {
	Type       string
	Key        string
	Secret     string
	Passphrase string
}

type Response struct {
	Value    float64 `json:"value"`
	Currency string  `json:"currency"`
	Holdings []Asset `json:"holdings"`
}

type Asset struct {
	Exchange      string  `json:"exchange"`
	Name          string  `json:"name"`
	Symbol        string  `json:"symbol"`
	Price         float64 `json:"price"`
	Amount        float64 `json:"amount"`
	QuoteCurrency string  `json:"quote_currency"`
	Value         float64 `json:"value"`
}

func NewExchange(t, key, secret, passphrase string) *Exchange {
	return &Exchange{
		Type:       t,
		Key:        key,
		Secret:     secret,
		Passphrase: passphrase,
	}
}

func (e *Exchange) Value(cmcKey, quoteCurrency string) (Response, error) {
	var r Response
	var err error

	switch e.Type {
	case "gdax":
		gc := gdax.NewClient(e.Key, e.Secret, e.Passphrase)

		assets, err := gc.Assets(quoteCurrency)
		if err != nil {
			return r, err
		}

		// do this better (avoid the copy)
		for _, a := range assets {
			asset := Asset{
				Exchange: "gdax",
				Symbol:   a.Symbol,
				Amount:   a.Amount,
			}

			r.Holdings = append(r.Holdings, asset)
		}
	case "binance":
		//bc := binance.NewClient(e.Key, e.Secret)
		fmt.Println("checking binance")
	default:
		return r, errors.New("unknown exchange type")
	}

	// loop over all assets and get values
	fiatSymbols := []string{"EUR", "GBP", "USD"}
	var cryptos []string

	for i, a := range r.Holdings {
		var fiatSymbol bool
		for _, e := range fiatSymbols {
			if a.Symbol == e {
				fiatSymbol = true
				break
			}
		}

		if a.Symbol == quoteCurrency {
			a.Value = a.Amount
		} else if fiatSymbol {
			value, err := fiatValue(a.Symbol, quoteCurrency, a.Amount)
			if err != nil {
				return r, err
			}

			r.Holdings[i].QuoteCurrency = quoteCurrency
			r.Holdings[i].Value = value
		} else {
			cryptos = append(cryptos, a.Symbol)
		}
	}

	//	fmt.Printf("r = %+v\n", r)

	err = cryptoValue(r.Holdings, cryptos, cmcKey, quoteCurrency)
	if err != nil {
		return r, err
	}

	// compute totals
	r.Currency = quoteCurrency

	for i, _ := range r.Holdings {
		r.Value = r.Value + r.Holdings[i].Value
	}

	//	fmt.Printf("r = %+v\n", r)

	return r, err
}

func cryptoValue(holdings []Asset, cryptos []string, cmcKey, quoteCurrency string) error {
	var err error

	cmc := coinmarketcap.NewClient(cmcKey)

	params := map[string]string{"symbol": strings.Join(cryptos, ","), "convert": quoteCurrency}

	quotes, err := cmc.QuotesLatest(params)
	if err != nil {
		return err
	}

	for _, c := range cryptos {
		for i, h := range holdings {
			if c == h.Symbol {
				price := quotes.Data[h.Symbol].Quote[quoteCurrency].Price
				value := (price / 100000000) * (h.Amount * 100000000)

				holdings[i].Price = price
				holdings[i].QuoteCurrency = quoteCurrency
				holdings[i].Value = value
			}
		}
	}

	return err
}

func fiatValue(from, to string, amount float64) (float64, error) {
	fc := forex.NewClient()

	amountStr := fmt.Sprintf("%f", amount)

	params := map[string]string{"from": from, "to": to, "amount": amountStr}

	c, err := fc.Convert(params)
	if err != nil {
		return 0, err
	}

	return c.Result, nil
}

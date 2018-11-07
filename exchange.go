package main

import (
	"errors"
	"fmt"

	"github.com/g3kk0/aggy/gdax"
)

type Exchange struct {
	Type       string
	Key        string
	Secret     string
	Passphrase string
}

type Response struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
	// define type for this!
	//	Exchanges []
	Holdings []Asset `json:"holdings"`
	//Holdings []gdax.Asset
}

type Asset struct {
	Exchange      string  `json:"exchange"`
	Name          string  `json:"name"`
	Symbol        string  `json:"symbol"`
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

func (e *Exchange) Value(quoteCurrency string) (Response, error) {
	var r Response
	var err error

	//	fiatSymbols := []string{"EUR", "GBP", "USD"}
	//	cryptoSymbols := []string{}

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

	return r, err
}

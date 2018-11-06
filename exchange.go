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
	Holdings interface{} `json:"holdings"`
	//Holdings []gdax.Asset
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

	switch e.Type {
	case "gdax":
		gc := gdax.NewClient(e.Key, e.Secret, e.Passphrase)
		r.Holdings, err = gc.Assets(quoteCurrency)
		if err != nil {
			return r, err
		}

	case "binance":
		//bc := binance.NewClient(e.Key, e.Secret)
		fmt.Println("checking binance")
	default:
		return r, errors.New("unknown exchange type")
	}

	return r, nil
}

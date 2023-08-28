package main

import (
	"encoding/json"
	"os"
)

type Account struct {
	IgnoredField string `json:"-"`
	Number       int    `json:"n"`
	Balance      int    `json:"s"`
}

func main() {
	account := Account{Number: 1, Balance: 100, IgnoredField: "this should not appear in the json"}
	res, err := json.Marshal(account)
	if err != nil {
		println(err)
	}
	println(string(res))

	encoder := json.NewEncoder(os.Stdout)
	err = encoder.Encode(account)
	if err != nil {
		println(err)
	}

	rawJson := []byte(`{"n":2,"s":200}`)
	var accountX Account
	err = json.Unmarshal(rawJson, &accountX)
	if err != nil {
		println(err)
	}
	println(accountX.Balance)
}

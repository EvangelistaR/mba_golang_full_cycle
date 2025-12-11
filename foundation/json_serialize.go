package main

import (
	"encoding/json"
	"os"
)

type Account struct {
	Number  int `json:"number"`
	Balance int `json:"balance"`
}

func main() {
	account := Account{Number: 1, Balance: 1_000}

	res, err := json.Marshal(account)

	if err != nil {
		println(err)
	}

	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(account)

	if err != nil {
		println(err)
	}

	newJson := []byte(`{"number": 2, "balance": 3555}`)
	var newAccount Account

	err = json.Unmarshal(newJson, &newAccount)
	if err != nil {
		println(err)
	}

	println(newAccount.Balance)

}

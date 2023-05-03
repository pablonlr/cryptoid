package main

import (
	crypto "github.com/pablonlr/cryptoid"
)

func main() {
	client := crypto.NewCryptoIDClient("your_key")
	resp, err := client.BlockCount("btc")
	if err != nil {
		panic(err)
	}
	println(resp)
}

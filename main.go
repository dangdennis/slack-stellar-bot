package main

import (
	"context"
	"fmt"
	// "github.com/stellar/go/clients/horizon"
	"github.com/stellar/go/clients/horizonclient"
	// "github.com/stellar/go/keypair"
	hProtocol "github.com/stellar/go/protocols/horizon"
	// "io/ioutil"
	// "log"
	// "net/http"
	"time"
)

func main() {
	// Streaming transactions
	client := horizonclient.DefaultTestNetClient
	// all transactions
	transactionRequest := horizonclient.TransactionRequest{Limit: 50}

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		// Stop streaming after 60 seconds.
		time.Sleep(60 * time.Second)
		cancel()
	}()

	printHandler := func(tr hProtocol.Transaction) {
		fmt.Println("TRANSACTION", tr.ID)
	}

	err := client.StreamTransactions(ctx, transactionRequest, printHandler)
	if err != nil {
		fmt.Println(err)
	}

	// Generating account

	// pair, err := keypair.Random()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println(pair.Seed())
	// // SD4HKWZ7JW2UTTCQNY5BH5GBCEJL5YI23COBSFJGG6M2VQFP47WMHHCY
	// address := pair.Address()
	// log.Println(address)
	// // address := "GBDLNH4K6GJ7EYYJ62NI5SVFTB7FY6P2SFOWQ5BHHZURI6CT25BN3NOF"
	// //
	// resp, err := http.Get("https://friendbot.stellar.org/?addr=" + address)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(body))

	// account, err := horizon.DefaultTestNetClient.LoadAccount(address)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Balances for account: ", address)

	// for _, balance := range account.Balances {
	// 	log.Println(balance)
	// }
}

package accounts

import (
	"fmt"
	"testing"
	// "github.com/stellar/go/build"

	"github.com/stellar/go/clients/horizon"
	// "github.com/stellar/go/keypair"
	"io/ioutil"
	"log"
	"net/http"
)

func TestAccounts(t *testing.T) {
	// pair, err := keypair.Random()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println(pair.Seed())
	// address := pair.Address()
	seed := "SD4HKWZ7JW2UTTCQNY5BH5GBCEJL5YI23COBSFJGG6M2VQFP47WMHHCY"
	address := "GBDLNH4K6GJ7EYYJ62NI5SVFTB7FY6P2SFOWQ5BHHZURI6CT25BN3NOF"

	log.Println(seed)
	log.Println(address)

	resp, err := http.Get("https://friendbot.stellar.org/?addr=" + address)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))

	account, err := horizon.DefaultTestNetClient.LoadAccount(address)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Balances for account: ", address)

	for _, balance := range account.Balances {
		log.Println(balance)
	}
}

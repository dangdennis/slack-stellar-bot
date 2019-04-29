package transactions

import (
	"context"
	"fmt"
	"github.com/stellar/go/clients/horizonclient"
	hProtocol "github.com/stellar/go/protocols/horizon"
	"testing"
	"time"
)

func TestStreamTransactions(t *testing.T) {
	client := horizonclient.DefaultTestNetClient
	// all transactions
	transactionRequest := horizonclient.TransactionRequest{Cursor: "760209215489"}

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		// Stop streaming after 60 seconds.
		time.Sleep(60 * time.Second)
		cancel()
	}()

	printHandler := func(tr hProtocol.Transaction) {
		fmt.Println(tr)
	}

	err := client.StreamTransactions(ctx, transactionRequest, printHandler)
	if err != nil {
		fmt.Println(err)
	}
}

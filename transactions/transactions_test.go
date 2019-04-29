package transactions

import (
	"fmt"
	"github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"
	"testing"
)

func TestTransactions(t *testing.T) {

	// seed1 := "SD4HKWZ7JW2UTTCQNY5BH5GBCEJL5YI23COBSFJGG6M2VQFP47WMHHCY"
	src := "GBDLNH4K6GJ7EYYJ62NI5SVFTB7FY6P2SFOWQ5BHHZURI6CT25BN3NOF"

	// seed2 := "SBQH63YSWBW4VXIZ7ZLFP7DSPPYETVNI33BPPKWVC2SZ4IRPUTUI7OSC"
	dest := "GAKPB2WR2AUBQBHQLBWVX4ORFSA5KLRMGQWF3D2YM3MMCACLCQ5H6NGA"

	// src := "SCZANGBA5YHTNYVVV4C3U252E2B6P6F5T3U6MM63WBSBZATAQI3EBTQ4"
	// dest := "GA2C5RFPE6GCKMY3US5PAB6UZLKIGSPIUKSLRB6Q723BM2OARMDUYEJ5"

	if _, err := horizon.DefaultTestNetClient.LoadAccount(dest); err != nil {
		panic(err)
	}

	tx, err := build.Transaction(
		build.TestNetwork,
		build.SourceAccount{AddressOrSeed: src},
		build.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
		build.Payment(
			build.Destination{AddressOrSeed: dest},
			build.NativeAmount{Amount: "10"},
		),
	)

	if err != nil {
		panic(err)
	}

	txe, err := tx.Sign(src)
	if err != nil {
		panic(err)
	}

	txeB64, err := txe.Base64()
	if err != nil {
		panic(err)
	}

	resp, err := horizon.DefaultTestNetClient.SubmitTransaction(txeB64)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successful Transaction:")
	fmt.Println("Ledger:", resp.Ledger)
	fmt.Println("Hash:", resp.Hash)
}

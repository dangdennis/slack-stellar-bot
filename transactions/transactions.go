package transactions

import (
	"fmt"
	"github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"
)

func main() {

	seed1 := "SD4HKWZ7JW2UTTCQNY5BH5GBCEJL5YI23COBSFJGG6M2VQFP47WMHHCY"
	src := "GBDLNH4K6GJ7EYYJ62NI5SVFTB7FY6P2SFOWQ5BHHZURI6CT25BN3NOF"

	seed2 := "SBQH63YSWBW4VXIZ7ZLFP7DSPPYETVNI33BPPKWVC2SZ4IRPUTUI7OSC"
	dest := "GAKPB2WR2AUBQBHQLBWVX4ORFSA5KLRMGQWF3D2YM3MMCACLCQ5H6NGA"

	if _, err := horizon.DefaultTestNetClient.LoadAccount(dest); err != nil {
		panic(err)
	}

	tx, err := build.Transaction(
		build.TestNetwork,
		build.SourceAccount(src),
		build.AutoSequence(horizon.DefaultTestNetClient),
		build.Payment(
			build.Destination{dest},
			build.NativeAmount{"10"},
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

package main

import (
	"fmt"
	// "flag"
	// "os"
	"btc/key"
)

func main() {
	kg := key.NewGenerator()
	//argsWithoutProg := os.Args[1:]

	private, public := kg.NewPair()

	fmt.Println("Private key: ", private.Base58())
	fmt.Println("             ", private.Hex())
	fmt.Println("Public key:  ", public.Base58())
	fmt.Println("             ", public.HexCompressed())
}

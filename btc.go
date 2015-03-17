package main

import(
	"fmt"
	// "flag"
	// "os"
	"btc/key"
	"btc/key/ec"
)

func main(){
	kg := key.NewGenerator()
	//argsWithoutProg := os.Args[1:]
	
	private, public := kg.NewPair()
	curve := ec.NewCurve()

	fmt.Println("GeneratorPoint: x ", curve.GeneratorPoint().XHex())
	fmt.Println("GeneratorPoint: y ", curve.GeneratorPoint().YHex())
	fmt.Println("Prime: ", curve.Prime().BigInt())

	fmt.Println("Private key: ", private.Base58())
	fmt.Println("             ", private.Hex())
	fmt.Println("Public key:  ", public.Base58(), public.Hex())
	fmt.Println("             ", public.Hex())
}

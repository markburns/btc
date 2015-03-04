package main

import(
	"fmt"
	"btc/keys"
)

func main(){
	kg := keys.NewGenerator()
	result := kg.NewPair()

	fmt.Println("Private key: ", result.Private().Hex())
	fmt.Println("Public key:  ", result.Public().Hex())

}

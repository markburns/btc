package ec

import (. "math/big")

type Wif struct {
	Int
}

func W(i interface{}) (w *Wif){
	defer func(){
		if r := recover(); r != nil {
			value := i.(*Int)
		
		  w  = NewWif(value.String())
		}
	}()

	value := i.(string)

	w  = NewWif(value)
	return w
}

func NewWif(s string) *Wif{
	b, success := NewInt(0).SetString(s, 10)

	if !success{ msg:= "Couldn't parse: " + s; panic(msg) }
	return &Wif{*b}
}

func (w *Wif) BigInt() *Int{
	return &w.Int
}

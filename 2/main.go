package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	b1, _ := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	b2, _ := hex.DecodeString("686974207468652062756c6c277320657965")
	res := make([]byte, len(b1))
	for i := range res {
		res[i] = b1[i] ^ b2[i]
	}
	fmt.Println(hex.EncodeToString(res))
}

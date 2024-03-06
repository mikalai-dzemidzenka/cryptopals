package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func Hex2Base64(s string) (string, error) {
	b, err := hex.DecodeString(s)
	if err != nil {
		return "", err
	}

	str := base64.StdEncoding.EncodeToString(b)
	return str, nil
}

func main() {
	fmt.Println(Hex2Base64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"))
}

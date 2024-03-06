package main

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	f, err := os.ReadFile("7/7.txt")
	pif(err)
	f, err = base64.StdEncoding.DecodeString(string(f))
	pif(err)

	res := DecryptAes128Ecb(f, []byte("YELLOW SUBMARINE"))
	fmt.Println(string(res))
}

func DecryptAes128Ecb(data, key []byte) []byte {
	cipher, _ := aes.NewCipher(key)
	decrypted := make([]byte, len(data))
	size := 16

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		cipher.Decrypt(decrypted[bs:be], data[bs:be])
	}

	return decrypted
}

func pif(err error) {
	if err != nil {
		panic(err)
	}
}

package main

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	key := []byte("YELLOW SUBMARINE")
	test := EncryptAes128Ecb([]byte("roflan"), key)
	res := DecryptAes128Ecb(test, key)

	f, _ := os.ReadFile("10/10.txt")
	s, _ := base64.StdEncoding.DecodeString(string(f))
	_ = s

	iv := make([]byte, 16)
	h := EncryptAes128Cbc(iv, []byte("yellow submarine and pidors"), key)
	res = DecryptAes128Cbc(iv, h, key)
	fmt.Println(string(res))

	res = DecryptAes128Cbc(iv, s, key)
	fmt.Println(string(res))
}

func pad(b []byte, alignment int) []byte {
	r := alignment - (len(b) % alignment)
	for i := 0; i < r; i++ {
		b = append(b, byte(r))
	}
	return b
}

func DecryptAes128Cbc(iv, data, key []byte) []byte {
	cipher, _ := aes.NewCipher(key)
	decrypted := make([]byte, len(data))
	size := 16

	var prev []byte
	for i, j := len(data)-size, len(data); i >= size; i, j = i-size, j-size {
		prev = make([]byte, size)
		cipher.Decrypt(prev, data[i:j])
		for k := range prev {
			decrypted[i+k] = data[i-size+k] ^ prev[k]
		}
	}

	prev = make([]byte, size)
	cipher.Decrypt(prev, data[:size])
	for j := range prev {
		decrypted[j] = prev[j] ^ iv[j]
	}

	return decrypted
}

func EncryptAes128Cbc(iv, data, key []byte) []byte {
	data = pad(data, len(key))
	cipher, _ := aes.NewCipher(key)
	encrypted := make([]byte, len(data))
	size := 16

	prev := iv
	for i := 0; i < len(data); i += size {
		cipher.Encrypt(encrypted[i:i+size], x(data[i:i+size], prev))
		prev = encrypted[i : i+size]
	}

	return encrypted
}

func x(a, b []byte) []byte {
	if len(a) != len(b) {
		panic("rofl")
	}
	res := make([]byte, len(a))
	for i := range a {
		res[i] = a[i] ^ b[i]
	}
	return res
}

func DecryptAes128Ecb(data, key []byte) []byte {
	cipher, _ := aes.NewCipher(key)
	encrypted := make([]byte, len(data))
	size := 16

	for i, j := 0, size; i < len(data); i, j = i+size, j+size {
		cipher.Decrypt(encrypted[i:j], data[i:j])
	}

	return encrypted
}

func EncryptAes128Ecb(data, key []byte) []byte {
	data = pad(data, len(key))
	cipher, _ := aes.NewCipher(key)
	encrypted := make([]byte, len(data))
	size := 16

	for i, j := 0, size; i < len(data); i, j = i+size, j+size {
		cipher.Encrypt(encrypted[i:j], data[i:j])
	}

	return encrypted
}

package main

import (
	"crypto/aes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"math/big"
)

func main() {
	oracle := encryptionOracle()
	// block size
	var s int
	for size := 2; size < 256; size++ {
		enc := oracle(make([]byte, 1000))
		if detectEcb(enc, size) {
			s = size
			break
		}
	}
	// our input 		 | text
	// A A A A A ... A A | x U z P [ D c
	// remove 1 byte
	// A A A A A ... A x | U z P [ D c
	{
		m := make(map[string]byte)
		for i := 0; i < 256; i++ {
			b := make([]byte, s)
			b[s-1] = byte(i)
			enc := oracle(b)
			m[string(enc[:s])] = byte(i)
		}

		b := make([]byte, s-1)
		enc := oracle(b)
		var res []byte
		if v, ok := m[string(enc[:s])]; ok {
			res = append(res, v)
		}
	}
	// our input 		 | text
	// A A A A A ... x A | x U z P [ D c
	// remove 2 bytes
	// A A A A A ... x U | z P [ D c

	m := make(map[string]byte)
	var res []byte
	for {
		off := len(res) / s
		rofl := s - (len(res) % s) - 1

		for i := 0; i < 256; i++ {
			b := make([]byte, rofl) // 16
			b = append(b, res...)
			b = append(b, byte(i))
			enc := oracle(b)
			m[string(enc[off*s:(off+1)*s])] = byte(i)
		}

		b := make([]byte, rofl)

		enc := oracle(b)
		if v, ok := m[string(enc[off*s:(off+1)*s])]; ok {
			res = append(res, v)
		} else {
			break
		}
	}
	fmt.Println(string(res))
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func detectEcb(enc []byte, size int) bool {
	for i := size; i < len(enc); i += size {
		if i+size > len(enc) {
			return false
		}
		if cmp(enc[i:i+size], enc[i-size:i]) {
			return true
		}
	}

	return false
}

func cmp(b1, b2 []byte) bool {
	for i := range b1 {
		if b1[i] != b2[i] {
			return false
		}
	}
	return true
}

func encryptionOracle() func(src []byte) []byte {
	size := 16
	key := randBytes(size)
	return func(src []byte) []byte {

		//src = randEdges(src)
		b, _ := base64.StdEncoding.DecodeString("Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkg\naGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBq\ndXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUg\nYnkK")
		src = append(src, b...)
		src = padPKCS7(src, size)

		return EncryptAes128Ecb(src, key)
	}
}

func randEdges(src []byte) []byte {
	s, _ := rand.Int(rand.Reader, big.NewInt(6)) // [0,6)
	e, _ := rand.Int(rand.Reader, big.NewInt(6))
	start := make([]byte, s.Int64()+5) // [5, 10]
	end := make([]byte, e.Int64()+5)   // [5, 10]

	rand.Read(start)
	rand.Read(end)

	return append(start, append(src, end...)...)
}

func randBytes(len int) []byte {
	b := make([]byte, len)
	rand.Reader.Read(b)
	return b
}
func padPKCS7(b []byte, alignment int) []byte {
	r := alignment - (len(b) % alignment)
	for i := 0; i < r; i++ {
		b = append(b, byte(r))
	}
	return b
}

func EncryptAes128Cbc(iv, data, key []byte) []byte {
	data = padPKCS7(data, len(key))
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

func EncryptAes128Ecb(data, key []byte) []byte {
	data = padPKCS7(data, len(key))
	cipher, _ := aes.NewCipher(key)
	encrypted := make([]byte, len(data))
	size := 16

	for i, j := 0, size; i < len(data); i, j = i+size, j+size {
		cipher.Encrypt(encrypted[i:j], data[i:j])
	}

	return encrypted
}

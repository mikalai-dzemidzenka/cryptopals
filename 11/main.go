package main

import (
	"crypto/aes"
	"crypto/rand"
	"fmt"
	"math/big"
)

var expEcb, expCbc int

func main() {

	oracle := encryptionOracle()
	var ecb, cbc int
	in := make([]byte, 100)
	for i := 0; i < 1000; i++ {
		enc := oracle(in)
		if detectEcb(enc) {
			ecb++
		} else {
			cbc++
		}
	}
	fmt.Println("expected: ", expEcb, expCbc)
	fmt.Println("actual: ", ecb, cbc)
}

func detectEcb(enc []byte) bool {
	size := 16
	for i := size; i < len(enc); i += size {
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
		src = randEdges(src)
		src = pad(src, size)

		alg, _ := rand.Int(rand.Reader, big.NewInt(2))
		if alg.Int64() == 0 {
			expEcb++
			return EncryptAes128Ecb(src, key)
		} else {
			expCbc++
			iv := randBytes(size)
			return EncryptAes128Cbc(iv, src, key)
		}
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
func pad(b []byte, alignment int) []byte {
	r := alignment - (len(b) % alignment)
	for i := 0; i < r; i++ {
		b = append(b, byte(r))
	}
	return b
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

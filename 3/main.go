package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	ans := sXor("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	fmt.Println(string(ans))
}

func sXor(in string) []byte {
	m := make(map[rune]float64)
	file, err := os.ReadFile("3/pg19033.txt")
	if err != nil {
		panic(err)
	}
	for _, c := range string(file) {
		m[c]++
	}
	total := utf8.RuneCountInString(string(file))
	for i := range m {
		m[i] = m[i] / float64(total)
	}

	b, _ := hex.DecodeString(in)
	var maxmax float64
	var resres []byte
	for key := 0; key < 256; key++ {
		res := make([]byte, len(b))
		for i := range b {
			res[i] = b[i] ^ byte(key)
		}
		var rofl float64
		for _, c := range string(res) {
			rofl += m[c]
		}
		if rofl > maxmax {
			maxmax = rofl
			resres = res
		}
	}

	return resres
}

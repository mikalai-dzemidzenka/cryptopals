package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
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

	file2, err := os.ReadFile("4/4.txt")
	if err != nil {
		panic(err)
	}
	strstr := strings.Split(string(file2), "\n")
	for _, str := range strstr {
		b, _ := hex.DecodeString(str)
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

		if maxmax > 1.0 {
			fmt.Println(maxmax, string(resres))
		}
	}

}

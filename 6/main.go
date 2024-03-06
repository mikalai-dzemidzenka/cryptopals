package main

import (
	"encoding/base64"
	"fmt"
	"math"
	"os"
	"unicode/utf8"
)

// TODO rewrite from scratch
func main() {
	file, err := os.ReadFile("6/6.txt")
	if err != nil {
		panic(err)
	}

	f2 := make([]byte, len(file))
	_, err = base64.StdEncoding.Decode(f2, file)
	if err != nil {
		panic(err)
	}
	file = f2

	var min = math.MaxFloat64
	var keyLen int
	for i := 2; i < 40; i++ {
		bsize := i * 4
		dist := float64(h2(file[0:bsize], file[bsize:bsize*2])) / float64(i)
		if dist < min {
			keyLen = i
			min = dist
		}
	}
	fmt.Println(keyLen)

	blocks := make([][]byte, keyLen)
	for i := range blocks {
		blocks[i] = make([]byte, 0, len(file)/keyLen)
	}

	for i := range file {
		blocks[i%keyLen] = append(blocks[i%keyLen], file[i])
	}

	key := make([]byte, keyLen)
	for i := range blocks {
		_, char, _ := findSingleXORKey(blocks[i], metric())
		key[i] = char
	}
	fmt.Println(string(key), "\n")

	res := make([]byte, len(file))
	for i := range file {
		res[i] = file[i] ^ key[i%keyLen]
	}

	fmt.Println(string(res))
}

func h1(b1, b2 byte) int {
	c := 0
	for i := 0; i < 8; i++ {
		c1 := (b1 >> i) & 1
		c2 := (b2 >> i) & 1
		if c1 != c2 {
			c++
		}
	}
	return c
}

func h2(s1, s2 []byte) int {
	if len(s1) != len(s2) {
		return 0
	}
	var c int
	for i := range s1 {
		c += h1(s1[i], s2[i])
	}
	return c
}

func sXor(b []byte) ([]byte, byte) {
	m := metric()

	var score float64
	var res []byte
	var char byte
	for key := 0; key < 256; key++ {
		text := make([]byte, len(b))
		for i := range b {
			text[i] = b[i] ^ byte(key)
		}
		var s float64
		for _, c := range string(text) {
			s += m[c]
		}
		if s > score {
			score = s
			res = text
			char = byte(key)
		}
	}

	return res, char
}

func metric() map[rune]float64 {
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
	return m
}

// filippo

func scoreEnglish(text string, c map[rune]float64) float64 {
	var score float64
	for _, char := range text {
		score += c[char]
	}
	return score / float64(utf8.RuneCountInString(text))
}

func singleXOR(in []byte, key byte) []byte {
	res := make([]byte, len(in))
	for i, c := range in {
		res[i] = c ^ key
	}
	return res
}

func findSingleXORKey(in []byte, c map[rune]float64) (res []byte, key byte, score float64) {
	for k := 0; k < 256; k++ {
		out := singleXOR(in, byte(k))
		s := scoreEnglish(string(out), c)
		if s > score {
			res = out
			score = s
			key = byte(k)
		}
	}
	return
}

func findRepeatingXORKey(in []byte, keySize int, c map[rune]float64) []byte {

	column := make([]byte, (len(in)+keySize-1)/keySize)
	key := make([]byte, keySize)
	for col := 0; col < keySize; col++ {
		for row := range column {
			if row*keySize+col >= len(in) {
				continue
			}
			column[row] = in[row*keySize+col]
		}
		_, k, _ := findSingleXORKey(column, c)
		key[col] = k
	}
	return key
}

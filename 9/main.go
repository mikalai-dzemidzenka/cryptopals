package main

import "fmt"

func main() {
	p := pad([]byte("YELLOW SUBMARINE"), 20)
	fmt.Println(string(p))
}

func pad(b []byte, p int) []byte {
	r := p - len(b)
	for i := 0; i < r; i++ {
		b = append(b, byte(r))
	}
	return b
}

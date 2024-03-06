package main

import (
	"encoding/hex"
	"fmt"
	"strings"
)

func main() {
	in := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	out := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272\na282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
	material := "ICE"

	inb := []byte(in)
	dest := make([]byte, len(inb))
	j := 0
	for i := range dest {
		dest[i] = inb[i] ^ material[j%len(material)]
		j++
	}
	_ = out
	rofl := hex.EncodeToString(dest)
	fmt.Println(rofl)
	fmt.Println("------")
	out = strings.ReplaceAll(out, "\n", "")
	fmt.Println(out)
	fmt.Println(rofl == out)
}

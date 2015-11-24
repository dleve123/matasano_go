package main

import(
	"encoding/hex"
	"fmt"
)

func fixedXor(buffer1 string, buffer2 string) string {
	firstBufferBytes, _ := hex.DecodeString(buffer1)
	secondBufferBytes, _ := hex.DecodeString(buffer2)

	n := len(firstBufferBytes)
	xorResult := make([]byte, n)

	for i := 0; i < n; i++ {
		xorResult[i] = firstBufferBytes[i] ^ secondBufferBytes[i]
	}

	return hex.EncodeToString(xorResult)
}

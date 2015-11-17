package main

import "encoding/hex"
import "encoding/base64"

func convert(hexInput string) string {
	inputBytes, _ := hex.DecodeString(hexInput)
	base64Encoded := base64.StdEncoding.EncodeToString(inputBytes)

	return base64Encoded
}

package main

import "testing"

func TestConvertFromHexToBase64(t *testing.T) {
  hexInput := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
  base64Result := convert(hexInput)

  if base64Result != "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t" {
    t.Error("conversion from hex to base 64 failed!")
  }

}

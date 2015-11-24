package main

import "testing"

func TestXor(t *testing.T) {
	buffer1 := "1c0111001f010100061a024b53535009181c"
	buffer2 := "686974207468652062756c6c277320657965"

	result := fixedXor(buffer1, buffer2)

	expectedValue := "746865206b696420646f6e277420706c6179"
	if result != expectedValue {
		t.Errorf("expected XOR to give %s, got %s", expectedValue, result)
	}
}

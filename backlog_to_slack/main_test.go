package main

import (
	"testing"
)

func TestHandler(t *testing.T) {
	t.Run("Successful make text no ticket", func(t *testing.T) {
		expected, _ := makeText([]byte{91, 93})
		actual := "前営業日に発番されたチケットはありません。"
		if actual != expected {
			t.Errorf("\ngot %v\nwant %v", actual, expected)
		}
	})
	t.Run("Failed make text no ticket", func(t *testing.T) {
		expected, _ := makeText([]byte{})
		actual := "failed"
		if actual != expected {
			t.Errorf("\ngot %v\nwant %v", actual, expected)
		}
	})
}

package ezclassify_test

import (
	"strings"
	"testing"

	"github.com/ezclassify/client-go"
)

func TestUTF8OffsetsRoundTripCanonicalBoundaries(t *testing.T) {
	text := "A\U0001f642e\u0301\r\nZ"
	wantByteOffsets := []int{0, 1, 5, 6, 8, 9, 10, 11}
	for runeIndex, wantByteOffset := range wantByteOffsets {
		byteOffset, err := ezclassify.UTF8ByteOffset(text, runeIndex)
		if err != nil {
			t.Fatalf("UTF8ByteOffset(%d): %v", runeIndex, err)
		}
		if byteOffset != wantByteOffset {
			t.Errorf("UTF8ByteOffset(%d) = %d, want %d", runeIndex, byteOffset, wantByteOffset)
		}
		gotRuneIndex, err := ezclassify.UTF8RuneIndex(text, byteOffset)
		if err != nil {
			t.Fatalf("UTF8RuneIndex(%d): %v", byteOffset, err)
		}
		if gotRuneIndex != runeIndex {
			t.Errorf("UTF8RuneIndex(%d) = %d, want %d", byteOffset, gotRuneIndex, runeIndex)
		}
	}
}

func TestUTF8OffsetsRejectInvalidInput(t *testing.T) {
	invalidUTF8 := string([]byte{0xff})
	tests := []struct {
		name string
		call func() error
		want string
	}{
		{"negative rune index", func() error { _, err := ezclassify.UTF8ByteOffset("x", -1); return err }, "out of range"},
		{"large rune index", func() error { _, err := ezclassify.UTF8ByteOffset("x", 2); return err }, "out of range"},
		{"negative byte offset", func() error { _, err := ezclassify.UTF8RuneIndex("x", -1); return err }, "out of range"},
		{"large byte offset", func() error { _, err := ezclassify.UTF8RuneIndex("x", 2); return err }, "out of range"},
		{"mid-rune byte offset", func() error { _, err := ezclassify.UTF8RuneIndex("\u00e9", 1); return err }, "not a rune boundary"},
		{"invalid UTF-8 to byte", func() error { _, err := ezclassify.UTF8ByteOffset(invalidUTF8, 0); return err }, "not valid UTF-8"},
		{"invalid UTF-8 to rune", func() error { _, err := ezclassify.UTF8RuneIndex(invalidUTF8, 0); return err }, "not valid UTF-8"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.call()
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("error = %v, want containing %q", err, test.want)
			}
		})
	}
}

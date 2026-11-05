package ezclassify_test

import (
	"strings"
	"testing"

	"github.com/ezclassify/sdk-go"
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

// TestUTF8BatchOffsetsAgreeWithTheSingleIndexFunctions is the contract for the
// one-pass converters: same answers as the per-index functions, in the
// caller's own order, for unsorted and duplicated requests alike.
func TestUTF8BatchOffsetsAgreeWithTheSingleIndexFunctions(t *testing.T) {
	text := "A\U0001f642é\r\nZ"
	runeIndexes := []int{7, 0, 3, 3, 1, 6}

	gotOffsets, err := ezclassify.UTF8ByteOffsets(text, runeIndexes)
	if err != nil {
		t.Fatalf("UTF8ByteOffsets: %v", err)
	}
	if len(gotOffsets) != len(runeIndexes) {
		t.Fatalf("len(offsets) = %d, want %d", len(gotOffsets), len(runeIndexes))
	}
	for position, runeIndex := range runeIndexes {
		want, err := ezclassify.UTF8ByteOffset(text, runeIndex)
		if err != nil {
			t.Fatalf("UTF8ByteOffset(%d): %v", runeIndex, err)
		}
		if gotOffsets[position] != want {
			t.Errorf("offsets[%d] for rune %d = %d, want %d", position, runeIndex, gotOffsets[position], want)
		}
	}

	gotIndexes, err := ezclassify.UTF8RuneIndexes(text, gotOffsets)
	if err != nil {
		t.Fatalf("UTF8RuneIndexes: %v", err)
	}
	for position, byteOffset := range gotOffsets {
		want, err := ezclassify.UTF8RuneIndex(text, byteOffset)
		if err != nil {
			t.Fatalf("UTF8RuneIndex(%d): %v", byteOffset, err)
		}
		if gotIndexes[position] != want {
			t.Errorf("indexes[%d] for offset %d = %d, want %d", position, byteOffset, gotIndexes[position], want)
		}
	}
}

// TestUTF8BatchOffsetsRejectInvalidInput pins the failure contract: one bad
// member fails the whole call, so a caller never gets a half-converted set.
func TestUTF8BatchOffsetsRejectInvalidInput(t *testing.T) {
	text := "abé"

	if _, err := ezclassify.UTF8ByteOffsets(text, []int{0, 99}); err == nil {
		t.Error("UTF8ByteOffsets past the end must fail")
	}
	if _, err := ezclassify.UTF8ByteOffsets(text, []int{-1}); err == nil {
		t.Error("UTF8ByteOffsets on a negative index must fail")
	}
	if _, err := ezclassify.UTF8ByteOffsets("\xff", []int{0}); err == nil {
		t.Error("UTF8ByteOffsets on invalid UTF-8 must fail")
	}
	if _, err := ezclassify.UTF8RuneIndexes(text, []int{0, 3}); err == nil {
		t.Error("UTF8RuneIndexes inside a multi-byte rune must fail")
	}
	if _, err := ezclassify.UTF8RuneIndexes(text, []int{99}); err == nil {
		t.Error("UTF8RuneIndexes past the end must fail")
	}
	if _, err := ezclassify.UTF8RuneIndexes("\xff", []int{0}); err == nil {
		t.Error("UTF8RuneIndexes on invalid UTF-8 must fail")
	}
}

// TestUTF8BatchOffsetsHandleTheEmptyAndBoundaryCases covers the
// end-exclusive boundary and the empty request, which the scan loop reaches
// only through its tail.
func TestUTF8BatchOffsetsHandleTheEmptyAndBoundaryCases(t *testing.T) {
	offsets, err := ezclassify.UTF8ByteOffsets("", []int{0})
	if err != nil || len(offsets) != 1 || offsets[0] != 0 {
		t.Fatalf("UTF8ByteOffsets on empty text = %v, %v", offsets, err)
	}
	indexes, err := ezclassify.UTF8RuneIndexes("", []int{0})
	if err != nil || len(indexes) != 1 || indexes[0] != 0 {
		t.Fatalf("UTF8RuneIndexes on empty text = %v, %v", indexes, err)
	}
	text := "éZ"
	offsets, err = ezclassify.UTF8ByteOffsets(text, []int{3})
	if err != nil || offsets[0] != len(text) {
		t.Fatalf("UTF8ByteOffsets at the end boundary = %v, %v", offsets, err)
	}
	indexes, err = ezclassify.UTF8RuneIndexes(text, []int{len(text)})
	if err != nil || indexes[0] != 3 {
		t.Fatalf("UTF8RuneIndexes at the end boundary = %v, %v", indexes, err)
	}
	if offsets, err := ezclassify.UTF8ByteOffsets(text, nil); err != nil || len(offsets) != 0 {
		t.Fatalf("UTF8ByteOffsets on an empty request = %v, %v", offsets, err)
	}
}

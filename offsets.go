package ezclassify

import (
	"fmt"
	"unicode/utf8"
)

// UTF8ByteOffset returns the zero-based UTF-8 byte offset for a zero-based
// rune index in text. The index may equal the rune count to address the
// end-exclusive boundary. Invalid UTF-8 and out-of-range indices return an
// error rather than silently substituting replacement characters.
func UTF8ByteOffset(text string, runeIndex int) (int, error) {
	if !utf8.ValidString(text) {
		return 0, errInvalidUTF8()
	}
	if runeIndex < 0 {
		return 0, errRuneIndexRange(runeIndex)
	}

	index := 0
	for byteOffset := range text {
		if index == runeIndex {
			return byteOffset, nil
		}
		index++
	}
	if index == runeIndex {
		return len(text), nil
	}
	return 0, errRuneIndexRange(runeIndex)
}

// UTF8RuneIndex returns the zero-based rune index for a zero-based UTF-8 byte
// offset in text. The offset may equal len(text) to address the end-exclusive
// boundary. Offsets inside a multi-byte encoding, invalid UTF-8, and
// out-of-range offsets return an error.
func UTF8RuneIndex(text string, byteOffset int) (int, error) {
	if !utf8.ValidString(text) {
		return 0, errInvalidUTF8()
	}
	if byteOffset < 0 || byteOffset > len(text) {
		return 0, errByteOffsetRange(byteOffset)
	}
	if byteOffset < len(text) && !utf8.RuneStart(text[byteOffset]) {
		return 0, errByteOffsetBoundary(byteOffset)
	}
	return utf8.RuneCountInString(text[:byteOffset]), nil
}

func errInvalidUTF8() error {
	return fmt.Errorf("ezclassify: text is not valid UTF-8")
}

func errRuneIndexRange(runeIndex int) error {
	return fmt.Errorf("ezclassify: rune index %d is out of range", runeIndex)
}

func errByteOffsetRange(byteOffset int) error {
	return fmt.Errorf("ezclassify: byte offset %d is out of range", byteOffset)
}

func errByteOffsetBoundary(byteOffset int) error {
	return fmt.Errorf("ezclassify: byte offset %d is not a rune boundary", byteOffset)
}

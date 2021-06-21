package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"testing"
)

func TestGetCells(t *testing.T) {
	data := []byte{}

	for _, c := range []cell{[]byte{1, 2, 3}, []byte{3, 4}} {
		vib := make([]byte, 1)
		binary.PutUvarint(vib, uint64(len(c)+1))
		data = append(data, append(vib, c...)...)
	}

	r := bufio.NewReader(bytes.NewReader(data))
	cells := getCells(r)

	if len(cells) != 2 {
		t.Errorf("Expected cells length of 2, but got %v", len(cells))
	}
}

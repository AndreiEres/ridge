package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"io"
	"log"
)

type cell []byte

func getCells(r *bufio.Reader) []cell {
	cells := []cell{}
	vib := make([]byte, 1)

	for {
		n, err := r.Read(vib)
		if n == 0 && err == io.EOF {
			break
		}

		vi, _ := binary.ReadUvarint(bytes.NewReader(vib))
		if vi <= 0 {
			break
		}

		buf := make([]byte, vi-1)
		m, err := r.Read(buf)
		if m == 0 && err == io.EOF {
			break
		}

		cells = append(cells, append(vib, buf...))

		if n != 0 && err == nil {
			continue
		}

		log.Fatal(err)
	}

	return cells
}

package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"
)

func main() {
	fileName := "data.bin"

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error to read [file=%v]: %v", fileName, err.Error())
	}

	cells := getCells(bufio.NewReader(f))

	for _, c := range cells {
		b := bytes.Split(c, []byte{0})
		log.Println(len(b))
	}

	vib := make([]byte, 1)
	binary.PutUvarint(vib, uint64(0))
	fmt.Println(">>>", vib)
}

package main

import (
	"fmt"

	"github.com/diemonster/transaction-processor/circular_buffer"
	"github.com/diemonster/transaction-processor/filereader"
	"github.com/diemonster/transaction-processor/processor"
)

func main() {
	buffer := circular_buffer.New(10)
	fr := filereader.NewFileReader("data.json")
	lines, err := fr.ReadLines()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	p := processor.New(buffer)
	p.ProcessData(lines)

	// Flush remaining entries in the buffer
	for {
		data, err := buffer.Delete()
		if err != nil {
			break
		}
		fmt.Println(data)
	}
}

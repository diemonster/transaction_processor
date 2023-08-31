package main

import (
	"fmt"

	"github.com/diemonster/transaction-processor/circular_buffer"
	"github.com/diemonster/transaction-processor/filereader"
	"github.com/diemonster/transaction-processor/processor"
)

func main() {
	buffer := circular_buffer.NewCircularBuffer(10)
	fr := filereader.NewFileReader("data.json")
	lines, err := fr.ReadLines()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	p := processor.NewProcessor(buffer)
	p.ProcessData(lines)

	// Flush remaining entries in the buffer
	for {
		if data, err := buffer.Delete(); err == nil {
			fmt.Println(data)
		} else {
			break
		}
	}
}

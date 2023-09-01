package processor

import (
	"encoding/json"
	"fmt"

	"github.com/diemonster/transaction-processor/circular_buffer"
	"github.com/diemonster/transaction-processor/data"
)

type Processor struct {
	buffer circular_buffer.Buffer
}

func New(buffer circular_buffer.Buffer) *Processor {
	return &Processor{buffer: buffer}
}

func (p *Processor) ProcessData(input <-chan string) {
	var entry data.Entry
	for line := range input {
		if err := json.Unmarshal([]byte(line), &entry); err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			continue
		}

		if err := p.buffer.Add(entry); err == nil {
			continue
		}

		for {
			data, err := p.buffer.Delete()
			if err != nil {
				break
			}
			fmt.Println(data)
		}
	}
}

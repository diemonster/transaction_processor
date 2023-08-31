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

func NewProcessor(buffer circular_buffer.Buffer) *Processor {
	return &Processor{buffer: buffer}
}

func (p *Processor) ProcessData(input <-chan string) {
	for line := range input {
		entry := data.Entry{}
		err := json.Unmarshal([]byte(line), &entry)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			continue
		}

		if err := p.buffer.Add(entry); err != nil {
			for {
				if data, err := p.buffer.Delete(); err == nil {
					fmt.Println(data)
				} else {
					break
				}
			}
		}
	}
}

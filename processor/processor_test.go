package processor

import (
	"testing"

	"github.com/diemonster/transaction-processor/circular_buffer"
	"github.com/diemonster/transaction-processor/data"
)

func TestProcessor_ProcessData(t *testing.T) {
	input := make(chan string)
	go func() {
		input <- `{"datetime":"2023-06-27 22:22:19.62710.192501","value":"1","partition":"p5"}`
		input <- `{"datetime":"2023-06-27 22:22:19.62710.193409","value":"2","partition":"p4"}`
		close(input)
	}()

	buffer := circular_buffer.New(5)
	processor := New(buffer)
	processor.ProcessData(input)

	expectedData := []data.Entry{
		{Datetime: "2023-06-27 22:22:19.62710.192501", Value: "1", Partition: "p5"},
		{Datetime: "2023-06-27 22:22:19.62710.193409", Value: "2", Partition: "p4"},
	}
	for i, expectedEntry := range expectedData {
		actualEntry, err := buffer.Delete()
		if err != nil {
			t.Fatalf("unexpected error while deleting from buffer: %v", err)
		}
		if expectedEntry != actualEntry {
			t.Errorf("buffer entry at index %d mismatch: expected %v, got %v", i, expectedEntry, actualEntry)
		}
	}
}

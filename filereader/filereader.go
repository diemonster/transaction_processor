package filereader

import (
	"bufio"
	"os"
)

type FileReader struct {
	FileName string
}

func NewFileReader(fileName string) *FileReader {
	return &FileReader{FileName: fileName}
}

func (fr *FileReader) ReadLines() (<-chan string, error) {
	file, err := os.Open(fr.FileName)
	if err != nil {
		return nil, err
	}

	lines := make(chan string)
	go func() {
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines <- scanner.Text()
		}
		close(lines)
	}()

	return lines, nil
}

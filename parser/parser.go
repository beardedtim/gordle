package parser

import (
	"bufio"
	"os"
)

type Reader struct {
	scanner *bufio.Scanner
	file    *os.File
}

func (reader *Reader) Scan() bool {
	return reader.scanner.Scan()
}

func (reader *Reader) NextLine() string {
	return reader.scanner.Text()
}

func (reader *Reader) Close() error {
	return reader.file.Close()
}

func readFileByLine(path string) (Reader, error) {
	readFile, err := os.Open(path)

	if err != nil {
		return Reader{}, err
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	return Reader{
		scanner: fileScanner,
		file:    readFile,
	}, nil
}

func ParseFile(filePath string) ([]string, error) {
	reader, err := readFileByLine(filePath)
	list := make([]string, 14855)

	if err != nil {
		return nil, err
	}

	count := 0
	for reader.Scan() {
		line := reader.NextLine()
		list[count] = line
		count++
	}

	reader.Close()

	return list, nil
}

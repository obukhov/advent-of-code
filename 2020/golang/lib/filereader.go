package lib

import (
	"bufio"
	"log"
	"os"
	"sync"
)

type ReadFileCallback func(line string)

// based on https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
func ReadFileWg(name string, callback ReadFileCallback, wg *sync.WaitGroup) {
	ReadFile(name, callback, func() {wg.Done()})
}

func ReadFile(name string, callback ReadFileCallback, completed func()) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		callback(scanner.Text())
	}

	completed()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func ReadToSlice(name string) []string {
	s := make([]string, 0)
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		s = append(s, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return s
}
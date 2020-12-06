package common

import (
	"bufio"
	"log"
	"os"
	"sync"
)

type ReadFileCallback func(line string)

// based on https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
func ReadFile(name string, callback ReadFileCallback, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		callback(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

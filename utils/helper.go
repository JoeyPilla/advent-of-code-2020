package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadTextFile(fileName string) []string {
	resp := []string{}
	f, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {

		resp = append(resp, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return resp
}

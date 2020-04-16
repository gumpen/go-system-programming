package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	records := [][]string{
		{"name", "score"},
		{"bob", "13"},
		{"john", "34"},
		{"tanaka", "2"},
	}

	w := csv.NewWriter(os.Stdout)

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln(err)
		}
	}

	w.Flush()
}

package main

import (
	"fmt"
	"indexer/internals/reader"
	"log"
)

func main() {
	reader.ReadDocument("1362dd7384ca37d2656eb3d38bdd7eb814eed437adc3711ccee0ac3d2a1cec1c.json")

	data, err := reader.ReadWordFile("stop_words_english.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}

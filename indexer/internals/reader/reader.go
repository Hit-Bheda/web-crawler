package reader

import (
	"encoding/json"
	"indexer/internals/types"
	"log"
	"os"
)

func ReadDocument(name string) (types.Document, error) {
	bytes, err := os.ReadFile(name)
	if err != nil {
		log.Fatal("Failed to read the file", err)
	}

	var data []types.Document

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return types.Document{}, err
	}
	return data[0], nil
}

func ReadWordFile(name string) ([]string, error) {
	bytes, err := os.ReadFile(name)
	if err != nil {
		log.Fatal("Failed to read the file", err)
	}

	var data []string

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil

}

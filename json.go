package main

import (
	"encoding/json"
	"io"
)

func decodeJsonFile(f io.Reader) ([]person, error) {

	decoder := json.NewDecoder(f)
	var people []person

	err := decoder.Decode(&people)
	if err != nil {
		return nil, err
	}

	return people, nil
}

type person struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

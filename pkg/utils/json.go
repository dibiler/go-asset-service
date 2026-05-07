package utils

import (
	"encoding/json"
	"os"
)

func LoadJSON[T any](path string) ([]T, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var data []T
	err = json.Unmarshal(file, &data)
	return data, err
}

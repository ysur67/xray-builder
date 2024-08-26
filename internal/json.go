package internal

import (
	"encoding/json"
	"io/fs"
	"os"
)

func ReadJson[T any](path string, obj *T) error {
	file, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, obj)
	if err != nil {
		return err
	}
	return nil
}

func WriteToFile[T any](path string, obj *T) error {
	result, _ := json.MarshalIndent(obj, "", "    ")
	result = append(result, '\n')
	return os.WriteFile(path, result, fs.FileMode(0644))
}

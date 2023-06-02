package internal

import (
	"encoding/json"
	"io/ioutil"
)

func ReadJson[T any](path string, obj *T) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(file), obj)
	if err != nil {
		return err
	}
	return nil
}

func WriteToFile[T any](path string, obj *T) error {
	result, _ := json.Marshal(obj)
	return ioutil.WriteFile(path, result, 0644)
}

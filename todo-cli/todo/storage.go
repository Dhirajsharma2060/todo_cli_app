package todo

import (
	"encoding/json"
	"os"
)

func SaveFile(filename string, todos []Todo) error {
	data, err := json.MarshalIndent(todos, "", " ")
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// to make sure that data is loaded before the operation

func LoadFile(filename string) (Todos, error) {
	data, err := os.ReadFile(filename)
	var todos Todos
	if err != nil {
		if os.IsNotExist(err) {
			return todos, nil
		}
		return nil, err

	}
	err = json.Unmarshal(data, &todos)
	return todos, err
}

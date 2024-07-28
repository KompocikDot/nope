package internal

import (
	"encoding/json"
	"os"
	"path"
)

func SaveTodos(todos []todo) {
	dir, err := os.UserCacheDir()
	if err != nil {
		panic(err)
	}

	err = os.Mkdir(path.Join(dir, "nope"), 0755)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}

	filePath := path.Join(dir, "nope", "todos.nope")
	f, err := os.Create(filePath)
	defer f.Close()

	if err != nil {
		panic(err)
	}

	err = json.NewEncoder(f).Encode(todos)
	if err != nil {
		panic(todos)
	}
}

func ReadTodos() []todo {
	f, err := os.Open("todos.nope")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	out := []todo{}
	err = json.NewDecoder(f).Decode(&out)

	return out
}

package internal

import (
	"encoding/json"
	"io"
	"os"
	"path"
)

func getOrCreateTodosPath() string {
	dir, err := os.UserCacheDir()
	if err != nil {
		panic(err)
	}

	err = os.Mkdir(path.Join(dir, "nope"), 0755)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}

	filePath := path.Join(dir, "nope", "todos.nope")

	return filePath
}

func SaveTodos(todos []todo) {
	filePath := getOrCreateTodosPath()
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

func ReadTodos() *[]todo {
	filePath := getOrCreateTodosPath()
	f, err := os.Open(filePath)
	defer f.Close()

	if err != nil {
		panic(err)
	}

	out := new([]todo)

	err = json.NewDecoder(f).Decode(out)
	if err != nil && err != io.EOF {
		panic(err)
	}
	return out
}

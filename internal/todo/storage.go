package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

//  # Read/Write logic for interacting with the filesystem

const fileName = "todo.json"

func SaveTaskToFiles() {
	file , err := os.Create(fileName)
	if err!=nil {
		fmt.Println("error creating file:", err)
		return
	}

	defer file.Close()

	jsonData, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		fmt.Println("error marshalling Tasks:", err)
		return
	}
	file.Write(jsonData)
	fmt.Println("Task saved to", fileName);
}

func LoadTasksFromFile() {
	file , err := os.Open(fileName)
	if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

	jsonParser := json.NewDecoder(file)
	err = jsonParser.Decode(&tasks)
	if err != nil {
        fmt.Println("Error decoding tasks:", err)
    }
}
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

	// jsonData, err := json.MarshalIndent(tasks, "", " ")
	encoder := json.NewEncoder(file)
	encoder.SetIndent(""," ")

	err = encoder.Encode(tasks)
	if err != nil {
		fmt.Println("error marshalling Tasks:", err)
		return
	}
	fmt.Println("Task saved to", fileName);
}

func LoadTasksFromFile() {
	// Check if file exists; if not, initialize an empty task list
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		fmt.Println("No task file found, starting with an empty task list.")
		tasks = []Task{} // Initialize an empty task slice
		return
	}

	// Open the file for reading
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Decode the tasks from the file
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		fmt.Println("Error decoding tasks, file might be corrupted:", err)
		tasks = []Task{} // Initialize an empty task slice if decoding fails
		return
	}

	fmt.Println("Tasks loaded from", fileName)
}

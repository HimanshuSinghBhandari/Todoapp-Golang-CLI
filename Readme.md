# Todo CLI Application

A simple and efficient command-line interface (CLI) for managing your tasks. Built using **Go** and **Cobra**, this app allows you to add, delete, list, and display tasks in a tabular format, while persisting data to the file system.

## Features

- **Add tasks**: Add new tasks with a description.
- **Delete tasks**: Remove tasks by their ID.
- **List tasks**: View all tasks in a simple list.
- **Display tasks in a table**: Display tasks in a formatted table.
- **Save and load tasks**: Tasks persist via a JSON file (`todo.json`).

## Project Structure

todo-cli/ │ ├── cmd/ │ ├── add.go # Add task command │ ├── delete.go # Delete task command │ ├── list.go # List tasks command │ ├── save.go # Save tasks command │ ├── table.go # Display tasks in table command │ ├── internal/ │ └── todo/ │ ├── todo.go # Task operations (add, delete, display) │ └── storage.go # Read/write logic for saving/loading tasks │ ├── go.mod # Go module file ├── main.go # Main application entry point └── README.md # Documentation


## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/yourusername/todo-cli.git
    ```

2. Navigate to the project directory:
    ```bash
    cd todo-cli
    ```

3. Initialize Go modules:
    ```bash
    go mod init todo-cli
    ```

4. Install dependencies:
    ```bash
    go get -u github.com/spf13/cobra
    ```

## Usage

### Add a Task
```bash
go run main.go add "Your task description here"

## Delete a Task
go run main.go delete <task_id>

## List All Tasks
go run main.go list

## Display Task in a Table
go run main.go table

## Save task to files
go run main.go save

# Load Tasks from File
Tasks are automatically loaded when you use commands like list or table.

# License
This project is licensed under the MIT License.


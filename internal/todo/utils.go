package todo

// # Utility functions (tabular data, formatting, etc.)


import (
    "fmt"
)

func PrintSuccess(message string) {
    fmt.Println("[SUCCESS]:", message)
}

func PrintError(message string) {
    fmt.Println("[ERROR]:", message)
}

// # Logic for Todo operations (Add, Delete, etc.)

package todo

import "fmt"

type Task struct {
	ID int
	Description string
	Completed bool
}

var tasks []Task
var taskIDCounter = 1

func AddTask(description string){
	task:= Task{
		ID: taskIDCounter,
		Description: description,
		Completed: false,
	}
	tasks = append(tasks, task)
	taskIDCounter++

	SaveTaskToFiles() // save task after adding
}

func DeleteTask(id int){
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...) // slice up means select all element till i (but not include it ) and then select all element from i+1 till end , and ... is variadic which allows yout o pass the element as individual argument
			SaveTaskToFiles()
			return
		}  
	}
    fmt.Println("Task not found"); 
}

func ListTasks() {
	LoadTasksFromFile()
	if len(tasks) == 0 {
		fmt.Println("No task found");
		return
	}
    
	for _ , task := range tasks { // The _ known as Blank Identifier ,  It is used when you need to ignore a value that is returned but you don't want to use it.
		status := "Incomplete"
		if task.Completed {
			status = "completed"
		}
		fmt.Printf("%d: %s [%s] \n",task.ID,task.Description,status);
	}
}

func DisplayTaskTable() {

	LoadTasksFromFile()
	if len(tasks) == 0 {
		fmt.Println("No task dispaly");
		return
	}
	fmt.Println("ID\tTask\tStatus")
	fmt.Println("--------------------------")
	for _, task := range tasks {
		status:= "Incomplete"
		if task.Completed {
			status = "Complete"
		}
		fmt.Printf("%d\t%s\t\t%s\n", task.ID, task.Description, status)
	}
}

func SaveAndLoadtasks() {
	fmt.Println("saving tasks...")
	SaveTaskToFiles()
	fmt.Println("Loading task...")
	LoadTasksFromFile()

	// Display loaded tasks to verify
    fmt.Println("Loaded tasks from file:")
    for _, task := range tasks {
        status := "Incomplete"
        if task.Completed {
            status = "Complete"
        }
        fmt.Printf("%d: %s [%s]\n", task.ID, task.Description, status)
    }
}
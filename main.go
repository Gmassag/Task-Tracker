package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Task represents a single task in the task tracker
type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

// TaskList represents a list of tasks in the task tracker
type TaskList struct {
	Tasks []Task `json:"tasks"`
}

const filename = "tasks.json"

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		handleAdd()
	case "list":
		handleList()
	case "update":
		handleUpdate()
	case "delete":
		handleDelete()
	case "mark-done":
		handleMarkDone()
	case "mark-in-progress":
		handleMarkInProgress()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
	}
}

// printUsage displays the usage instructions for the task tracker CLI
func printUsage() {
	fmt.Println("Task Tracker CLI")
	fmt.Println("Usage:")
	fmt.Println("  task-cli add \"description task\"")
	fmt.Println("  task-cli list [done|todo|in-progress]")
	fmt.Println("  task-cli update <id> \"new description\"")
	fmt.Println("  task-cli delete <id>")
	fmt.Println("  task-cli mark-done <id>")
	fmt.Println("  task-cli mark-in-progress <id>")
}

// load all the task from the json
func loadTasks() (TaskList, error) {
	var taskList TaskList
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return taskList, nil
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		return taskList, err
	}

	err = json.Unmarshal(data, &taskList)
	return taskList, err
}

// saves all tasks to the json file
func saveTasks(taskList TaskList) error {
	data, err := json.MarshalIndent(taskList, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

// generates a new unique ID for a task
func generateID(tasks []Task) int {
	maxID := 0
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	return maxID + 1
}

// add a new task to the list
func handleAdd() {	
	if len(os.Args) < 3 {
		fmt.Println("Please provide a description for the task.")
		return
	}

	description := os.Args[2];

	taskList, err := loadTasks()
	if err != nil {
		fmt.Printf("Error during tasks loading: %v\n", err)
		return
	}

	now := time.Now().Format(time.RFC3339)

	newTask := Task{
		ID:          generateID(taskList.Tasks),
		Description: description,
		Status:      "todo",
		CreatedAt:   now,
		UpdatedAt:   now,
	}
		taskList.Tasks = append(taskList.Tasks, newTask)

	err = saveTasks(taskList)
	if err != nil {
		fmt.Printf("Error while saving: %v\n", err)
		return
	}

	fmt.Printf("Task added succesfully (ID: %d)\n", newTask.ID)
}

// set the status from the "status" parameter and returns the task
func setTaskStatus(id int, status string) *Task{
	taskList, err := loadTasks();
	if err != nil {
		fmt.Printf("Error during tasks's loading: %v\n", err);
	}

	for i := range taskList.Tasks {
		if taskList.Tasks[i].ID == id {
			taskList.Tasks[i].Status = getStatusIcon(status);
			taskList.Tasks[i].UpdatedAt = time.Now().Format(time.RFC3339);

			if err := saveTasks(taskList); err != nil {
				fmt.Printf("Error saving tasks: %v\n", err)
			}


			return &taskList.Tasks[i];
		}
	}

	fmt.Printf("No task found with ID: %d", id);
	return nil;
}

// print all the task as a list 
func handleList() {
	taskList, err := loadTasks()
	if err != nil {
		fmt.Printf("Error during tasks's loading: %v\n", err)
		return
	}

	if len(taskList.Tasks) == 0 {
		fmt.Println("No task found.")
		return
	}

	filter := "all"
	if len(os.Args) >= 3 {
		filter = os.Args[2]
	}

	fmt.Println("Tasks:")
	for _, task := range taskList.Tasks {
		if filter != "all" && task.Status != filter {
			continue
		}

		statusIcon := getStatusIcon(task.Status)
		fmt.Printf("[%d] %s %s - %s\n", task.ID, statusIcon, task.Description, task.Status)
	}
}

// returns the state
func getStatusIcon(status string) string {
	switch status {
	case "todo":
		return "[ ]"
	case "in-progress":
		return "[~]"
	case "done":
		return "[✓]"
	default:
		return "[?]"
	}
}

func handleUpdate() {
	//TODO
}

func handleDelete() {
	//TODO
}

func handleMarkDone() {

	fmt.Printf("w2222");
	if len(os.Args) < 3 {
		fmt.Println("You must input a valid ID.")
		return
	}

	id, err := strconv.ParseInt(os.Args[2], 10, 64); 
	if err != nil {
		fmt.Printf("Input a valid ID");
		return;
	}

	fmt.Printf("test");

	task := setTaskStatus(int(id), "done");
	
	fmt.Println("Task marked as done: %s" +  task.Description);
}

func handleMarkInProgress() {
	//TODO
}
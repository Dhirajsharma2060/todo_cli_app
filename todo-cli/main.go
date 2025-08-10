package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"todo/todo-cli/todo"
)

func UsageFlag() {
	fmt.Println("Usage: todo-cli [command] [arguments]")
	fmt.Println("Commands:")
	fmt.Println("  add <task>       Add a new task")
	fmt.Println("  list             List all tasks")
	fmt.Println("  completed <id>      Change the status of a task")
	fmt.Println("  delete <id>      Delete a task")
}

func main() {
	fmt.Println("Welcome to the Todo CLI!")

	flag.Usage = UsageFlag

	command := os.Args[1]

	//[0]  [1] [2]
	//main.go add/list/completed
	todos := todo.Todos{}
	switch command {
	case "add":
		desciption := os.Args[2]
		todos.Add(desciption)
	case "list":
		todos.List()
	case "completed", "delete":
		ID, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Printf("The argument after the %s should be integer ", command)
		}
		if command == "completed" {
			todos.StatusChange(ID)
		} else {
			todos.Delete(ID)
		}
	}

}

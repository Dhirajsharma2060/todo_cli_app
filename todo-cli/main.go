package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"todo/todo-cli/todo"
)

func UsageFlag() {
	fmt.Println("Welcome to the Todo CLI!")
	fmt.Println("Usage: todo-cli [command] [arguments]")
	fmt.Println("Commands:")
	fmt.Println("  add <task>       Add a new task")
	fmt.Println("  list             List all tasks")
	fmt.Println("  completed <id>      Change the status of a task")
	fmt.Println("  delete <id>      Delete a task")
}

func main() {

	flag.Usage = UsageFlag

	if len(os.Args) > 1 && os.Args[1] == "--help" {
		flag.Usage()
		return
	}

	if len(os.Args) < 2 {
		fmt.Print("Please provide a command.\n")
		flag.Usage()
		return
	}
	command := os.Args[1]
	//[0]  [1] [2]
	//main.go add/list/completed
	// todos := todo.Todos{}
	// this is neccessary to load the data so data is not lost
	todos, err := todo.LoadFile("todo.json")
	if err != nil {
		fmt.Println("Error loading todos:", err)
		return
	}
	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a description for the todo.")
			return

		}
		description := strings.Join(os.Args[2:], " ")
		todos.Add(description)
		todo.SaveFile("todo.json", todos)
	case "list":
		fmt.Println("Welcome to the Todo CLI!")
		fmt.Println("Here are your tasks:")
		todos.List()
	case "completed", "delete":
		ID, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("The argument after the %s should be integer ", command)
		}
		if command == "completed" {
			fmt.Println("Welcome to the Todo CLI!")
			if err := todos.StatusChange(ID); err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("Welcome to the Todo CLI!")
			if err := todos.Delete(ID); err != nil {
				fmt.Println("Error:", err)
				return
			}

		}
		todo.SaveFile("todo.json", todos)
	default:
		fmt.Printf("Unknown command: %s\n", command)
		fmt.Print("Use go run main.go --help to see the list of commands.\n")

	}

}

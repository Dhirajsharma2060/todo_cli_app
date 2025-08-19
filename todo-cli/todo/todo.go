package todo

import (
	"fmt"
	"time"
)

type Todo struct {
	ID          int
	Description string
	Completed   bool
	CreatedAt   time.Time
}

type Todos []Todo

func (todos *Todos) Add(description string) error {

	newID := 1
	if description == "" {
		return fmt.Errorf("Todo description cannot be empty")
	}
	for _, todo := range *todos {
		if todo.ID >= newID {
			newID = todo.ID + 1
		}
	}
	todo := Todo{
		ID:          newID,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
	}

	*todos = append(*todos, todo)
	fmt.Printf("Todo added successfully (ID:%d)\n", todo.ID)
	return nil

}

func (todos *Todos) List() {
	for _, todo := range *todos {
		status := "(Pending)"
		check := ""
		if todo.Completed {
			status = "(Completed)"
			check = "✓ "
		}
		fmt.Printf("[%d] %s%s %s - Created: %s\n",
			todo.ID,
			check,
			todo.Description,
			status,
			todo.CreatedAt.Format("2006-01-02"),
		)
	}

}

func (todos *Todos) Delete(id int) error {
	for i, todo := range *todos {
		if todo.ID == id {
			// note the ... here is a variadic argument that allows use to pass a slice as indivodual argument
			*todos = append((*todos)[:i], (*todos)[i+1:]...)
			fmt.Printf("Todo %d deleted successfully\n", id)
			return nil

		}
	}
	return fmt.Errorf("todo with ID %d not found", id)

}

func (todos *Todos) StatusChange(id int) error {
	for i, todo := range *todos {
		if todo.ID == id {
			(*todos)[i].Completed = true
			fmt.Printf("Todo %d marked as completed\n", id)
			return nil
		}
	}
	return fmt.Errorf("todo with ID %d not found", id)
}

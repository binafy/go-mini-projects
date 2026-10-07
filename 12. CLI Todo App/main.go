package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

// Tasks are saved here so they survive between runs.
const dataFile = "tasks.json"

// Task => represent task model
type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:"created_at"`
}

// go run main.go -add="Buy milk" -description="2 liters"
// go run main.go -list
// go run main.go -done=1
// go run main.go -delete=1
func main() {
	add := flag.String("add", "", "Title of a new task")
	description := flag.String("description", "", "Description of the new task (used with -add)")
	list := flag.Bool("list", false, "List all tasks")
	done := flag.Int("done", 0, "ID of the task to mark as done")
	del := flag.Int("delete", 0, "ID of the task to delete")
	flag.Parse()

	tasks, err := load()
	if err != nil {
		log.Fatal(err)
	}

	switch {
	case *add != "":
		task := Task{
			ID:          nextID(tasks),
			Title:       *add,
			Description: *description,
			CreatedAt:   time.Now(),
		}
		tasks = append(tasks, task)
		fmt.Printf("Added task #%d: %s\n", task.ID, task.Title)

	case *done != 0:
		i := find(tasks, *done)
		if i == -1 {
			log.Fatalf("Task #%d not found", *done)
		}
		tasks[i].Done = true
		fmt.Printf("Completed task #%d: %s\n", tasks[i].ID, tasks[i].Title)

	case *del != 0:
		i := find(tasks, *del)
		if i == -1 {
			log.Fatalf("Task #%d not found", *del)
		}
		fmt.Printf("Deleted task #%d: %s\n", tasks[i].ID, tasks[i].Title)
		tasks = append(tasks[:i], tasks[i+1:]...)

	case *list:
		printTasks(tasks)
		return

	default:
		flag.Usage()
		os.Exit(1)
	}

	if err := save(tasks); err != nil {
		log.Fatal(err)
	}
}

func printTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks yet. Add one with -add=\"...\"")
		return
	}

	for _, t := range tasks {
		mark := " "
		if t.Done {
			mark = "x"
		}
		fmt.Printf("[%s] #%d %s", mark, t.ID, t.Title)
		if t.Description != "" {
			fmt.Printf(" — %s", t.Description)
		}
		fmt.Printf(" (%s)\n", t.CreatedAt.Format("2006-01-02 15:04"))
	}
}

func load() ([]Task, error) {
	data, err := os.ReadFile(dataFile)
	if errors.Is(err, os.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)

	return tasks, err
}

func save(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(dataFile, data, 0644)
}

func nextID(tasks []Task) int {
	max := 0
	for _, t := range tasks {
		if t.ID > max {
			max = t.ID
		}
	}

	return max + 1
}

func find(tasks []Task, id int) int {
	for i, t := range tasks {
		if t.ID == id {
			return i
		}
	}

	return -1
}

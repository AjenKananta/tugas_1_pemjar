package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("TASKMOOD TRACKER")
	fmt.Println("========================")

	storage := NewStorage("taskmood_data.json")
	manager := NewAppManager()

	data, err := storage.Load()
	if err != nil {
		fmt.Printf("Error loading data: %v\n", err)
		fmt.Println("Starting with empty data...")
	} else {
		manager.data = data
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		showMenu()
		fmt.Print("Choose an option (1-8): ")

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())

		switch input {
		case "1":
			addTask(scanner, manager)
		case "2":
			completeTask(scanner, manager)
		case "3":
			deleteTask(scanner, manager)
		case "4":
			addMood(scanner, manager)
		case "5":
			manager.ShowTasks()
		case "6":
			manager.ShowMoods()
		case "7":
			manager.ShowStats()
		case "8":
			if err := storage.Save(manager.data); err != nil {
				fmt.Printf("Error saving data: %v\n", err)
			} else {
				fmt.Println("Data saved successfully!")
			}
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func showMenu() {
	fmt.Println("\nMAIN MENU:")
	fmt.Println("1. Add Task")
	fmt.Println("2. Complete Task")
	fmt.Println("3. Delete Task")
	fmt.Println("4. Add Mood Entry")
	fmt.Println("5. Show All Tasks")
	fmt.Println("6. Show Mood History")
	fmt.Println("7. Show Statistics")
	fmt.Println("8. Exit and Save")
	fmt.Println("========================")
}

func addTask(scanner *bufio.Scanner, manager *AppManager) {
	fmt.Print("Enter task description: ")
	scanner.Scan()
	description := strings.TrimSpace(scanner.Text())

	if description == "" {
		fmt.Println("Task description cannot be empty")
		return
	}

	fmt.Print("Enter priority (1=Low, 2=Medium, 3=High): ")
	scanner.Scan()
	priorityStr := strings.TrimSpace(scanner.Text())

	priority, err := strconv.Atoi(priorityStr)
	if err != nil || priority < 1 || priority > 3 {
		fmt.Println("Invalid priority. Using Medium (2)")
		priority = 2
	}

	manager.AddTask(description, priority)
}

func completeTask(scanner *bufio.Scanner, manager *AppManager) {
	manager.ShowTasks()

	if len(manager.GetTasks()) == 0 {
		return
	}

	fmt.Print("Enter task ID to complete: ")
	scanner.Scan()
	taskIDStr := strings.TrimSpace(scanner.Text())

	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		fmt.Println("Invalid task ID")
		return
	}

	if err := manager.CompleteTask(taskID); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func deleteTask(scanner *bufio.Scanner, manager *AppManager) {
	manager.ShowTasks()

	if len(manager.GetTasks()) == 0 {
		return
	}

	fmt.Print("Enter task ID to delete: ")
	scanner.Scan()
	taskIDStr := strings.TrimSpace(scanner.Text())

	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		fmt.Println("Invalid task ID")
		return
	}

	if err := manager.DeleteTask(taskID); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func addMood(scanner *bufio.Scanner, manager *AppManager) {
	fmt.Println("Choose your mood:")
	fmt.Println("1. Happy")
	fmt.Println("2. Sad")
	fmt.Println("3. Angry")
	fmt.Println("4. Tired")
	fmt.Println("5. Excited")
	fmt.Println("6. Normal")

	fmt.Print("Enter mood choice (1-6): ")
	scanner.Scan()
	choice := strings.TrimSpace(scanner.Text())

	var mood Mood
	switch choice {
	case "1":
		mood = Happy
	case "2":
		mood = Sad
	case "3":
		mood = Angry
	case "4":
		mood = Tired
	case "5":
		mood = Excited
	case "6":
		mood = Normal
	default:
		fmt.Println("Invalid choice. Using Normal mood")
		mood = Normal
	}

	fmt.Print("Add a note (optional): ")
	scanner.Scan()
	note := strings.TrimSpace(scanner.Text())

	manager.AddMood(mood, note)
}

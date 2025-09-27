package main

import (
	"fmt"
	"sort"
)

type AppManager struct {
	data *AppData
}

func NewAppManager() *AppManager {
	return &AppManager{
		data: &AppData{
			Tasks:      []Task{},
			Moods:      []MoodEntry{},
			LastTaskID: 0,
			LastMoodID: 0,
		},
	}
}

func (am *AppManager) AddTask(description string, priority int) {
	am.data.LastTaskID++
	task := NewTask(description, priority)
	task.ID = am.data.LastTaskID
	am.data.Tasks = append(am.data.Tasks, task)
	fmt.Printf("Task added: %s\n", description)
}

func (am *AppManager) CompleteTask(taskID int) error {
	for i := range am.data.Tasks {
		if am.data.Tasks[i].ID == taskID {
			am.data.Tasks[i].Completed = true
			fmt.Printf("Task completed: %s\n", am.data.Tasks[i].Description)
			return nil
		}
	}
	return fmt.Errorf("Task with ID %d not found", taskID)
}

func (am *AppManager) DeleteTask(taskID int) error {
	for i, task := range am.data.Tasks {
		if task.ID == taskID {
			am.data.Tasks = append(am.data.Tasks[:i], am.data.Tasks[i+1:]...)
			fmt.Printf("Task deleted: %s\n", task.Description)
			return nil
		}
	}
	return fmt.Errorf("Task with ID %d not found", taskID)
}

func (am *AppManager) AddMood(mood Mood, note string) {
	am.data.LastMoodID++
	moodEntry := NewMoodEntry(mood, note)
	moodEntry.ID = am.data.LastMoodID
	am.data.Moods = append(am.data.Moods, moodEntry)
	fmt.Printf("Mood recorded: %s - %s\n", string(mood), note)
}

func (am *AppManager) ShowTasks() {
	if len(am.data.Tasks) == 0 {
		fmt.Println("No tasks yet!")
		return
	}

	fmt.Println("\nYOUR TASKS:")
	fmt.Println("========================")

	// Sort by priority and completion status
	tasks := make([]Task, len(am.data.Tasks))
	copy(tasks, am.data.Tasks)

	sort.Slice(tasks, func(i, j int) bool {
		if tasks[i].Completed != tasks[j].Completed {
			return !tasks[i].Completed // Uncompleted tasks first
		}
		return tasks[i].Priority > tasks[j].Priority // Higher priority first
	})

	for _, task := range tasks {
		task.Print()
	}
	fmt.Println()
}

func (am *AppManager) ShowMoods() {
	if len(am.data.Moods) == 0 {
		fmt.Println("No mood entries yet!")
		return
	}

	fmt.Println("\nMOOD HISTORY:")
	fmt.Println("========================")

	for i := len(am.data.Moods) - 1; i >= 0; i-- {
		am.data.Moods[i].Print()
	}
	fmt.Println()
}

func (am *AppManager) ShowStats() {
	totalTasks := len(am.data.Tasks)
	completedTasks := 0
	for _, task := range am.data.Tasks {
		if task.Completed {
			completedTasks++
		}
	}

	fmt.Println("\nYOUR STATS:")
	fmt.Println("========================")
	fmt.Printf("Total Tasks: %d\n", totalTasks)
	fmt.Printf("Completed: %d\n", completedTasks)
	fmt.Printf("Pending: %d\n", totalTasks-completedTasks)

	if len(am.data.Moods) > 0 {
		fmt.Printf("Mood Entries: %d\n", len(am.data.Moods))
		fmt.Printf("Latest Mood: %s\n", string(am.data.Moods[len(am.data.Moods)-1].Mood))
	}

	if len(am.data.Moods) > 0 {
		moodCount := make(map[Mood]int)
		for _, mood := range am.data.Moods {
			moodCount[mood.Mood]++
		}

		fmt.Println("\nMood Frequency:")
		for mood, count := range moodCount {
			fmt.Printf("- %s: %d times\n", string(mood), count)
		}
	}
	fmt.Println()
}

func (am *AppManager) GetMoods() []MoodEntry {
	return am.data.Moods
}

func (am *AppManager) GetTasks() []Task {
	return am.data.Tasks
}

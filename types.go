package main

import (
	"fmt"
	"time"
)

type Mood string

const (
	Happy   Mood = "Happy"
	Sad     Mood = "Sad"
	Angry   Mood = "Angry"
	Tired   Mood = "Tired"
	Excited Mood = "Excited"
	Normal  Mood = "Normal"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	Priority    int       `json:"priority"` // 1-3: 3 paling penting
}

type MoodEntry struct {
	ID        int       `json:"id"`
	Mood      Mood      `json:"mood"`
	Note      string    `json:"note"`
	CreatedAt time.Time `json:"created_at"`
}

type AppData struct {
	Tasks      []Task      `json:"tasks"`
	Moods      []MoodEntry `json:"moods"`
	LastTaskID int         `json:"last_task_id"`
	LastMoodID int         `json:"last_mood_id"`
}

func NewTask(description string, priority int) Task {
	return Task{
		ID:          0,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
		Priority:    priority,
	}
}

func NewMoodEntry(mood Mood, note string) MoodEntry {
	return MoodEntry{
		ID:        0,
		Mood:      mood,
		Note:      note,
		CreatedAt: time.Now(),
	}
}

func (t Task) Print() {
	status := "[ ]"
	if t.Completed {
		status = "[X]"
	}

	priorityText := ""
	switch t.Priority {
	case 1:
		priorityText = "(Low)"
	case 2:
		priorityText = "(Medium)"
	case 3:
		priorityText = "(High)"
	}

	fmt.Printf("%d. %s %s %s - %s\n",
		t.ID, status, t.Description, priorityText, t.CreatedAt.Format("02 Jan 15:04"))
}

func (m MoodEntry) Print() {
	fmt.Printf("%d. %s: %s - %s\n",
		m.ID, string(m.Mood), m.Note, m.CreatedAt.Format("02 Jan 15:04"))
}

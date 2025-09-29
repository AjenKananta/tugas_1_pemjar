package main

import (
	"encoding/json"
	"os"
)

type Storage struct {
	filename string
}

func NewStorage(filename string) *Storage {
	return &Storage{filename: filename}
}

func (s *Storage) Save(data *AppData) error {
	file, err := os.Create(s.filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

func (s *Storage) Load() (*AppData, error) {
	file, err := os.Open(s.filename)
	if err != nil {
		if os.IsNotExist(err) {
			return &AppData{
				Tasks:      []Task{},
				Moods:      []MoodEntry{},
				LastTaskID: 0,
				LastMoodID: 0,
			}, nil
		}
		return nil, err
	}
	defer file.Close()

	var data AppData
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

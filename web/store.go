package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Store struct {
	Path string
}

func NewStore(filePath string) *Store {
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	return &Store{Path: filePath}
}

func (s Store) GetStudents() ([]*Student, error) {

	file, err := os.Open(s.Path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	students := []*Student{}
	if err := json.NewDecoder(file).Decode(&students); err != nil {
		if err == io.EOF {
			return students, nil
		}
		return nil, err
	}

	return students, nil
}

func (s Store) SaveStudent(student *Student) error {

	student.Presave()
	if err := student.Validate(); err != nil {
		return err
	}

	file, err := os.OpenFile(s.Path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer file.Close()

	students, err := s.GetStudents()
	if err != nil {
		return err
	}
	students = append(students, student)

	if err := json.NewEncoder(file).Encode(students); err != nil {
		return err
	}

	return nil
}

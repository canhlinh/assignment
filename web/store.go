package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"sync"
)

type Store struct {
	Path  string
	mutex *sync.Mutex
}

func NewStore(filePath string) *Store {

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
	}

	return &Store{Path: filePath, mutex: &sync.Mutex{}}
}

func (s Store) save(students []*Student) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	file, err := os.OpenFile(s.Path, os.O_WRONLY|os.O_TRUNC, os.ModeAppend)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	return encoder.Encode(students)
}

func (s Store) GetStudents() ([]*Student, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

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

	students, err := s.GetStudents()
	if err != nil {
		return err
	}
	students = append(students, student)
	return s.save(students)
}

func (s Store) GetStudent(studentID string) (*Student, error) {
	students, err := s.GetStudents()
	if err != nil {
		return nil, err
	}

	for _, student := range students {
		if student.ID == studentID {
			return student, nil
		}
	}

	return nil, errors.New("Student not found")
}

func (s Store) UpdateStudent(student *Student) error {

	if err := student.Validate(); err != nil {
		return err
	}

	students, err := s.GetStudents()
	if err != nil {
		return err
	}

	for i := range students {
		if students[i].ID == student.ID {
			students[i] = student
			return s.save(students)
		}
	}

	return errors.New("Student not found")
}

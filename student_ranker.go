package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

type Student struct {
	Name    string   `json:"name"`
	Class   string   `json:"class"`
	Subject *Subject `json:"subject"`
}

type Students []*Student

type Subject struct {
	Physics int `json:"physics"`
	Math    int `json:"math"`
	History int `json:"history"`
}

func (s *Subject) average() float64 {
	return math.Round(float64(s.Physics+s.Math+s.History)/3*10) / 10
}

func (s *Student) calculateAverageScore() float64 {
	return s.Subject.average()
}

func parseStudentsFromJSON(filepath string) (Students, error) {

	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var students []*Student
	if err := json.NewDecoder(file).Decode(&students); err != nil {
		return nil, err
	}

	return students, nil
}

// Sort sort the students
func (students Students) Sort() {
	sort.SliceStable(students, func(i, j int) bool {
		return students[i].calculateAverageScore() > students[j].calculateAverageScore()
	})
}

// Println print the students
func (students Students) Println() {
	for _, student := range students {
		fmt.Printf("%s - %s - %0.1f - %d - %d - %d\n",
			student.Name, student.Class, student.calculateAverageScore(),
			student.Subject.Math, student.Subject.Physics, student.Subject.History)
	}
}

func main() {

	filepath := "./student_ranker.json"
	if len(os.Args) >= 2 {
		filepath = os.Args[1]
	}

	students, err := parseStudentsFromJSON(filepath)
	if err != nil {
		log.Fatal(err)
	}

	students.Sort()
	students.Println()
}

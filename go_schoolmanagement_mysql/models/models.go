package models

import "fmt"

type (
	// Student ...
	Student struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Class    string `json:"class"`
		Teachers string `json:"teachers"` // TODO use Set
		Comments string `json:"comments"`
	}

	// Teacher ...
	Teacher struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Disciplines string `json:"disciplines"` // TODO use Set
		Comments    string `json:"comments"`
	}

	// Discipline ...
	Discipline struct {
		ID               string `json:"id"`
		Name             string `json:"name"`
		NumberOfLecture  int    `json:"numberoflecture"`
		NumberOfExercise int    `json:"numberofexercise"`
		Comments         string `json:"comments"`
	}

	// Class ....
	Class struct {
		ID       string `json:"id"`
		Comments string `json:"comments"`
	}
)

// ToString ...
func (t Student) ToString() string {
	return fmt.Sprintf("Student[ID: %s, Name: %s, Class ID: %s, Teachers: %s, Comments: %s]", t.ID, t.Name, t.Class, t.Teachers, t.Comments)
}

// ToString ...
func (t Teacher) ToString() string {
	return fmt.Sprintf("Teacher[ID: %s, Name: %s, Disciplines : %s, Comments: %s]", t.ID, t.Name, t.Disciplines, t.Comments)
}

// ToString ...
func (t Discipline) ToString() string {
	return fmt.Sprintf("Discipline[ID: %s, Name: %s, Number of lecture : %d, Number of exercise: %d, Commments: %s]", t.ID, t.Name, t.NumberOfLecture, t.NumberOfExercise, t.Comments)
}

// ToString ...
func (t Class) ToString() string {
	return fmt.Sprintf("Class[ID: %s, Name: %s]", t.ID, t.Comments)
}

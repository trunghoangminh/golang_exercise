package models

import "fmt"

type (
	// Student ...
	Student struct {
		ID       string   `bson:"id"`
		Name     string   `bson:"name"`
		Class    string   `bson:"class"`
		Teachers []string `bson:"teachers"` // TODO use Set
		Comments []string `bson:"comments"`
	}

	// Teacher ...
	Teacher struct {
		ID          string   `bson:"id"`
		Name        string   `bson:"name"`
		Disciplines []string `bson:"disciplines"` // TODO use Set
		Comments    []string `bson:"comments"`
	}

	// Discipline ...
	Discipline struct {
		ID               string   `bson:"id"`
		Name             string   `bson:"name"`
		NumberOfLecture  int      `bson:"numberoflecture"`
		NumberOfExercise int      `bson:"numberofexercise"`
		Comments         []string `bson:"comments"`
	}

	// Class ....
	Class struct {
		Identifier string   `bson:"id"`
		Comments   []string `bson:"comments"`
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
	return fmt.Sprintf("Class[ID: %s, Name: %s]", t.Identifier, t.Comments)
}

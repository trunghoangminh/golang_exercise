package main

import (
	"database/sql"
	"log"

	"github.com/trunghoangminh/schoolmanagement/database"
	"github.com/trunghoangminh/schoolmanagement/models"
)

// Student repository interface
type IStudentRepository interface {
	FindAll() ([]models.Student, error)
	Add(student models.Student) error
	Update(student models.Student) error
	Delete(id string) error
}

// Student repository
type StudentRepository struct {
	db *sql.DB
}

// Find all student
func (studentRepos StudentRepository) FindAll() ([]models.Student, error) {
	res, _ := studentRepos.db.Query("SELECT * FROM Student")
	students := []models.Student{}
	defer res.Close()
	for res.Next() {

		var student models.Student
		res.Scan(&student.ID, &student.Name, &student.Class, &student.Teachers, &student.Comments)
		students = append(students, student)
	}
	return students, nil
}

// Add new Student
func (studentRepos StudentRepository) Add(student models.Student) error {
	_, err := studentRepos.db.Query("INSERT INTO Student VALUES (?,?,?,?,?)", student.ID, student.Name, student.Class, student.Teachers, student.Comments)
	return err
}

// Update Student
func (studentRepos StudentRepository) Update(student models.Student) error {
	_, err := studentRepos.db.Query("UPDATE Student SET id=?, name=?,class=?,teachers=?,comments=?)", student.ID, student.Name, student.Class, student.Teachers, student.Teachers)
	return err
}

// Delete student
func (studentRepos StudentRepository) Delete(id string) error {
	_, err := studentRepos.db.Query("DELETE  FROM Student WHERE id=?", id)
	return err
}

func main() {
	db, err := database.ConnectMySQLDB()

	if err != nil {
		log.Fatal(err)
	}
	studentRepository := StudentRepository{db}
	student := models.Student{"51303198", "Hoang Minh Trung", "1350301", []string{"123,456"}, []string{"123", "456"}}
	studentRepository.Add(student)

}

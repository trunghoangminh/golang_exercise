package repositories

import (
	"database/sql"

	"github.com/trunghoangminh/schoolmanagement/models"
)

// Teacher repository interface
type ITeacherRepository interface {
	FindAll() ([]models.Teacher, error)
	Add(teacher models.Teacher) error
	Update(teacher models.Teacher) error
	Delete(id string) error
}

// Teacher repository
type TeacherRepository struct {
	db *sql.DB
}

// Find all teachers
func (teacherRepos TeacherRepository) FindAll() ([]models.Teacher, error) {
	res, _ := teacherRepos.db.Query("SELECT * FROM Teacher")
	teachers := []models.Teacher{}
	defer res.Close()
	for res.Next() {

		var teacher models.Teacher
		res.Scan(&teacher.ID, &teacher.Name, teacher.Disciplines, &teacher.Comments)
		teachers = append(teachers, teacher)
	}
	return teachers, nil
}

// Add new Teacher
func (teacherRepos TeacherRepository) Add(teacher models.Teacher) error {
	_, err := teacherRepos.db.Query("INSERT INTO Teacher VALUES (?,?,?,?)", teacher.ID, teacher.Name, teacher.Disciplines, teacher.Comments)
	return err
}

// Update Teacher
func (teacherRepos TeacherRepository) Update(teacher models.Teacher) error {
	_, err := teacherRepos.db.Query("UPDATE Teacher SET id=?, name=?,disciplines=?comments=?)", teacher.ID, teacher.Name, teacher.Disciplines, teacher.Comments)
	return err
}

// Delete Teacher
func (teacherRepos TeacherRepository) Delete(id string) error {
	_, err := teacherRepos.db.Query("DELETE  FROM Teacher WHERE id=?", id)
	return err
}

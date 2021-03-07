package repositories

import (
	"database/sql"

	"github.com/trunghoangminh/schoolmanagement/models"
)

// Class repository interface
type IClassRepository interface {
	FindAll() ([]models.Class, error)
	Add(class models.Class) error
	Update(class models.Class) error
	Delete(id string) error
}

// Class repository
type ClassRepository struct {
	db *sql.DB
}

// Find all Class
func (classRepos ClassRepository) FindAll() ([]models.Class, error) {
	res, _ := classRepos.db.Query("SELECT * FROM Class")
	classes := []models.Class{}
	defer res.Close()
	for res.Next() {

		var class models.Class
		res.Scan(&class.ID, &class.Comments)
		classes = append(classes, class)
	}
	return classes, nil
}

// Add new Class
func (classRepos ClassRepository) Add(class models.Class) error {
	_, err := classRepos.db.Query("INSERT INTO Class VALUES (?,?)", class.ID, class.Comments)
	return err
}

// Update Class
func (classRepos ClassRepository) Update(class models.Class) error {
	_, err := classRepos.db.Query("UPDATE Class SET id=?,comments=?)", class.ID, class.Comments)
	return err
}

// Delete Class
func (classRepos ClassRepository) Delete(id string) error {
	_, err := classRepos.db.Query("DELETE  FROM Class WHERE id=?", id)
	return err
}

package repositories

import (
	"database/sql"

	"github.com/trunghoangminh/schoolmanagement/models"
)

// Discipline repository interface
type IDisciplineRepository interface {
	FindAll() ([]models.Discipline, error)
	Add(discipline models.Discipline) error
	Update(discipline models.Discipline) error
	Delete(id string) error
}

// Discipline repository
type DisciplineRepository struct {
	db *sql.DB
}

// Find all disciplines
func (disciplineRepos DisciplineRepository) FindAll() ([]models.Discipline, error) {
	res, _ := disciplineRepos.db.Query("SELECT * FROM Discipline")
	disciplines := []models.Discipline{}
	defer res.Close()
	for res.Next() {

		var discipline models.Discipline
		res.Scan(&discipline.ID, &discipline.Name, discipline.NumberOfLecture, &discipline.NumberOfExercise, &discipline.Comments)
		disciplines = append(disciplines, discipline)
	}
	return disciplines, nil
}

// Add new Discipline
func (disciplineRepos DisciplineRepository) Add(discipline models.Discipline) error {
	_, err := disciplineRepos.db.Query("INSERT INTO Discipline VALUES (?,?,?,?,?)", discipline.ID, discipline.Name, discipline.NumberOfLecture, discipline.NumberOfExercise, discipline.Comments)
	return err
}

// Update Discipline
func (disciplineRepos DisciplineRepository) Update(discipline models.Discipline) error {
	_, err := disciplineRepos.db.Query("UPDATE Discipline SET id=?, name=?,numberoflecture=?,numberofexercise=?,comments=?)", discipline.ID, discipline.Name, discipline.NumberOfLecture, discipline.NumberOfExercise, discipline.Comments)
	return err
}

// Delete Discipline
func (disciplineRepos DisciplineRepository) Delete(id string) error {
	_, err := disciplineRepos.db.Query("DELETE  FROM Discipline WHERE id=?", id)
	return err
}

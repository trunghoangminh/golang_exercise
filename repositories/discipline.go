package repositories

import (
	"context"
	"fmt"
	"log"

	"github.com/trunghoangminh/schoolmanagement/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Discipline repository interface
type IDisciplineRepository interface {
	FindAll() ([]models.Discipline, error)
	Add(class models.Discipline) error
	Update(class models.Discipline) error
	Delete(id string) error
}

// Discipline repository struct
type DisciplineRepository struct {
	collection *mongo.Collection
	ctx        context.Context
}

// Find all Discipline
func (disciplineRepos DisciplineRepository) FindAll() ([]models.Discipline, error) {
	cursor, err := disciplineRepos.collection.Find(disciplineRepos.ctx, bson.M{})
	if err != nil {
		fmt.Println("Failed to get all discipline!")
		log.Fatal(err)
	}

	var discipline []models.Discipline
	if err = cursor.All(disciplineRepos.ctx, &discipline); err != nil {
		log.Fatal(err)
	}

	return discipline, nil
}

// Add new Discipline
func (disciplineRepos DisciplineRepository) Add(discipline models.Discipline) error {
	_, err := disciplineRepos.collection.InsertOne(disciplineRepos.ctx, discipline)
	return err
}

// Update Teacher
func (disciplineRepos DisciplineRepository) Update(discipline models.Discipline) error {
	filter := bson.D{{"id", discipline.ID}}
	after := options.After
	returnOpt := options.FindOneAndUpdateOptions{ReturnDocument: &after}
	update := bson.D{{"$set", bson.D{{"name", discipline.Name}, {"numberoflecture", discipline.NumberOfLecture}, {"numberofexercise", discipline.NumberOfExercise}, {"comments", discipline.Comments}}}}
	return disciplineRepos.collection.FindOneAndUpdate(disciplineRepos.ctx, filter, update, &returnOpt).Err()
}

// Delete Teacher
func (disciplineRepos DisciplineRepository) Delete(id string) error {
	return disciplineRepos.collection.FindOneAndDelete(disciplineRepos.ctx, bson.D{{"id", id}}).Err()
}

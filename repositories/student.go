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

// Student repository interface
type Repository interface {
	FindAll() ([]models.Student, error)
	Add(student models.Student) error
	Update(student models.Student) error
	Delete(id string) error
}

// Student repository struct
type StudentRepository struct {
	collection *mongo.Collection
	ctx        context.Context
}

// Find all Student
func (studentRepos StudentRepository) FindAll() ([]models.Student, error) {
	cursor, err := studentRepos.collection.Find(studentRepos.ctx, bson.M{})
	if err != nil {
		fmt.Println("Failed to get all student!")
		log.Fatal(err)
	}

	var student []models.Student
	if err = cursor.All(studentRepos.ctx, &student); err != nil {
		log.Fatal(err)
	}

	return student, nil
}

// Add new Student
func (studentRepos StudentRepository) Add(student models.Student) error {
	_, err := studentRepos.collection.InsertOne(studentRepos.ctx, student)
	return err
}

// Update Student
func (studentRepos StudentRepository) Update(student models.Student) error {
	filter := bson.D{{"id", student.ID}}
	after := options.After
	returnOpt := options.FindOneAndUpdateOptions{ReturnDocument: &after}
	update := bson.D{{"$set", bson.D{{"name", student.Name}, {"class", student.Class}, {"teachers", student.Teachers}, {"comments", student.Comments}}}}
	return studentRepos.collection.FindOneAndUpdate(studentRepos.ctx, filter, update, &returnOpt).Err()
}

// Delete Student
func (studentRepos StudentRepository) Delete(id string) error {
	return studentRepos.collection.FindOneAndDelete(studentRepos.ctx, bson.D{{"id", id}}).Err()
}

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/trunghoangminh/schoolmanagement/database"
	model "github.com/trunghoangminh/schoolmanagement/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Student repository interface
type Repository interface {
	FindAll() ([]model.Student, error)
	Add(student model.Student) error
	Update(student model.Student) error
	Delete(id string) error
}

// Student repositoru struct
type StudentRepository struct {
	collection *mongo.Collection
	ctx        context.Context
}

// Find all Student
func (studentRepos StudentRepository) FindAll() ([]model.Student, error) {
	cursor, err := studentRepos.collection.Find(studentRepos.ctx, bson.M{})
	if err != nil {
		fmt.Println("Failed to get all student!")
		log.Fatal(err)
	}

	var student []model.Student
	if err = cursor.All(studentRepos.ctx, &student); err != nil {
		log.Fatal(err)
	}

	return student, nil
}

// Add new Student
func (studentRepos StudentRepository) Add(student model.Student) error {
	document, err := studentRepos.collection.InsertOne(studentRepos.ctx, student)
	fmt.Print(document)
	return err
}

// Update Student
func (studentRepos StudentRepository) Update(student model.Student) error {
	filter := bson.D{{"id", student.ID}}
	after := options.After
	returnOpt := options.FindOneAndUpdateOptions{ReturnDocument: &after}
	update := bson.D{{"$set", bson.D{{"name", student.Name}}}}
	return studentRepos.collection.FindOneAndUpdate(studentRepos.ctx, filter, update, &returnOpt).Err()
}

// Delete Student
func (studentRepos StudentRepository) Delete(id string) error {
	return studentRepos.collection.FindOneAndDelete(studentRepos.ctx, bson.D{{"id", id}}).Err()
}

func main() {
	connection := database.ConnectMongoDB()
	studentRepository := StudentRepository{connection.Collection("Student"), context.TODO()}
	// students, err := studentRepository.FindAll()
	// if err == nil {
	// 	for _, student := range students {
	// 		fmt.Println(student.ToString())
	// 	}
	// }
	student := model.Student{"5303201", "Hoang Minh TrungS", "1305031", []string{"123", "456"}, []string{"123", "456"}}
	// // studentRepository.Add(student)
	// studentRepository.Delete("5303201")
	studentRepository.Update(student)
}

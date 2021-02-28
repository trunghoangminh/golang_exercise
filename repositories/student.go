package main

import (
	"context"
	"fmt"
	"log"

	"github.com/trunghoangminh/schoolmanagement/database"
	model "github.com/trunghoangminh/schoolmanagement/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	FindAll() ([]model.Student, error)
	Add(student model.Student) error
	Update(student model.Student) error
	Delete(id string) error
}
type StudentRepository struct {
	col *mongo.Collection
	ctx context.Context
}

func (studentRepos StudentRepository) FindAll() ([]model.Student, error) {
	cursor, err := studentRepos.col.Find(studentRepos.ctx, bson.M{})
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

func main() {
	connection := database.ConnectMongoDB()
	studentRepository := StudentRepository{connection.Collection("Student"), context.TODO()}
	students, err := studentRepository.FindAll()
	if err == nil {
		for _, student := range students {
			fmt.Println(student.ToString())
		}
	}
}

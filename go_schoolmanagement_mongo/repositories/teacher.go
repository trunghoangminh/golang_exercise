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

// Teacher repository interface
type ITeacherRepository interface {
	FindAll() ([]models.Teacher, error)
	Add(class models.Teacher) error
	Update(class models.Teacher) error
	Delete(id string) error
}

// Teacher repositoru struct
type TeacherRepository struct {
	collection *mongo.Collection
	ctx        context.Context
}

// Find all Teacher
func (teacherRepos TeacherRepository) FindAll() ([]models.Teacher, error) {
	cursor, err := teacherRepos.collection.Find(teacherRepos.ctx, bson.M{})
	if err != nil {
		fmt.Println("Failed to get all teacher!")
		log.Fatal(err)
	}

	var teacher []models.Teacher
	if err = cursor.All(teacherRepos.ctx, &teacher); err != nil {
		log.Fatal(err)
	}

	return teacher, nil
}

// Add new Teacher
func (teacherRepos TeacherRepository) Add(teacher models.Teacher) error {
	_, err := teacherRepos.collection.InsertOne(teacherRepos.ctx, teacher)
	return err
}

// Update Teacher
func (teacherRepos TeacherRepository) Update(teacher models.Teacher) error {
	filter := bson.D{{"id", teacher.ID}}
	after := options.After
	returnOpt := options.FindOneAndUpdateOptions{ReturnDocument: &after}
	update := bson.D{{"$set", bson.D{{"name", teacher.Name}, {"disciplines", teacher.Disciplines}, {"comments", teacher.Comments}}}}
	return teacherRepos.collection.FindOneAndUpdate(teacherRepos.ctx, filter, update, &returnOpt).Err()
}

// Delete Teacher
func (teacherRepos TeacherRepository) Delete(id string) error {
	return teacherRepos.collection.FindOneAndDelete(teacherRepos.ctx, bson.D{{"id", id}}).Err()
}

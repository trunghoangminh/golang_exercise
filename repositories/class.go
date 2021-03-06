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

// Class repository interface
type IClassRepository interface {
	FindAll() ([]models.Class, error)
	Add(class models.Class) error
	Update(class models.Class) error
	Delete(id string) error
}

// Class repository struct
type ClassRepository struct {
	collection *mongo.Collection
	ctx        context.Context
}

// Find all Class
func (classRepos ClassRepository) FindAll() ([]models.Class, error) {
	cursor, err := classRepos.collection.Find(classRepos.ctx, bson.M{})
	if err != nil {
		fmt.Println("Failed to get all class!")
		log.Fatal(err)
	}

	var class []models.Class
	if err = cursor.All(classRepos.ctx, &class); err != nil {
		log.Fatal(err)
	}

	return class, nil
}

// Add new Class
func (classRepos ClassRepository) Add(class models.Class) error {
	_, err := classRepos.collection.InsertOne(classRepos.ctx, class)
	return err
}

// Update Class
func (classRepos ClassRepository) Update(class models.Class) error {
	filter := bson.D{{"id", class.ID}}
	after := options.After
	returnOpt := options.FindOneAndUpdateOptions{ReturnDocument: &after}
	update := bson.D{{"$set", bson.D{{"comments", class.Comments}}}}
	return classRepos.collection.FindOneAndUpdate(classRepos.ctx, filter, update, &returnOpt).Err()
}

// Delete Class
func (classRepos ClassRepository) Delete(id string) error {
	return classRepos.collection.FindOneAndDelete(classRepos.ctx, bson.D{{"id", id}}).Err()
}

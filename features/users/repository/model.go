package repository

import (
	"context"
	"gomongo/features/users"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserModel struct {
	Name  string `bson:"name"`
	Age   int    `bson:"age"`
	Email string `bson:"email"`
}

type UserQuery struct {
	db         *mongo.Database
	collection string
}

func New(client *mongo.Database, collection string) users.Repository {
	return &UserQuery{
		db:         client,
		collection: collection,
	}
}

func (ad *UserQuery) AddUser(newUser users.Contoh) (users.Contoh, error) {
	var inputData = new(UserModel)
	inputData.Name = newUser.Name
	inputData.Age = newUser.Age
	inputData.Email = newUser.Email

	_, err := ad.db.Collection(ad.collection).InsertOne(context.Background(), inputData)
	if err != nil {
		return users.Contoh{}, err
	}

	return newUser, nil
}

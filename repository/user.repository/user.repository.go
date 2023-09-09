package user_repository

import (
	"context"

	"github.com/api-rest-go/database"
	m "github.com/api-rest-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// User services.
var userCollection = database.GetColllection("users")
var ctx = context.Background()

func Create(user m.User) error {
	var err error
	_, err = userCollection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func Read() (m.Users, error) {
	var users m.Users
	filter := bson.D{}
	cur, err := userCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var user m.User
		err := cur.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func ReadById(userId string) (primitive.M, error) {
	oid, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.M{"_id": oid}
	User := userCollection.FindOne(ctx, filter)
	doc := bson.M{}
	decodeErr := User.Decode(&doc)
	return doc, decodeErr
}

func Update(user m.User, userId string) (primitive.M, error) {
	oid, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.M{"_id": oid}
	update := bson.M{
		"$set": bson.M{
			"name":     user.Name,
			"email":    user.Email,
			"updateat": user.UpdateAt,
		},
	}
	after := options.After
	upsert := true
	otp := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	result := userCollection.FindOneAndUpdate(ctx, filter, update, &otp)
	doc := bson.M{}
	decodeErr := result.Decode(&doc)
	return doc, decodeErr
}

func Delete(userId string) error {
	oid, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.M{
		"_id": oid,
	}
	_, err := userCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

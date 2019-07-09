package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `json:"id" bson:"id"`
	Name     string             `json:"name" bson:"name"`
	Username string             `json:"username" bson:"username"`
	Password string             `json:"password" bson:"password"`
}

const USER_COL = "users"

func (u *User) Insert() error {
	u.Id = primitive.NewObjectID()
	res, err := getDatabase().Collection(USER_COL).InsertOne(context.Background(), u)
	u.Id = res.InsertedID.(primitive.ObjectID)
	return err
}

func (u *User) Delete() error {
	_, err := getDatabase().Collection(USER_COL).DeleteOne(context.Background(), bson.D{{"_id", u.Id}})
	return err
}

func (u *User) Update() error {
	update := bson.M{"$set": bson.M{
		"id":       u.Id,
		"name":     u.Name,
		"username": u.Username,
		"password": u.Password,
	}}
	_, err := getDatabase().Collection(USER_COL).UpdateOne(context.Background(), bson.D{{"_id", u.Id}}, update)
	return err
}

func FindUser(id primitive.ObjectID) (u User, err error) {
	filter := bson.D{{"_id", id}}
	err = getDatabase().Collection(USER_COL).FindOne(context.Background(), filter).Decode(&u)
	return
}

func AllUsers() (users []User, err error) {
	cursor, err := getDatabase().Collection(USER_COL).Find(context.Background(), bson.D{{}})
	if err != nil {
		return
	}
	var u User
	for cursor.Next(context.TODO()) {
		err = cursor.Decode(&u)
		if err != nil {
			return
		}
		users = append(users, u)
	}
	if err = cursor.Err(); err != nil {
		return
	}
	cursor.Close(context.TODO())
	return
}

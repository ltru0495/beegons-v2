package models

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id           primitive.ObjectID `json:"id" bson:"id"`
	Name         string             `json:"name" bson:"name"`
	Username     string             `json:"username" bson:"username"`
	Password     string             `json:"password" bson:"password"`
	HashPassword []byte             `json:"hashpassword" bson:"hashPassword"`
	Role         string             `json:"role" bson:"role"`
}

const USER_COL = "users"

func (u *User) Insert() error {
	u.Id = primitive.NewObjectID()

	hpass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.HashPassword = hpass
	u.Password = ""

	res, err := GetAppDatabase().Collection(USER_COL).InsertOne(context.Background(), u)
	u.Id = res.InsertedID.(primitive.ObjectID)
	return err
}

func Login(user User) (u User, err error) {
	u, err = FindUserByUsername(user.Username)
	if err != nil {
		return
	}
	err = bcrypt.CompareHashAndPassword(u.HashPassword, []byte(user.Password))
	if err != nil {
		u = User{}
	}
	return
}

func (u *User) Delete() error {
	_, err := GetAppDatabase().Collection(USER_COL).DeleteOne(context.Background(), bson.D{{"_id", u.Id}})
	return err
}

func (u *User) Update() error {
	update := bson.M{"$set": bson.M{
		"id":       u.Id,
		"name":     u.Name,
		"username": u.Username,
		"password": u.Password,
	}}
	_, err := GetAppDatabase().Collection(USER_COL).UpdateOne(context.Background(), bson.D{{"_id", u.Id}}, update)
	return err
}

func FindUser(id primitive.ObjectID) (u User, err error) {
	filter := bson.D{{"_id", id}}
	err = GetAppDatabase().Collection(USER_COL).FindOne(context.Background(), filter).Decode(&u)
	return
}

func FindUserByUsername(username string) (u User, err error) {
	filter := bson.D{{"username", username}}
	err = GetAppDatabase().Collection(USER_COL).FindOne(context.Background(), filter).Decode(&u)
	return
}

func AllUsers() (users []User, err error) {
	cursor, err := GetAppDatabase().Collection(USER_COL).Find(context.Background(), bson.D{{}})
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

package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Test struct {
	Id   primitive.ObjectID `json:"_id" bson:"_id"`
	Name string             `json:"name", `
}

const TEST_COL = "tests"

func (t *Test) Insert() error {
	t.Id = primitive.NewObjectID()
	res, err := getDatabase().Collection(TEST_COL).InsertOne(context.Background(), t)
	//fmt.Print(result)
	t.Id = res.InsertedID.(primitive.ObjectID)
	return err
}

func (t *Test) Delete() error {
	_, err := getDatabase().Collection(TEST_COL).DeleteOne(context.Background(), bson.D{{"_id", t.Id}})
	return err
}

func (t *Test) Update() error {
	update := bson.M{"$set": bson.M{
		"name": t.Name,
	}}
	_, err := getDatabase().Collection(TEST_COL).UpdateOne(context.Background(), bson.D{{"_id", t.Id}}, update)
	return err
}

func FindTest(id primitive.ObjectID) (test Test, err error) {
	filter := bson.D{{"_id", id}}
	err = getDatabase().Collection(TEST_COL).FindOne(context.Background(), filter).Decode(&test)
	return
}

func AllTests() (tests []Test, err error) {
	cursor, err := getDatabase().Collection(TEST_COL).Find(context.Background(), bson.D{{}})
	if err != nil {
		return
	}
	var test Test
	for cursor.Next(context.TODO()) {
		err = cursor.Decode(&test)
		if err != nil {
			return
		}
		tests = append(tests, test)
	}
	if err = cursor.Err(); err != nil {
		return
	}
	cursor.Close(context.TODO())
	return
}

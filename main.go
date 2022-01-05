package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

const uri = "mongodb://localhost:27017"

func main() {
	credential := options.Credential{
		Username: "jack",
		Password: "12345678",
	}
	clientOptions := options.Client()
	clientOptions.ApplyURI(uri).SetAuth(credential)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	//fmt.Println("Successfully to connect and pinged")

	// coll := client.Database("db_enigma").Collection("students")
	productColl := client.Database("db_enigma").Collection("products")

	// InsertOne for Students
	//const layout = "2006-01-02"
	//dt, _ := time.Parse(layout, "2022-01-05")
	//newStudent := Student{
	//	Id:       primitive.NewObjectID(),
	//	Name:     "Wilda",
	//	Gender:   "F",
	//	Age:      23,
	//	JoinDate: dt,
	//	IdCard:   "303",
	//	Senior:   false,
	//}
	// InsertOneStudent(ctx, coll, newStudent)

	// FindAllStudent(ctx, coll)

	//fmt.Println("With bson.D")
	//FindStudentByGenderAndAge(ctx, coll, "M", 26)
	//
	//fmt.Println("With Student Struct")
	//FindStudentStructByGenderAndAge(ctx, coll, "M", 26)

	CountProductByCategory(ctx, productColl, "handphone")
}

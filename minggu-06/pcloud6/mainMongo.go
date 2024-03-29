a-tama silakan import package yang dibutuhkan.

package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson"
)

var ctx = context.Background()

type student struct {
    Name  string `bson:"name"`
    Grade int    `bson:"Grade"`
}

func connect() (*mongo.Database, error) {
    clientOptions := options.Client()
    clientOptions.ApplyURI("mongodb://localhost:27017")
    client, err := mongo.NewClient(clientOptions)
    if err != nil {
        return nil, err
    }

    err = client.Connect(ctx)
    if err != nil {
        return nil, err
    }

    return client.Database("belajar_golang"), nil
}

func insert() {
    db, err := connect()
    if err != nil {
        log.Fatal(err.Error())
    }

    _, err = db.Collection("student").InsertOne(ctx, student{"Wick", 2})
    if err != nil {
        log.Fatal(err.Error())
    }

    _, err = db.Collection("student").InsertOne(ctx, student{"Ethan", 2})
    if err != nil {
        log.Fatal(err.Error())
    }

    fmt.Println("Insert success!")
}

func main() {
    insert()
}

func insert() {
    db, err := connect()
    if err != nil {
        log.Fatal(err.Error())
    }

    _, err = db.Collection("student").InsertOne(ctx, student{"Wick", 2})
    if err != nil {
        log.Fatal(err.Error())
    }

    _, err = db.Collection("student").InsertOne(ctx, student{"Ethan", 2})
    if err != nil {
        log.Fatal(err.Error())
    }

    fmt.Println("Insert success!")
}

func main() {
    insert()
}
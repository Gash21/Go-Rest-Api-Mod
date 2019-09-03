package model

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Photo struct {
	Id    int    `json:"Id"`
	Title string `json:"Title"`
	Url   string `json:"Url"`
	Thumb string `json:"ThumbnailUrl"`
}

var Photos []Photo
var client *mongo.Client
var clientOptions *options.ClientOptions
var ctx context.Context
var latihanCollection *mongo.Collection
var hostMongo string
var dbMongo string

func FindPhoto(id int) Photo {
	filter := bson.D{{"id", id}}
	var result Photo
	err := latihanCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println("error:", err)
		return Photo{Id: 0}
	} else {
		return result
	}
}

func CreateNewPhoto(photo Photo) {
	insertResult, err := latihanCollection.InsertOne(context.TODO(), photo)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Inserted a Single Document: ", insertResult.InsertedID)
}

func UpdatePhoto(photo Photo) {
	filter := bson.M{
		"id": photo.Id,
	}
	update := bson.M{
		"$set": bson.M{
			"title": photo.Title,
			"url":   photo.Url,
			"thumb": photo.Thumb,
		},
	}

	updateResult, err := latihanCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}

func populate() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	hostMongo, exists := os.LookupEnv("MONGO_HOST")
	dbMongo, exists1 := os.LookupEnv("MONGO_DB")
	if !exists {
		fmt.Println("Mongo Link Not found")
	}
	if !exists1 {
		fmt.Println("Mongo DB Not found")
	}
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions = options.Client().ApplyURI(hostMongo)
	client, err5 := mongo.Connect(ctx, clientOptions)
	if err5 != nil {
		fmt.Printf("Mongo request failed")
		return
	}
	fmt.Printf("Mongo request success")
	latihanCollection = client.Database(dbMongo).Collection("latihan")
	cur, err := latihanCollection.Find(context.TODO(), bson.D{})
	defer cur.Close(context.TODO())
	if err != nil {

		fmt.Println("error:", err)
		return
	} else {
		count := 0
		for cur.Next(context.TODO()) {
			// create a value into which the single document can be decoded
			var elem Photo
			err := cur.Decode(&elem)
			if err != nil {
				fmt.Println("error:", err)
			} else {
				Photos = append(Photos, elem)
				count += 1
				//fmt.Println(elem.Url)
			}
		}

		if count == 0 {
			response, err := http.Get("https://jsonplaceholder.typicode.com/photos")
			if err != nil {
				fmt.Printf("The HTTP request failed with error %s\n", err)
				return
			} else {
				data, _ := ioutil.ReadAll(response.Body)
				err2 := json.Unmarshal(data, &Photos)

				if err2 != nil {
					fmt.Println("error:", err2)
					return
				}

				Photos = Photos[0:1000]
				var photoInterface []interface{}
				for _, t := range Photos {
					photoInterface = append(photoInterface, t)
				}
				fmt.Printf("Input to array success")
				_, err4 := latihanCollection.InsertMany(context.TODO(), photoInterface)

				if err4 != nil {
					fmt.Println(err4)
				} else {
					fmt.Println("Inserted multiple documents")
				}
			}
		}
		fmt.Println(Photos[0].Url)
		fmt.Println("data ready")
	}
}

func init() {
	populate()
}

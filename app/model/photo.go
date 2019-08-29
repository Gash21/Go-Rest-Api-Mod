package model

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

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

func FindPhoto(id int) Photo {
	// for _, photo := range Photos {
	// 	if photo.Id == id {
	// 		return photo
	// 	}
	// }
	filter := bson.D{{"id", id}}
	var result Photo
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err5 := mongo.Connect(ctx, clientOptions)
	if err5 != nil {
		fmt.Println("error:", err5)
		return Photo{Id: 0}
	}

	collection := client.Database("newsfeed").Collection("latihan")
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println("error:", err)
		return Photo{Id: 0}
	} else {
		return result
	}
}

func CreateNewPhoto(photo Photo) {
	fmt.Println(photo)
	Photos = append(Photos, photo)
	//return article
}

func UpdatePhoto(phot Photo) {
	for ii, photo := range Photos {
		if photo.Id == phot.Id {
			Photos[ii].Id = phot.Id
			Photos[ii].Title = phot.Title
			Photos[ii].Url = phot.Url
			Photos[ii].Thumb = phot.Thumb
		}
	}
}

func populate() {
	response, err := http.Get("https://jsonplaceholder.typicode.com/photos")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		err2 := json.Unmarshal(data, &Photos)
		var photoInterface []interface{}

		if err2 != nil {
			fmt.Println("error:", err2)
			return
		}

		err3 := json.Unmarshal(data, &Photos)
		//err4 := json.Unmarshal(data, &photoInterface)

		for _, t := range Photos {
			photoInterface = append(photoInterface, t)
		}

		if err3 != nil { //|| err4 != nil {
			fmt.Printf("input to array failed")
			return
		}
		fmt.Printf("Input to array success")
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
		client, err5 := mongo.Connect(ctx, clientOptions)
		if err5 != nil {
			fmt.Printf("Mongo request failed")
			return
		}
		fmt.Printf("Mongo request success")
		collection := client.Database("newsfeed").Collection("latihan")
		filter := bson.D{{"id", 1}}
		var result Photo
		err := collection.FindOne(context.TODO(), filter).Decode(&result)
		if err != nil {
			fmt.Println("error:", err)
			_, err4 := collection.InsertMany(context.TODO(), photoInterface)
			if err4 != nil {
				fmt.Println(err4)
			}
			fmt.Println("Inserted multiple documents")
			return
		} else {
			fmt.Println("data ready")
		}
	}
}

func init() {
	populate()
}

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type MetaData interface {
	MetaDataUser | MetaDataOrder | map[string]interface{}
}
type Event[T MetaData] struct {
	SessionID string `bson:"session_id"`
	Name      string `bson:"name"`
	MetaData  T      `bson:"meta_data"`
}
type BatchName uint8

const (
	ALL BatchName = iota
	USER
	ORDER
)

type User struct {
	Name  string `bson:"name"`
	Phone string `bson:"phone"`
}
type MetaDataUser struct {
	Ids  []string `bson:"ids"`
	Info []User   `bson:info`
}

type MetaDataOrder struct {
	Number int    `bson:"number"`
	Name   string `bson:"name"`
}

func (u User) Get[T any](test T) T {
	return test
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx,
		options.
			Client().
			ApplyURI("mongodb://codefin:1qazxsw2@localhost:27017/event_mornitoring_db"),
	)
	if err != nil {
		panic(err)
	}
	collection := client.Database("event_mornitoring_db").Collection("batch_monitoring")
	_, err = collection.InsertOne(
		context.Background(),
		Event[map[string]interface{}]{
			SessionID: "0001",
			Name:      "test",
			MetaData: map[string]interface{}{
				"ids": []string{
					"id001",
					"id002",
					"id003",
				},
				"info": []User{
					{
						Name:  "user-test",
						Phone: "0000000000",
					},
					{
						Name:  "user-test02",
						Phone: "0000000000",
					},
				},
			},
		},
	)
	if err != nil {
		panic(err)
	}

	_, err = collection.InsertOne(
		context.Background(),
		Event[map[string]interface{}]{
			SessionID: "0002",
			Name:      "test2",
			MetaData: map[string]interface{}{
				"number": 10,
				"name":   "TEST2",
			},
		},
	)
	if err != nil {
		panic(err)
	}
	var result Event[MetaDataOrder]
	find(*collection, result)
	// fmt.Println("Result: ", result)
}

func find[T any](collection mongo.Collection, result T) {
	cur, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		// To decode into a struct, use cursor.Decode()
		var nresp = result
		err := cur.Decode(&nresp)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Result: ", nresp)

		// switch reflect.TypeOf(destination) {
		// case reflect.TypeOf(new(EventUser)):
		// 	fmt.Println("eventUser")
		// 	destination = InitEventUser(result)
		// 	fmt.Println("Result: ", destination)

		// }
		// do something with result...

		// To get the raw bson bytes use cursor.Current
		// raw := cur.Current

		// do something with raw...
	}
	if err := cur.Err(); err != nil {
		panic(err)
	}

}




func CallChangeFundcodePortService(old, new, scheduleRef string) (bool, error) {

}
func CallChangeFundcodeFundService(old, new, scheduleRef string) (bool, error) {

}
func CallChangeFundcodeRoboService(old, new, scheduleRef string) (bool, error) {

}
func CallChangeFundcodePlanService(old, new, scheduleRef string) (bool, error) {

}

type ChangeFundnameProcessors []fn(old, new, scheduleRef) (bool, error)

const(
	CFPs ChangeFundnameProcessors = {
		CallCallChangeFundcodePortService,
		
	}
)

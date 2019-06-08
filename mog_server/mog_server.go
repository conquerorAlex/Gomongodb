package mog_server

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type User struct{
	Id primitive.ObjectID
	Name string
	Password string
}

func AddUser(db *mongo.Database){
	coll := db.Collection("user")
	/*err := coll.Drop(context.Background())
	if err != nil {
		fmt.Println("err:", err)
	}*/
	/*user := bson.D{
		{"name", "tom"},
		{"password", "passwd"},
	}*/
	result, err := coll.InsertOne(
		context.Background(),
		bson.D{
			{"name", "tom"},
			{"password", "passwd"},
		})
	if err != nil{
		fmt.Println("err:", err)
	}else{
		fmt.Println("result:", result)
	}
}

func FindUser(db *mongo.Database){
	fmt.Println("finduser")
	coll := db.Collection("user")
	/*err := coll.Drop(context.Background())
	if err != nil {
		fmt.Println("err:", err)
	}*/

	cur, err := coll.Find(context.Background(),	bson.D{{"name", "tom"}})
	if err != nil{
		fmt.Println("err:", err)
	}

	if cur == nil{
		fmt.Println("cur none.")
	}

	var results []*User
	//var r1 User
	//var tmp bson.RawValue
	var tmpId primitive.ObjectID
	for cur.Next(context.Background()) {
		fmt.Println("current:", cur.Current)
		ret := cur.Current.Lookup("_id").Unmarshal(&tmpId)
		if ret != nil{
			fmt.Println(ret)
		}
		fmt.Println("tmpId:", tmpId)

		var elem User
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	// Close the cursor once finished
	cur.Close(context.TODO())

	//
	/*ret := coll.FindOne(context.Background(), bson.D{{"_id", tmpId}}).Decode(&r1)
	if ret != nil{
		fmt.Println(ret)
	}
	fmt.Println("r1=", r1)*/
	rawBytes, err := coll.FindOne(context.Background(), bson.D{{"_id", tmpId}}).DecodeBytes()
	if err != nil{
		fmt.Println(err)
	}
	tmp2:=rawBytes.Lookup("_id")
	fmt.Println("tmp2:", tmp2)

	fmt.Println(len(results))
	fmt.Println("name:", results[0].Name)


}

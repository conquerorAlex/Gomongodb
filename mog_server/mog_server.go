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
	users := []interface{}{
		bson.D{
			{"name", "tom"},
			{"password", "passwd0"},
		},
		bson.D{
			{"name", "lili"},
			{"password", "passwd1"},
		},
		bson.D{
			{"name", "jan"},
			{"password", "passwd2"},
		},

	}
	result, err := coll.InsertMany(
		context.Background(),
		users,
		)
	if err != nil{
		fmt.Println("err:", err)
	}else{
		fmt.Println("result:", result)
	}
}

//and
func FindUser(db *mongo.Database){
	fmt.Println("find user with and:")
	coll := db.Collection("user")
	/*err := coll.Drop(context.Background())
	if err != nil {
		fmt.Println("err:", err)
	}*/

	cur, err := coll.Find(
		context.Background(),
		bson.D{
			{"name", "tom"},
			{"password", "passwd0"},
		})
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
}

//or
func FindManyUser(db *mongo.Database){
	fmt.Println("find users by or:")
	coll := db.Collection("user")

	cur, err := coll.Find(
		context.Background(),
		bson.D{
			{"$or", bson.A{
				bson.D{
					{"name", "tom"},
				},
				bson.D{
					{"name", "lili"},
				},
			}},
		})
	if err != nil{
		fmt.Println("err:", err)
	}

	if cur == nil{
		fmt.Println("cur none.")
	}

	var results []*User
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

	fmt.Println("num of found results:", len(results))
	fmt.Println("name:", results[0].Name)
}

//findone
func FindOneUser(db *mongo.Database){
	fmt.Println("find one user:")
	coll := db.Collection("user")

	//
	/*ret := coll.FindOne(context.Background(), bson.D{{"_id", tmpId}}).Decode(&r1)
	if ret != nil{
		fmt.Println(ret)
	}
	fmt.Println("r1=", r1)*/

	tmpId, err := primitive.ObjectIDFromHex("5cfbafae2d49c9055d3b5a17")
	if err != nil{
		fmt.Println(err)
	}
	rawBytes, err := coll.FindOne(
		context.Background(),
		bson.D{
		{"_id", tmpId},
	}).DecodeBytes()
	if err != nil{
		fmt.Println(err)
	}
	tmp2:=rawBytes.Lookup("_id")
	fmt.Println("tmp2:", tmp2)

}

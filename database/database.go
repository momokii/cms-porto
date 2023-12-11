package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collections struct {
	UserCollection          *mongo.Collection
	SelfInfoCollection      *mongo.Collection
	BannerInfoCollection    *mongo.Collection
	MyPageInfoCollection    *mongo.Collection
	HobbyInfoCollection     *mongo.Collection
	EducationInfoCollection *mongo.Collection
	ContactInfoCollection   *mongo.Collection
}

func InitMongoDB() {
	ctx := context.TODO()

	// * mongo db connect
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// * checking database connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	db := client.Database(os.Getenv("MONGODB_DB"))
	Collections.UserCollection = db.Collection("users")
	Collections.SelfInfoCollection = db.Collection("selfInfo")
	Collections.BannerInfoCollection = db.Collection("bannerInfo")
	Collections.MyPageInfoCollection = db.Collection("myPageInfo")
	Collections.HobbyInfoCollection = db.Collection("hobbyInfo")
	Collections.EducationInfoCollection = db.Collection("educationInfo")
	Collections.ContactInfoCollection = db.Collection("contactInfo")
}

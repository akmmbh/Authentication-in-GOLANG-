package database

import(
	"fmt"
	"log"
	"time"
	"os"
	"context"
	"github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client{

	err:= godotenv.Load()
  	if err != nil{

		log.Fatal("error loading .env file")

	}	
//read mongo db url from .env 
MongoDb := os.Getenv("MONGODB_URL")
if MongoDb == ""{
	log.Fatal("Mongodb_url not found in .env")
}

//set timeout context
ctx, cancel:= context.WithTimeout(context.Background(),10*time.Second)
defer cancel()

//conntect directly 
client ,err := mongo.Connect(ctx,options.Client().ApplyURI(MongoDb));
if err!=nil{
	log.Fatalf("Failed to connect to monogDB: %v",err)

}

//Ping to verify connection 
if err:= client.Ping(ctx,nil);err!=nil{
	log.Fatalf("Mongodb not responding : %v",err)
}
log.Println("connected to MongoDB!")
fmt.Println("Connected to MongoDB!")
return client 
}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client ,collectionName string )*mongo.Collection{
	var collection *mongo.Collection = client.Database("cluster0").Collection(collectionName)
	return collection
}
package container



import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client


 func Connect() {

	clientOptions := options.Client().ApplyURI("mongodb://app_customer_assets:db076e07e2a@DEV-SHARED-MONGODB-01A.dev-sicredi.in:27017,DEV-SHARED-MONGODB-02C.dev-sicredi.in:27017,DEV-SHARED-MONGODB-03A.dev-sicredi.in:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}

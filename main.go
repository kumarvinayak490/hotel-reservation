package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kumarvinayak490/hotel-reservation/api"
	"github.com/kumarvinayak490/hotel-reservation/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



const dburi = "mongodb://localhost:27017"


func main(){

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}
	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))
	fmt.Println(userHandler)
	app:= fiber.New()
	app.Get("/:id", userHandler.HandleGetUser)

	app.Listen(":8000")
}
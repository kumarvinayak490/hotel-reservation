package db

import (
	"context"

	"github.com/kumarvinayak490/hotel-reservation/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const dbName = "hotel-reservation"
const COLLECTION = "users"

type UserStore interface {
	GetUserByID(context.Context,  string) (*models.User, error)
}


type MongoUserStore struct{
	client 	*mongo.Client
	coll *mongo.Collection
}

func NewMongoUserStore(c *mongo.Client) *MongoUserStore{

	return &MongoUserStore{
		client: c,
		coll: c.Database(dbName).Collection(COLLECTION),
	}
}

func (m MongoUserStore) GetUserByID(ctx context.Context, id string) (*models.User, error){
	objId, err:=primitive.ObjectIDFromHex(id)
	if err!=nil{
		return nil, err
	}
	var user models.User
	if err := m.coll.FindOne(ctx, bson.M{"_id":objId}).Decode(&user); err!=nil{
		return nil, err
	}
	return &user, nil;
}
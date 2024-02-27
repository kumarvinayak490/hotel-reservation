package models

type User struct{
	ID 	string `bson:"_id,omitempty" json:"id,omitempty"`
	Firstname string `bson:"firstname" json:"firstname"`
	LastName string `bson:"lastname" json:"lastname"`
}
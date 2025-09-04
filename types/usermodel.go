package types

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string        `bson:"name" json:"name"`
	Email     string        `bson:"email" json:"email"`
	Password  string        `bson:"password" json:"-"`
	CreatedAt time.Time     `bson:"createdAt" json:"createdAt"`
}

// creates a user with time
func NewUser(name, email, password string) *User {
	tt := time.Now().UTC()
	//todo: salt the password
	return &User{
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: tt,
	}
}

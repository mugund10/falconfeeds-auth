package types

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string        `bson:"name" json:"name"`
	Email     string        `bson:"email" json:"email"`
	Password  string        `bson:"password" json:"-"`
	CreatedAt time.Time     `bson:"createdAt" json:"createdAt"`
}

// creates a user with time
func NewUser(name, email, password string) (*User, error) {
	epass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	tt := time.Now().UTC()
	return &User{
		Name:      name,
		Email:     email,
		Password:  string(epass),
		CreatedAt: tt,
	}, nil
}

// validates password with hash
func (u *User) ValidatePass(pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pass))
	return err == nil
}

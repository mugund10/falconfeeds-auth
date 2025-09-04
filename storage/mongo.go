package storage

import (
	"context"

	"github.com/mugund10/falconfeeds-auth/types"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MongoUserStore struct {
	db   *mongo.Database
	coll string
}

func NewMongoUserStore(db *mongo.Database) *MongoUserStore {
	return &MongoUserStore{db: db, coll: "users"}
}

// inserts users to mongodb
func (m *MongoUserStore) Insert(ctx context.Context, user *types.User) error {
	res, err := m.db.Collection(m.coll).InsertOne(ctx, user)
	if err != nil {
		return err
	}
	if id, ok := res.InsertedID.(bson.ObjectID); ok {
		user.ID = id
	}
	return err
}

// finds user by email
func (m *MongoUserStore) GetByEmail(ctx context.Context, email string) (*types.User, error) {
	var emailUser types.User
	err := m.db.Collection(m.coll).FindOne(ctx, bson.M{"email": email}).Decode(&emailUser)
	if err != nil {
		return nil, err
	}
	return &emailUser, err
}

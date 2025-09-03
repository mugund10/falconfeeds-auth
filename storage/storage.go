package storage

import (
	"context"

	"github.com/mugund10/falconfeeds-auth/types"
)

// standard db interface for user persistence
type UserStorer interface {
	Insert(context.Context, types.User)
	GetByEmail(context.Context, string) (types.User, error)
}

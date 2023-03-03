package repo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserStorageI interface {
	Create(ctx context.Context, u *User) (*User, error)
	Get(ctx context.Context, id string) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetAll(ctx context.Context, params *GetAllUsersParams) (*GetAllUsersResult, error)
	Update(ctx context.Context, u *User) (*User, error)
	Delete(ctx context.Context, userId string) error
}

type User struct {
	Id        primitive.ObjectID `bson:"_id"`
	FirstName string             `bson:"first_name"`
	LastName  string             `bson:"last_name"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	CreatedAt time.Time          `bson:"created_at"`
}

type GetAllUsersResult struct {
	Users []*User
	Count int32
}

type GetAllUsersParams struct {
	Limit  int32
	Page   int32
	Search string
}

package repo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserStorageI interface {
	Create(u *User) (*User, error)
	Get(id string) (*User, error)
	GetByEmail(email string) (*User, error)
	GetAll(params *GetAllUsersParams) (*GetAllUsersResult, error)
	Update(u *User) (*User, error)
	Delete(userId int64) error
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

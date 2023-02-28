package mongodb

import (
	"github.com/SaidovZohid/make-url-short/storage/repo"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepo struct {
	db *mongo.Database
}

func NewUser(mongoDB *mongo.Database) repo.UserStorageI {
	return &userRepo{
		db: mongoDB,
	}
}

func (u *userRepo) Create(user *repo.User) (*repo.User, error) {
	return nil, nil
}

func (u *userRepo) Get(id int64) (*repo.User, error) {
	return nil, nil
}

func (u *userRepo) GetByEmail(email string) (*repo.User, error) {
	return nil, nil
}

func (u *userRepo) GetAll(params *repo.GetAllUsersParams) (*repo.GetAllUsersResult, error) {
	return nil, nil
}

func (u *userRepo) Update(user *repo.User) (*repo.User, error) {
	return nil, nil
}

func (u *userRepo) Delete(userID int64) error {
	return nil
}

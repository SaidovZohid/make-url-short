package mongodb

import (
	"context"
	"time"

	"github.com/SaidovZohid/make-url-short/storage/repo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepo struct {
	db *mongo.Collection
}

func NewUser(mongoDB *mongo.Collection) repo.UserStorageI {
	return &userRepo{
		db: mongoDB,
	}
}

func (u *userRepo) Create(user *repo.User) (*repo.User, error) {
	user.Id = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	res, err := u.db.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}
	user.Id = res.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (u *userRepo) Get(id string) (*repo.User, error) {
	var res repo.User
	filter := bson.D{{"_id", id}}
	err := u.db.FindOne(context.Background(), filter).Decode(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
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

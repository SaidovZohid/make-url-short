package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/SaidovZohid/make-url-short/storage/repo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepo struct {
	db *mongo.Collection
}

func NewUser(mongoDB *mongo.Collection) repo.UserStorageI {
	return &userRepo{
		db: mongoDB,
	}
}

func (u *userRepo) Create(ctx context.Context, user *repo.User) (*repo.User, error) {
	user.Id = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	res, err := u.db.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}
	user.Id = res.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (u *userRepo) Get(ctx context.Context, id string) (*repo.User, error) {
	var res repo.User
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = u.db.FindOne(context.Background(), bson.M{"_id": objectId}).Decode(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (u *userRepo) GetByEmail(ctx context.Context, email string) (*repo.User, error) {
	var res repo.User
	err := u.db.FindOne(context.Background(), bson.M{"email": email}).Decode(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (u *userRepo) Update(ctx context.Context, user *repo.User) (*repo.User, error) {
	filter := bson.M{"_id": user.Id}
	update := bson.M{"$set": bson.M{"first_name": user.FirstName, "last_name": user.LastName}}
	res, err := u.db.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}
	if res.ModifiedCount == 0 {
		return nil, mongo.ErrNoDocuments
	}

	return user, nil
}

func (u *userRepo) Delete(ctx context.Context, userID string) error {
	objectId, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	res, err := u.db.DeleteOne(context.Background(), bson.M{"_id": objectId})
	if err != nil {
		return err
	}
	if res.DeletedCount != 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}

func (u *userRepo) GetAll(ctx context.Context, params *repo.GetAllUsersParams) (*repo.GetAllUsersResult, error) {
	offset := (params.Page - 1) * params.Limit
	findOptions := options.Find()
	findOptions.SetLimit(int64(params.Limit))
	findOptions.SetSkip(int64(offset))
	regexPattern := fmt.Sprintf("^%s", params.Search)
	regex := bson.M{"$regex": primitive.Regex{Pattern: regexPattern, Options: "i"}}

	filter := bson.M{
		"$or": []bson.M{
			{"first_name": regex},
			{"last_name": regex},
			{"email": regex},
		},
	}

	cursor, err := u.db.Find(context.Background(), filter, findOptions)
	if err != nil {
		return nil, err
	}

	results := repo.GetAllUsersResult{
		Users: make([]*repo.User, 0),
	}
	for cursor.Next(context.Background()) {
		var result repo.User
		err := cursor.Decode(&result)
		if err != nil {
			return nil, err
		}

		results.Users = append(results.Users, &result)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	err = cursor.Close(context.Background())
	if err != nil {
		return nil, err
	}

	return &results, nil
}

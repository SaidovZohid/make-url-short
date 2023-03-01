package storage

import (
	"github.com/SaidovZohid/make-url-short/storage/mongodb"
	"github.com/SaidovZohid/make-url-short/storage/repo"
	"go.mongodb.org/mongo-driver/mongo"
)

type StorageI interface {
	User() repo.UserStorageI
	Url() repo.UrlStorageI
}

type Storage struct {
	userRepo repo.UserStorageI
	urlRepo  repo.UrlStorageI
}

type Collections struct {
	UserCollection *mongo.Collection
	UrlCollection  *mongo.Collection
}

func NewStorage(c *Collections) StorageI {
	return &Storage{
		userRepo: mongodb.NewUser(c.UserCollection),
		urlRepo:  mongodb.NewUrl(c.UrlCollection),
	}
}

func (s *Storage) User() repo.UserStorageI {
	return s.userRepo
}

func (s *Storage) Url() repo.UrlStorageI {
	return s.urlRepo
}

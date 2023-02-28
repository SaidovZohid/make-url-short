package storage

import (
	"github.com/SaidovZohid/make-url-short/config"
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

func NewStorage(cfg *config.Config, db *mongo.Client) StorageI {
	urlCol := db.Database(cfg.MongoDB.UrlCollection)
	userCol := db.Database(cfg.MongoDB.UserCollection)
	return &Storage{
		userRepo: mongodb.NewUser(userCol),
		urlRepo:  mongodb.NewUrl(urlCol),
	}
}

func (s *Storage) User() repo.UserStorageI {
	return s.userRepo
}

func (s *Storage) Url() repo.UrlStorageI {
	return s.urlRepo
}

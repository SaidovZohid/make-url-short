package mongodb

import (
	"github.com/SaidovZohid/make-url-short/storage/repo"
	"go.mongodb.org/mongo-driver/mongo"
)

type urlRepo struct {
	db *mongo.Database
}

func NewUrl(db *mongo.Database) repo.UrlStorageI {
	return &urlRepo{
		db: db,
	}
}

func (u *urlRepo) Create(url *repo.Url) (*repo.Url, error) {
	return url, nil
}

func (u *urlRepo) Get(url string) (*repo.Url, error) {
	return nil, nil
}

func (u *urlRepo) GetAll(params *repo.GetAllUrlsParams) (*repo.GetAllUrlsResult, error) {
	return nil, nil
}

func (u *urlRepo) Update(url *repo.Url) (*repo.Url, error) {
	return nil, nil
}

func (u *urlRepo) Delete(id, userID int64) error {
	return nil
}

func (u *urlRepo) DecrementClick(url string) error {
	return nil
}

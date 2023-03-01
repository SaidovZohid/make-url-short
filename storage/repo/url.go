package repo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UrlStorageI interface {
	Create(u *Url) (*Url, error)
	Get(url string) (*Url, error)
	GetAll(params *GetAllUrlsParams) (*GetAllUrlsResult, error)
	DecrementClick(url string) error
	Update(u *Url) (*Url, error)
	Delete(id, userID int64) error
}

type Url struct {
	Id          primitive.ObjectID `bson:"_id"`
	UserId      int64              `bson:"user_id"`
	OriginalUrl string             `bson:"original_url"`
	HashedUrl   string             `bson:"hashed_url"`
	MaxClicks   *int64             `bson:"max_clicks"`
	ExpiresAt   *time.Time         `bson:"expires_at"`
	CreatedAt   time.Time          `bson:"created_at"`
}

type GetAllUrlsResult struct {
	Urls  []*Url
	Count int32
}

type GetAllUrlsParams struct {
	Limit  int32
	Page   int32
	UserID int64
	Search string
}

package main

import (
	"context"
	"fmt"

	"github.com/SaidovZohid/make-url-short/config"
	"github.com/SaidovZohid/make-url-short/pkg/logger"
	"github.com/SaidovZohid/make-url-short/storage"
	"github.com/SaidovZohid/make-url-short/storage/repo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	logger.Init()
	log := logger.GetLogger()

	log.Info("Getting Configuration Files...")
	cfg, err := config.Load(".")
	if err != nil {
		log.WithError(err).Fatal("Failed to Get Environtment Variables From --> .env <-- File")
	}

	serverApi := options.ServerAPI(options.ServerAPIVersion1)
	url := fmt.Sprintf("mongodb://%s:%s@%s:%s", cfg.MongoDB.User, cfg.MongoDB.Password, cfg.MongoDB.Host, cfg.MongoDB.Port)
	opts := options.Client().ApplyURI(url).SetServerAPIOptions(serverApi)

	log.Info("Creating a client and connect to the server")
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		log.WithError(err).Fatalf("failed to create client to mongodb -> %s", url)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	log.Info("Sending a ping to confirm a successful connection")
	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		panic(err)
	}
	log.Info("Pinged the primary node of the cluster. You successfully connected to MongoDB!")
	db := client.Database(cfg.MongoDB.Database)
	colUser := db.Collection(cfg.MongoDB.UserCollection)
	colUrl := db.Collection(cfg.MongoDB.UrlCollection)
	strg := storage.NewStorage(&storage.Collections{
		UserCollection: colUser,
		UrlCollection:  colUrl,
	})
	res, err := strg.User().GetAll(context.Background(), &repo.GetAllUsersParams{
		Limit: 10,
		Page:  1,
		Search: "ai",
	})
	if err != nil {
		log.Println(err)
	}
	for _, v := range res.Users {
		log.Println(v)
	}
	// objectId, err := primitive.ObjectIDFromHex("63ff0ebd20379a3c6912f1cc")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// _, err = strg.User().Update(context.Background(), &repo.User{
	// 	Id:        objectId,
	// 	FirstName: "Zohid",
	// 	LastName:  "Saidov",
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// res, err := strg.User().Get(context.Background(), "63ff0ebd20379a3c6912f1cc")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(res)

}

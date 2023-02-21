package infra

import (
	"context"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

var (
	once  sync.Once
	store *mongo.Database
)

func NewMongoConnection(databaseURI string) *mongo.Database {
	once.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		cs, err := connstring.Parse(databaseURI)
		if err != nil {
			log.Fatal(err)
		}

		clientOpts := options.Client().ApplyURI(cs.String())

		client, err := mongo.Connect(ctx, clientOpts)
		if err != nil {
			log.Fatal(err)
		}

		store = client.Database(cs.Database)
	})

	return store
}

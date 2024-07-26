package config

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/pkg/errors"

	// Add the MongoDB driver packages
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewDatabases returns a new instance of MongoDB client
func NewDatabases(c *Config) (*mongo.Client, error) {

	// Try default URI
	client, err := connectMongoDB(c.Database, "")
	if err != nil {
		// Try with authSource=admin if default fails
		client, err = connectMongoDB(c.Database, "?authSource=admin")
		if err != nil {
			return nil, errors.Wrap(err, "failed to connect to MongoDB")
		}
	}

	return client, nil
}

func connectMongoDB(m *DatabaseCfg, authOptions string) (*mongo.Client, error) {

	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s/%s%s", m.DBUser, m.DBPass, m.Host, m.DBName, authOptions)
	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to MongoDB")
	}

	// Ping the database to verify connection
	if err := client.Ping(context.Background(), nil); err != nil {
		return nil, errors.Wrap(err, "failed to ping MongoDB")
	}

	log.Infof("Successfully connected MongoDB with URI: %s", m.Host)

	return client, nil
}

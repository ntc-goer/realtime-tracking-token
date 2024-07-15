package database

import (
	"fmt"
	"github.com/ntc-goer/parser-exercise/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"time"
)

const (
	connectionStringTemplate = "mongodb://%s:%s@%s/%s"
)

type MongoDB struct {
	ConnectTimeout int      `json:"connect_timeout"`
	Username       string   `json:"username"`
	Password       string   `json:"password"`
	Hosts          []string `json:"hosts"`
	Options        string   `json:"options"`
	Database       string   `json:"database"`
}

func NewMongoDB(c *config.Config) *MongoDB {
	return &MongoDB{
		Username: c.MongoDB.UserName,
		Password: c.MongoDB.Password,
		Hosts:    []string{c.MongoDB.Host},
		Database: c.MongoDB.Database,
	}
}

func InitMongoDB(ctx context.Context, mongo *MongoDB) *mongo.Database {
	client, err := mongo.Connect()
	if err != nil {
		message := fmt.Sprintf("Cannot connect to mongoDB: %v", err)
		panic(message)
	}
	if err = client.Ping(ctx, nil); err != nil {
		message := fmt.Sprintf("Cannot connect to mongoDB: %v", err)
		panic(message)
	}
	return client.Database(mongo.Database)
}

func (m *MongoDB) Connect() (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	connectionStr := connectionStringTemplate
	uri := fmt.Sprintf(connectionStr, m.Username, m.Password, m.Hosts[0], m.Database)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, opts)
	client.Database(m.Database)
	return client, err
}

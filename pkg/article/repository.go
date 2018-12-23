package article

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo/readpref"

	"github.com/mongodb/mongo-go-driver/mongo"
)

type IRepository interface {
	List()
}

type IArticleRepository interface {
	IRepository
}

type MongoDB struct {
	db     *mongo.Database
	client *mongo.Client
}

func NewMongoDB() *MongoDB {
	client, err := mongo.NewClient("mongodb://localhost:27018")
	checkErrLogFatal(err)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	checkErrLogFatal(err)
	err = client.Connect(ctx)
	checkErrLogFatal(err)
	db := client.Database("mydb")
	m := &MongoDB{
		db:     db,
		client: client,
	}
	m.seed()
	fmt.Println("Seed Complete!")
	return m
}

func (m *MongoDB) seed() {
	// Load values from JSON file to model
	byteValues, err := ioutil.ReadFile("data/articles/seed_data.json")
	if err != nil {
		log.Fatal(err)
	}

	var articles []Article
	json.Unmarshal(byteValues, &articles)

	// Insert articles into DB
	var ppl []interface{}
	for _, p := range articles {
		ppl = append(ppl, p)
	}

	_, err = m.db.Collection("articles").InsertMany(context.Background(), ppl)
	checkErrLogFatal(err)
}

func (m *MongoDB) Ping() error {
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	err := m.client.Ping(ctx, readpref.Primary())
	return err
}

func (m *MongoDB) List() {
	fmt.Println("listed from the mongo Repo")
}

func checkErrLogFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

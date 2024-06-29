package models

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var client *mongo.Client

func init() {
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

type ResumeData struct {
	Skills     []string `json:"skills"`
	Education  []string `json:"education"`
	Experience []string `json:"experience"`
}

type User struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	UserType        string `json:"usertype"`
	Address         string `json:"address"`
	ProfileHeadline string `json:"profile_headline"`
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Profile struct {
	Applicant  string   `json:"applicant"`
	Resume     string   `json:"resume"`
	Skills     []string `json:"skills"`
	Education  []string `json:"education"`
	Experience []string `json:"experience"`
}

func (u *User) CreateUser() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := client.Database("recruitment").Collection("users")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	_, err = collection.InsertOne(ctx, u)
	return err
}

func (u *User) GetUserByEmail(email string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := client.Database("recruitment").Collection("users")
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&u)
	return err
}

func (p *Profile) SaveProfile() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := client.Database("recruitment").Collection("profiles")
	_, err := collection.InsertOne(ctx, p)
	return err
}

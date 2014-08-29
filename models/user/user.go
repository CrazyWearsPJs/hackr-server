package user

import (
	"code.google.com/p/go-uuid/uuid"
	"code.google.com/p/go.crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
	_ "os"
	"time"
)

type User struct {
	Id             bson.ObjectId `bson:"_id"`
	Email          string        `bson:"email"`
	HashedPassword string        `bson:"pass"`
	APIKey         string        `bson:"key"`
	Created        time.Time     `bson:"created"`
	Code           []string      `bson:"code"`
}

func New(email, password []byte) (*User, error) {
	s_email := string(email)
	id := bson.NewObjectId()
	hashed_password, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	s_hashed_password := string(hashed_password)
	api_key := uuid.New()
	created := time.Now()
	code := []string{}

	return &User{id, s_email, s_hashed_password, api_key, created, code}, nil
}

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
	Updated        time.Time     `bson:"updated"`
	Code           []string      `bson:"code"`
}

func New(email, password string) (*User, error) {
	id := bson.NewObjectId()
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	s_hashed_password := string(hashed_password)
	api_key := uuid.New()
	created := time.Now()
	updated := created
	code := []string{}

	return &User{id, email, s_hashed_password, api_key, created, updated, code}, nil
}

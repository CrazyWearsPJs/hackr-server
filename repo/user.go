package repo

import (
	"github.com/CrazyWearsPJs/hackr/models/user"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type UserRepo struct {
	Collection *mgo.Collection
}

func (r UserRepo) Add(u *user.User) error {
	if u.Id.Hex() == "" {
		u.Id = bson.NewObjectId()
	}

	if u.Created.IsZero() {
		u.Created = time.Now()
	}

	_, err := r.Collection.UpsertId(u.Id, u)
	return err
}

func (r UserRepo) FindUserByEmail(email string) (*user.User, error) {
	var u *user.User
	err := r.Collection.Find(bson.M{"email": email}).Limit(1).All(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

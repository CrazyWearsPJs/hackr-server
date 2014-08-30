package repo

import (
	"time"

	"github.com/CrazyWearsPJs/hackr/models/user"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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
	var u user.User
	err := r.Collection.Find(bson.M{"email": email}).One(&u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r UserRepo) SubmitCode(email string, code string) (*user.User, error) {
	var u user.User

	change := mgo.Change{
		ReturnNew: true,
		Update: bson.M{
			"$set": bson.M{
				"updated": time.Now(),
			},
			"$push": bson.M{
				"code": code,
			},
		},
	}

	_, err := r.Collection.Find(bson.M{"email": email}).Apply(change, &u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

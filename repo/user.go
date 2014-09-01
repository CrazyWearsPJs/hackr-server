package repo

import (
	"errors"
	"fmt"
	"github.com/CrazyWearsPJs/hackr/models/user"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type UserRepo struct {
	Collection *mgo.Collection
}

func (r UserRepo) Add(u *user.User) error {

	u_repo, _ := r.FindUserByEmail(u.Email)

	if u_repo != nil {
		msg := fmt.Sprintf("User with the email %v already exists", u.Email)
		return errors.New(msg)
	}

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
	fmt.Printf("wat")
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

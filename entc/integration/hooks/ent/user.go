// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/entc/integration/hooks/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Version holds the value of the "version" field.
	Version int `json:"version,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Worth holds the value of the "worth" field.
	Worth uint `json:"worth,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges          UserEdges `json:"edges"`
	best_friend_id *int
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Cards holds the value of the cards edge.
	Cards []*Card
	// Friends holds the value of the friends edge.
	Friends []*User
	// BestFriend holds the value of the best_friend edge.
	BestFriend *User
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// CardsOrErr returns the Cards value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CardsOrErr() ([]*Card, error) {
	if e.loadedTypes[0] {
		return e.Cards, nil
	}
	return nil, &NotLoadedError{edge: "cards"}
}

// FriendsOrErr returns the Friends value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FriendsOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.Friends, nil
	}
	return nil, &NotLoadedError{edge: "friends"}
}

// BestFriendOrErr returns the BestFriend value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) BestFriendOrErr() (*User, error) {
	if e.loadedTypes[2] {
		if e.BestFriend == nil {
			// The edge best_friend was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.BestFriend, nil
	}
	return nil, &NotLoadedError{edge: "best_friend"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullInt64{},  // version
		&sql.NullString{}, // name
		&sql.NullInt64{},  // worth
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*User) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // best_friend_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(values ...interface{}) error {
	if m, n := len(values), len(user.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	u.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field version", values[0])
	} else if value.Valid {
		u.Version = int(value.Int64)
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[1])
	} else if value.Valid {
		u.Name = value.String
	}
	if value, ok := values[2].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field worth", values[2])
	} else if value.Valid {
		u.Worth = uint(value.Int64)
	}
	values = values[3:]
	if len(values) == len(user.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field best_friend_id", value)
		} else if value.Valid {
			u.best_friend_id = new(int)
			*u.best_friend_id = int(value.Int64)
		}
	}
	return nil
}

// QueryCards queries the cards edge of the User.
func (u *User) QueryCards() *CardQuery {
	return (&UserClient{config: u.config}).QueryCards(u)
}

// QueryFriends queries the friends edge of the User.
func (u *User) QueryFriends() *UserQuery {
	return (&UserClient{config: u.config}).QueryFriends(u)
}

// QueryBestFriend queries the best_friend edge of the User.
func (u *User) QueryBestFriend() *UserQuery {
	return (&UserClient{config: u.config}).QueryBestFriend(u)
}

// Update returns a builder for updating this User.
// Note that, you need to call User.Unwrap() before calling this method, if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", version=")
	builder.WriteString(fmt.Sprintf("%v", u.Version))
	builder.WriteString(", name=")
	builder.WriteString(u.Name)
	builder.WriteString(", worth=")
	builder.WriteString(fmt.Sprintf("%v", u.Worth))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}

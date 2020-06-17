// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// CarsColumns holds the columns for the "cars" table.
	CarsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_car", Type: field.TypeInt, Nullable: true},
	}
	// CarsTable holds the schema information for the "cars" table.
	CarsTable = &schema.Table{
		Name:       "cars",
		Columns:    CarsColumns,
		PrimaryKey: []*schema.Column{CarsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "cars_users_car",
				Columns: []*schema.Column{CarsColumns[1]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:        "groups",
		Columns:     GroupsColumns,
		PrimaryKey:  []*schema.Column{GroupsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// PetsColumns holds the columns for the "pets" table.
	PetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "owner_id", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// PetsTable holds the schema information for the "pets" table.
	PetsTable = &schema.Table{
		Name:       "pets",
		Columns:    PetsColumns,
		PrimaryKey: []*schema.Column{PetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "pets_users_pets",
				Columns: []*schema.Column{PetsColumns[1]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeInt, Increment: true},
		{Name: "age", Type: field.TypeInt},
		{Name: "name", Type: field.TypeString, Size: 2147483647},
		{Name: "nickname", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString, Default: "unknown"},
		{Name: "buffer", Type: field.TypeBytes, Nullable: true},
		{Name: "title", Type: field.TypeString, Default: "SWE"},
		{Name: "renamed", Type: field.TypeString, Nullable: true},
		{Name: "blob", Type: field.TypeBytes, Nullable: true, Size: 1000},
		{Name: "state", Type: field.TypeEnum, Nullable: true, Enums: []string{"logged_in", "logged_out", "online"}},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "user_phone_age",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[4], UsersColumns[1]},
			},
		},
	}
	// FriendsColumns holds the columns for the "friends" table.
	FriendsColumns = []*schema.Column{
		{Name: "user", Type: field.TypeInt},
		{Name: "friend", Type: field.TypeInt},
	}
	// FriendsTable holds the schema information for the "friends" table.
	FriendsTable = &schema.Table{
		Name:       "friends",
		Columns:    FriendsColumns,
		PrimaryKey: []*schema.Column{FriendsColumns[0], FriendsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "friends_user",
				Columns: []*schema.Column{FriendsColumns[0]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:  "friends_friend",
				Columns: []*schema.Column{FriendsColumns[1]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CarsTable,
		GroupsTable,
		PetsTable,
		UsersTable,
		FriendsTable,
	}
)

func init() {
	CarsTable.ForeignKeys[0].RefTable = UsersTable
	PetsTable.ForeignKeys[0].RefTable = UsersTable
	FriendsTable.ForeignKeys[0].RefTable = UsersTable
	FriendsTable.ForeignKeys[1].RefTable = UsersTable
}

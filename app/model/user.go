package model

import (
	"database/sql"

	"errors"

	"github.com/graphql-go/graphql"
)

type User struct {
	ID          int         `json:"id"`
	Email       string      `json:"email"`
	Name        string      `json:"name"`
	Pwd         string      `json:"-"`
	Avatar      string      `json:"avatar"`
	Bio         string      `json:"bio"`
	Token       string      `json:"token"`
	Collections Collections `json:"collection"`
}

var AuthReturnGraph = graphql.NewObject(graphql.ObjectConfig{
	Name: "auth",
	Fields: graphql.Fields{
		"token": &graphql.Field{Type: graphql.String},
	},
})

var UserGraph = graphql.NewObject(graphql.ObjectConfig{
	Name: "user",
	Fields: graphql.Fields{
		"id":    &graphql.Field{Type: graphql.Int},
		"email": &graphql.Field{Type: graphql.String},
		"name":  &graphql.Field{Type: graphql.String},
		"pwd": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "", nil
			},
			// ignore this field
		},
		"avatar": &graphql.Field{Type: graphql.String},
		"bio":    &graphql.Field{Type: graphql.String},
		"token":  &graphql.Field{Type: graphql.String},
		"collections": &graphql.Field{
			// TODO:
			Type: graphql.String,
		},
	},
})

func NewUser(id int, email, name, pwd, avatar, bio, token string) *User {
	return &User{
		ID:     id,
		Email:  email,
		Name:   name,
		Pwd:    pwd,
		Avatar: avatar,
		Bio:    bio,
		Token:  token,
	}
}

// Save 用户注册的时候没有 uid，所以无法生成 token
func (u *User) Save(db *sql.DB) error {
	id := 0
	err := db.QueryRow("INSERT INTO users(email, name, pwd, avatar, bio) VALUES($1, $2, $3, $4, $5) RETURNING id", u.Email, u.Name, u.Pwd, u.Avatar, u.Bio).Scan(&id)
	if err != nil {
		return err
	}
	u.ID = id
	return nil
}

// Update just allow to update avatar, bio and token by userID
func (u *User) Update(db *sql.DB) error {
	_, err := db.Exec(`
	UPDATE users SET avatar=$1, bio=$2 WHERE id=$4
	`, u.Avatar, u.Bio, u.ID)
	return err
}

func FetchUserBy(db *sql.DB, id int) (*User, error) {
	rows, err := db.Query("SELECT DISTINCT ON (collections.cell) * FROM users LEFT JOIN collections ON users.id=collections.owner WHERE users.id=$1", id)
	if err != nil {
		return nil, err
	}
	users, err := getUsersInfoFrom(rows)
	if err != nil {
		return nil, err
	}
	return users[0], nil
}

func UserAuth(db *sql.DB, email, pwd string) (*User, error) {
	rows, err := db.Query("SELECT DISTINCT ON (collections.cell) * FROM users LEFT JOIN collection ON users.id=collections.owner WHERE users.email=$1 AND users.pwd=$2", email, pwd)
	if err != nil {
		return nil, err
	}
	users, err := getUsersInfoFrom(rows)
	if err != nil {
		return nil, err
	}
	return users[0], nil
}

func getUsersInfoFrom(rows *sql.Rows) ([]*User, error) {
	defer rows.Close()
	var users []*User
	for rows.Next() {
		var id int
		var email, name, pwd, avatar, bio, token string
		var collectionID, collectionCell, collectionOwner int
		if err := rows.Scan(&id, &email, &name, &pwd, &avatar, &bio, &token, &collectionID, &collectionCell, &collectionOwner); err != nil {
			return nil, err
		}
		collection := NewCollection(collectionCell, collectionOwner, collectionID)
		collections := Collections{collection}
		user := &User{
			ID:          id,
			Email:       email,
			Name:        name,
			Pwd:         pwd,
			Avatar:      avatar,
			Bio:         bio,
			Token:       token,
			Collections: collections,
		}
		users = append(users, user)
	}
	if len(users) == 0 {
		err := errors.New("no result")
		return nil, err
	}
	distinctedUsers := distinctUsers(users)
	return distinctedUsers, nil
}

// 从用户数据中去除重复的部分，留下不同的 collection， 整合之后重新构成一个用户列表返回出去
func distinctUsers(users []*User) []*User {
	if len(users) == 0 {
		return users
	}

	// 用户的 id -> Collections
	// var userCollecgtionMap map[int]Collections
	userCollecgtionMap := make(map[int]Collections)
	//  User 的 ID
	var distinctedUsers []*User

	for _, user := range users {
		// 开始去重 user 里面的数据
		var notIn = true
		for _, u := range distinctedUsers {
			if u.ID == user.ID {
				notIn = false
				break
			}
		}
		if notIn {
			distinctedUsers = append(distinctedUsers, user)
			notIn = true
		}

		// 开始整理 collection 的数据
		collections := user.Collections
		if len(collections) == 0 {
			continue
		}
		if len(userCollecgtionMap[user.ID]) == 0 {
			userCollecgtionMap[user.ID] = collections
		} else {
			// 因为是每一条数据，所以可以判定确实是有 collection 的，而且只有一条
			collectionsInMap := userCollecgtionMap[user.ID]
			userCollecgtionMap[user.ID] = append(collectionsInMap, collections[0])
		}
	}

	for _, user := range distinctedUsers {
		distinctedCollections := distinctCollections(userCollecgtionMap[user.ID])
		user.Collections = distinctedCollections
	}
	return distinctedUsers
}

func distinctCollections(collections Collections) (result Collections) {
	for _, collection := range collections {
		notIn := true
		for _, item := range result {
			if collection.Cell == item.Cell {
				notIn = false
			}
		}
		if notIn {
			result = append(result, collection)
			notIn = true
		}
	}
	return
}

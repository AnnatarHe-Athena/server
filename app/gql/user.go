package gql

import (
	"github.com/douban-girls/server/app/initial"
	"github.com/douban-girls/server/app/model"
	"github.com/douban-girls/server/app/utils"
	"github.com/graphql-go/graphql"

	"golang.org/x/crypto/scrypt"
)

var AuthArg = graphql.FieldConfigArgument{
	"email":    &graphql.ArgumentConfig{Type: graphql.String},
	"password": &graphql.ArgumentConfig{Type: graphql.String},
}

func CreateUserResolver(params graphql.ResolveParams) (interface{}, error) {
	email := params.Args["email"].(string)
	username := params.Args["username"].(string)
	password := params.Args["password"].(string)
	avatar := params.Args["avatar"].(string)
	bio := params.Args["bio"].(string)

	scrypt.Keys

	user := model.NewUser(0, email, username, password, avatar, bio, "")
	err := user.Save(initial.DB)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func AuthResolver(params graphql.ResolveParams) (interface{}, error) {
	user, err := model.UserAuth(initial.DB, params.Args["email"].(string), params.Args["password"].(string))
	if err != nil {
		return nil, err
	}

	token, err := utils.GenToken(user.ID)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"token": token,
		"id":    user.ID,
	}, nil
}

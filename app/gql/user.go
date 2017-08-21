package gql

import (
	"github.com/douban-girls/douban-girls-server/app/initial"
	"github.com/douban-girls/douban-girls-server/app/model"
	"github.com/graphql-go/graphql"
)

func CreateUserResolver(params graphql.ResolveParams) (interface{}, error) {
	email := params.Args["email"].(string)
	username := params.Args["username"].(string)
	password := params.Args["password"].(string)
	avatar := params.Args["avatar"].(string)
	bio := params.Args["bio"].(string)

	user := model.NewUser(0, email, username, password, avatar, bio, "")
	err := user.Save(initial.DB)
	if err != nil {
		return nil, err
	}
	return user, nil
}

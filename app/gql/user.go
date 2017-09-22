package gql

import (
	"errors"
	"strconv"
	"strings"

	"github.com/revel/revel"

	"github.com/douban-girls/server/app/initial"
	"github.com/douban-girls/server/app/model"
	"github.com/douban-girls/server/app/utils"
	"github.com/graphql-go/graphql"
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

	if !strings.Contains(email, "@") {
		err := errors.New("email valiation fail")
		return nil, err
	}
	realPasword := utils.GenPassword(password)

	user := model.NewUser(0, email, username, realPasword, avatar, bio, "")
	err := user.Save(initial.DB)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func AuthResolver(params graphql.ResolveParams) (interface{}, error) {
	email := params.Args["email"].(string)
	pwd := utils.GenPassword(params.Args["password"].(string))
	revel.INFO.Println(params.Args["password"], pwd)
	user, err := model.UserAuth(initial.DB, email, pwd)
	if err != nil {
		revel.INFO.Println(err)
		return nil, err
	}

	revel.INFO.Println(user.ID, user.Email)

	token, err := utils.GenToken(user.ID)
	if err != nil {
		revel.INFO.Println(err)
		return nil, err
	}

	controller := utils.GetController(params)
	controller.Session["userID"] = strconv.Itoa(user.ID)
	return map[string]interface{}{
		"token": token,
		"id":    user.ID,
	}, nil
}

func QueryUserResolver(params graphql.ResolveParams) (interface{}, error) {
	id, ok := params.Args["id"].(int)
	if !ok {
		return model.User{}, nil
	}
	user, err := model.FetchUserBy(initial.DB, id)
	if err != nil {
		return model.User{}, nil
	}
	return *user, nil
}

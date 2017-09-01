package gql

import (
	"errors"

	"github.com/douban-girls/server/app/initial"
	"github.com/douban-girls/server/app/model"
	"github.com/douban-girls/server/app/utils"
	"github.com/graphql-go/graphql"
	"github.com/revel/revel"
)

func QueryCollectionResolver(params graphql.ResolveParams) (interface{}, error) {
	isPair, err := utils.IsTokenPair(utils.GetController(params))
	if !isPair || err != nil {
		return nil, errors.New("token not pair")
	}

	userID := params.Args["userID"].(int)
	collections, err := model.FetchUserCollectionBy(initial.DB, userID)

	if err != nil {
		revel.INFO.Println(err)
		return nil, err
	}

	return collections, nil
}

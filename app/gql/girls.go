package gql

import (
	"github.com/douban-girls/douban-girls-server/app/initial"
	"github.com/douban-girls/douban-girls-server/app/model"
	"github.com/graphql-go/graphql"
)

// GirlsResolver is graphql resolver
func GirlsResolver(params graphql.ResolveParams) (interface{}, error) {
	// TODO: add redis cache here
	from := params.Args["from"].(int)
	take := params.Args["take"].(int)
	offset := params.Args["offset"].(int)
	girls, err := model.FetchGirls(initial.DB, from, take, offset)

	return girls, err
}

func CreateGirl(params graphql.ResolveParams) (interface{}, error) {
	return nil, nil
}
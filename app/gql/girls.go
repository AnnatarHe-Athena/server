package gql

import (
	"github.com/douban-girls/server/app/initial"
	"github.com/douban-girls/server/app/model"
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

func getGirlArgument(params graphql.ResolveParams, keys []string) (result map[string]string) {
	for _, val := range keys {
		result[val] = params.Args[val].(string)
	}
	return result
}

// CreateGirl will set a girl to database
func CreateGirl(params graphql.ResolveParams) (interface{}, error) {

	args := getGirlArgument(params, []string{"user", "img", "category", "text"})
	item := model.Cell{
		ID: args["user"].(int),
	}
	return nil, nil
}

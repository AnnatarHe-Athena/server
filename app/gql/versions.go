package gql

import (
	"github.com/douban-girls/server/app/initial"
	"github.com/douban-girls/server/app/model"
	"github.com/graphql-go/graphql"
)

func PlatformVersionResolver(params graphql.ResolveParams) (interface{}, error) {

	platform := params.Args["platform"].(string)
	getLastOne := params.Args["getLastOne"].(bool)

	return model.FetchPlatformSpecialOne(initial.DB, platform, getLastOne)
}

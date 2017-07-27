package controllers

import (
	"github.com/douban-girls/douban-girls-server/app/gql"
	"github.com/graphql-go/graphql"
	"github.com/revel/revel"
)

type GraphQLController struct {
	*revel.Controller
}

// is Get method for fetch data ^_^
func (g *GraphQLController) Fetch() revel.Result {
	query := g.Params.Get("query")
	revel.INFO.Println(query)

	params := graphql.Params{
		Schema:        gql.GraphQLSchema,
		RequestString: query,
	}
	revel.INFO.Println(params)

	result := graphql.Do(params)
	return g.RenderJSON(result)

}

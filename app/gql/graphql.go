package gql

import (
	"github.com/douban-girls/douban-girls-server/app/initial"
	"github.com/douban-girls/douban-girls-server/app/model"
	"github.com/graphql-go/graphql"
	"github.com/revel/revel"
)

var GraphQLSchema graphql.Schema

func getRootSchema() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		// 不能有空格等特殊字符
		Name: "RootSchema",
		Fields: graphql.Fields{
			"users": &graphql.Field{
				Type:        model.UserGraph,
				Description: "a user",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, ok := params.Args["id"].(int)
					if !ok {
						return model.User{}, nil
					}
					user, err := model.FetchUserBy(initial.DB, id)
					if err != nil {
						return model.User{}, nil
					}
					return *user, nil
				},
			},
		},
	})
}

func InitGraphQLSchema() {

	rootQuery := getRootSchema()
	var err error

	GraphQLSchema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
		// TODO:
		// Mutation:
	})
	if err != nil {
		revel.INFO.Println(err)
	}
	revel.INFO.Println(GraphQLSchema)
}

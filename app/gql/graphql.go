package gql

import (
	"github.com/douban-girls/server/app/initial"
	"github.com/douban-girls/server/app/model"
	"github.com/graphql-go/graphql"
	"github.com/revel/revel"
)

// GraphQLSchema is root schema
var GraphQLSchema graphql.Schema

func getRootSchema() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		// 不能有空格等特殊字符
		Name: "RootSchema",
		Fields: graphql.Fields{
			"auth": &graphql.Field{
				Type:        model.AuthReturnGraph,
				Description: "user auth by email and password",
				Args:        AuthArg,
				Resolve:     AuthResolver,
			},
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
			"categories": &graphql.Field{
				Type:        graphql.NewList(model.CategoryGraphqlSchema),
				Description: "categories",
				Resolve:     CategoriesResolver,
			},
			"girls": &graphql.Field{
				Type:        graphql.NewList(model.GirlGraphqlSchema),
				Description: "girls",
				Args: graphql.FieldConfigArgument{
					"offset": &graphql.ArgumentConfig{Type: graphql.Int},
					"take":   &graphql.ArgumentConfig{Type: graphql.Int},
					"from":   &graphql.ArgumentConfig{Type: graphql.Int},
				},
				Resolve: GirlsResolver,
			},
		},
	})
}

func getRootMutation() *graphql.Object {

	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"addUser": &graphql.Field{
				Type:        model.UserGraph,
				Description: "create a new user",
				Args: graphql.FieldConfigArgument{
					"email":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"username": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"avatar":   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"password": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"bio":      &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				},
				Resolve: CreateUserResolver,
			},
			// mutation: { addGirls: (cells: [{ img: "url", text: "hello", cate: 1, createdBy: hello }])}
			"addGirls": &graphql.Field{
				Type:        model.GirlGraphqlSchema,
				Description: "add some Girls",
				Args: graphql.FieldConfigArgument{
					"cells": &graphql.ArgumentConfig{Type: graphql.NewList(model.GirlInputSchema)},
				},
				Resolve: CreateGirl,
			},
		},
	})
}

// InitGraphQLSchema should init before app start
func InitGraphQLSchema() {

	rootQuery := getRootSchema()
	var err error

	GraphQLSchema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
		// TODO:
		Mutation: getRootMutation(),
	})
	if err != nil {
		revel.INFO.Println(err)
	}
	revel.INFO.Println(GraphQLSchema)
}

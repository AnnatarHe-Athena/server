package gql

import (
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
					"id": &graphql.ArgumentConfig{Type: graphql.Int},
				},
				Resolve: QueryUserResolver,
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
					"offset":   &graphql.ArgumentConfig{Type: graphql.Int},
					"take":     &graphql.ArgumentConfig{Type: graphql.Int},
					"from":     &graphql.ArgumentConfig{Type: graphql.Int},
					"hideOnly": &graphql.ArgumentConfig{Type: graphql.Boolean},
				},
				Resolve: GirlsResolver,
			},
			// 有 bug. ios 测出来的
			"collections": &graphql.Field{
				Type:        graphql.NewList(model.GirlGraphqlSchema),
				Description: "collections",
				Args: graphql.FieldConfigArgument{
					// from who
					"id": &graphql.ArgumentConfig{Type: graphql.Int},
					// from where(user.id)
					"from": &graphql.ArgumentConfig{Type: graphql.Int},
					// how much you want
					"size": &graphql.ArgumentConfig{Type: graphql.Int},
				},
				Resolve: QueryCollectionResolver,
			},
			"versions": &graphql.Field{
				Type:        graphql.NewList(model.MobAppVersionGraphQLSchema),
				Description: "versions",
				Args:        model.MobAppVersionGraphQLArgs,
				Resolve:     PlatformVersionResolver,
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
			// 有 bug. ios 测出来的
			// mutation: { addGirls: (cells: [{ img: "url", text: "hello", cate: 1, createdBy: hello }])}
			"addGirls": &graphql.Field{
				Type:        graphql.NewList(model.GirlGraphqlSchema),
				Description: "add some Girls",
				Args: graphql.FieldConfigArgument{
					"cells": &graphql.ArgumentConfig{Type: graphql.NewList(model.GirlInputSchema)},
				},
				Resolve: CreateGirl,
			},
			// TODO: add like button
			"addCollection": &graphql.Field{
				Type:        model.CollectionAddedReturnSchemaType,
				Description: "add collection",
				Args: graphql.FieldConfigArgument{
					// mutation: { addCollection: ( cells: [1,2,3] ) }
					"cells": &graphql.ArgumentConfig{Type: graphql.NewList(graphql.Int)},
				},
				Resolve: AddCollection,
			},
			"removeGirl": &graphql.Field{
				Type:        model.CollectionAddedReturnSchemaType,
				Description: "remove girl cell",
				Args: graphql.FieldConfigArgument{
					// mutation: { removeGirl: ( cells: [1,2,3] ) }
					"cells":    &graphql.ArgumentConfig{Type: graphql.NewList(graphql.Int)},
					"toRemove": &graphql.ArgumentConfig{Type: graphql.Boolean},
				},
				Resolve: RemoveGirl,
			},
		},
	})
}

// InitGraphQLSchema should init before app start
func InitGraphQLSchema() {
	var err error
	GraphQLSchema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query:    getRootSchema(),
		Mutation: getRootMutation(),
	})
	if err != nil {
		revel.INFO.Println(err)
	}
	revel.INFO.Println(GraphQLSchema)
}

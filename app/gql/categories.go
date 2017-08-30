package gql

import (
	"encoding/json"
	"time"

	"github.com/revel/revel"

	"github.com/douban-girls/server/app/initial"
	"github.com/douban-girls/server/app/model"
	"github.com/graphql-go/graphql"
)

// CategoriesResolver just a graphql resolver
func CategoriesResolver(params graphql.ResolveParams) (interface{}, error) {

	cates, err := initial.Redis.Get("girls:categories").Result()
	if err == nil {
		var v []model.Category
		if err := json.Unmarshal([]byte(cates), &v); err != nil {
			revel.INFO.Println(err)
		}
		return v, nil
	}

	categories, err := model.FetchAllCategories(initial.DB)

	go func() {
		cateJSON, err := json.Marshal(categories)
		if err != nil {
			revel.INFO.Println(err)
		}
		if err := initial.Redis.Set("girls:categories", cateJSON, time.Minute*5).Err(); err != nil {
			revel.INFO.Println("redis fail", err)
		}
	}()

	return categories, err
}

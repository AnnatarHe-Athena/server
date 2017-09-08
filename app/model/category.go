package model

import (
	"database/sql"

	"github.com/graphql-go/graphql"
)

type Category struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Src   int    `json:"src"`
	Count int    `json:"count"`
}

var CategoryGraphqlSchema = graphql.NewObject(graphql.ObjectConfig{
	Name: "category",
	Fields: graphql.Fields{
		"id":    &graphql.Field{Type: graphql.ID},
		"name":  &graphql.Field{Type: graphql.String},
		"src":   &graphql.Field{Type: graphql.Int},
		"count": &graphql.Field{Type: graphql.Int},
	},
})

func FetchAllCategories(db *sql.DB) ([]Category, error) {
	rows, err := db.Query(`SELECT categories.id, categories.name, categories.src, count(cells.id) AS count FROM categories, cells WHERE categories.id = cells.cate GROUP BY categories.id;`)
	defer rows.Close()
	categories := []Category{}
	if err != nil {
		return categories, err
	}

	for rows.Next() {
		var id, src, count int
		var name string
		rows.Scan(&id, &name, &src, &count)
		category := Category{
			ID:    id,
			Name:  name,
			Src:   src,
			Count: count,
		}
		categories = append(categories, category)
	}
	return categories, nil
}

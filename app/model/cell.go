package model

import (
	"database/sql"
	"fmt"

	"github.com/douban-girls/server/app/initial"
	"github.com/graphql-go/graphql"
)

type Cell struct {
	ID        int    `json:"id"`
	Img       string `json:"img"`
	Text      string `json:"text"`
	Cate      int    `json:"cate"`
	CreatedBy int    `json:"createdBy"`
}

var GirlInputSchema = graphql.NewObject(graphql.ObjectConfig{
	Name: "girl item",
	Fields: graphql.Fields{
		"img":       &graphql.Field{Type: graphql.String},
		"text":      &graphql.Field{Type: graphql.String},
		"cate":      &graphql.Field{Type: graphql.Int},
		"createdBy": &graphql.Field{Type: graphql.Int},
	},
})

var GirlGraphqlSchema = graphql.NewObject(graphql.ObjectConfig{
	Name: "girl",
	Fields: graphql.Fields{
		"id":        &graphql.Field{Type: graphql.ID},
		"img":       &graphql.Field{Type: graphql.String},
		"text":      &graphql.Field{Type: graphql.String},
		"cate":      &graphql.Field{Type: graphql.Int},
		"createdBy": &graphql.Field{Type: graphql.Int},
	},
})

type Cells []*Cell

func (cs Cells) Save(db *sql.DB) error {
	stat, err := db.Prepare("INSERT INTO cells(img, text, cate, createdBy) VALUES($1, $2, $3, $4) ON CONFLICT (img) DO NOTHING RETURNING id")
	if err != nil {
		return err
	}
	for _, cell := range cs {
		var id int
		err := stat.QueryRow(cell.Img, cell.Text, cell.Cate, cell.CreatedBy).Scan(&id)
		cell.ID = id
		if err != nil {
			fmt.Println(err)
		}
	}
	return nil
}

func fetchGilsFromDatabase(db *sql.DB, cate, row, offset int) ([]Cell, error) {
	rows, err := initial.DB.Query("SELECT id, text, img, cate FROM cells WHERE cate=$1 AND premission=2 ORDER BY id DESC LIMIT $2 OFFSET $3", cate, row, offset)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	result := []Cell{}

	for rows.Next() {
		var id int
		var text string
		var img string
		var cate int
		if err := rows.Scan(&id, &text, &img, &cate); err != nil {
			return nil, err
		}
		result = append(result, Cell{
			ID:   id,
			Img:  img,
			Text: text,
			Cate: cate,
		})
	}
	return result, nil
}

func FetchGirls(db *sql.DB, cate, row, offset int) ([]Cell, error) {
	// 需要保证返回的是最后几条数据，还没想好怎么存 redis 里面
	return fetchGilsFromDatabase(db, cate, row, offset)
}

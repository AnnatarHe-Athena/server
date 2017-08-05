package model

import (
	"database/sql"
	"fmt"

	"github.com/douban-girls/douban-girls-server/app/initial"
	"github.com/graphql-go/graphql"
)

type Cell struct {
	ID   int    `json:"id"`
	Img  string `json:"img"`
	Text string `json:"text"`
	Cate int    `json:"cate"`
}

var GirlGraphqlSchema = graphql.NewObject(graphql.ObjectConfig{
	Name: "girl",
	Fields: graphql.Fields{
		"id":   &graphql.Field{Type: graphql.ID},
		"img":  &graphql.Field{Type: graphql.String},
		"text": &graphql.Field{Type: graphql.String},
		"cate": &graphql.Field{Type: graphql.Int},
	},
})

type Cells []*Cell

func (cs Cells) Save(db *sql.DB) error {
	stat, err := db.Prepare("INSERT INTO cells(img, text, cate) VALUES($1, $2, $3) ON CONFLICT (img) DO NOTHING RETURNING id")
	if err != nil {
		return err
	}
	for _, cell := range cs {
		var id int
		err := stat.QueryRow(cell.Img, cell.Text, cell.Cate).Scan(&id)
		cell.ID = id
		if err != nil {
			fmt.Println(err)
		}
	}
	return nil
}

func FetchGirls(db *sql.DB, cate, row, offset int) ([]Cell, error) {
	rows, err := initial.DB.Query("SELECT * FROM cells WHERE cate=$1 ORDER BY id DESC LIMIT $2 OFFSET $3", cate, row, offset)
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
		if err := rows.Scan(&id, &img, &text, &cate); err != nil {
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

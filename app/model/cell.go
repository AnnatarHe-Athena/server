package model

import (
	"database/sql"

	"github.com/revel/revel"

	"github.com/douban-girls/server/app/initial"
	"github.com/douban-girls/server/app/utils"
	"github.com/graphql-go/graphql"
)

type Cell struct {
	ID         int    `json:"id"`
	Img        string `json:"img"`
	Text       string `json:"text"`
	Premission int    `json:"premission"`
	Cate       int    `json:"cate"`
	CreatedBy  int    `json:"createdBy"`
}

var GirlInputSchema = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "CellInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"img":        &graphql.InputObjectFieldConfig{Type: graphql.String},
		"text":       &graphql.InputObjectFieldConfig{Type: graphql.String},
		"cate":       &graphql.InputObjectFieldConfig{Type: graphql.Int},
		"premission": &graphql.InputObjectFieldConfig{Type: graphql.Int},
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
	stat, err := db.Prepare("INSERT INTO cells(img, text, cate, premission) VALUES($1, $2, $3, $4) ON CONFLICT (img) DO NOTHING RETURNING id")
	if err != nil {
		utils.Log("error when save cells", err)
		return err
	}
	for _, cell := range cs {
		var id int
		err := stat.QueryRow(cell.Img, cell.Text, cell.Cate, cell.Premission).Scan(&id)
		cell.ID = id
		revel.INFO.Println(*cell)
		if err != nil {
			utils.Log("error when save cells", err)
			return err
		}
	}
	return nil
}

func fetchGilsFromDatabase(db *sql.DB, cate, row, offset int) (Cells, error) {
	revel.INFO.Println("read from db")
	rows, err := initial.DB.Query("SELECT id, text, img, cate FROM cells WHERE cate=$1 AND premission=2 ORDER BY id DESC LIMIT $2 OFFSET $3", cate, row, offset)
	defer rows.Close()

	if err != nil {
		revel.INFO.Println("fetch girls from database error")
		revel.INFO.Println(err)
		return nil, err
	}

	result := GetCellsFromRows(rows)
	return result, nil
}

func GetCellsFromRows(rows *sql.Rows) (result Cells) {
	for rows.Next() {
		var id int
		var text string
		var img string
		var cate int
		if err := rows.Scan(&id, &text, &img, &cate); err != nil {
			revel.INFO.Println(err)
			return
		}
		result = append(result, &Cell{
			ID:   id,
			Img:  img,
			Text: text,
			Cate: cate,
		})
	}
	return
}

func FetchGirls(db *sql.DB, cate, row, offset int) (Cells, error) {
	// 需要保证返回的是最后几条数据，还没想好怎么存 redis 里面
	return fetchGilsFromDatabase(db, cate, row, offset)
}

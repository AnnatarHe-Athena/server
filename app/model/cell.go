package model

import (
	"database/sql"
	"time"

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
	FromID     string `json:"from_id"`
	FromURL    string `json:"from_url"`
	CreatedBy  int64  `json:"createdBy"`
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
		"id":         &graphql.Field{Type: graphql.ID},
		"img":        &graphql.Field{Type: graphql.String},
		"text":       &graphql.Field{Type: graphql.String},
		"cate":       &graphql.Field{Type: graphql.Int},
		"premission": &graphql.Field{Type: graphql.Int},
		"createdBy":  &graphql.Field{Type: graphql.Int},
		"from_id":    &graphql.Field{Type: graphql.String},
		"from_url":   &graphql.Field{Type: graphql.String},
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
	rows, err := initial.DB.Query("SELECT id, text, img, cate, premission, from_url, from_id, createdby FROM cells WHERE cate=$1 AND premission=2 ORDER BY id DESC LIMIT $2 OFFSET $3", cate, row, offset)
	defer rows.Close()

	if err != nil {
		revel.INFO.Println("fetch girls from database error")
		revel.INFO.Println(err)
		return nil, err
	}

	result := GetCellsFromRows(rows)
	return result, nil
}

func CellHideOrRemove(id int, shouldToRemove bool) {
	var rows *sql.Rows
	var err error
	if shouldToRemove {
		rows, err = initial.DB.Query("DELETE FROM cells WHERE id=$1", id)
	} else {
		rows, err = initial.DB.Query("UPDATE cells SET premission=3, updatedat=$1 WHERE id=$2", time.Now(), id)
	}
	if err != nil {
		revel.INFO.Println("error occean when cell to hide or remove", err)
	} else {
		rows.Close()
	}
}

func GetCellsFromRows(rows *sql.Rows) (result Cells) {
	for rows.Next() {
		var id, cate, premission int
		var text, img, fromID, fromURL, createdBy string
		if err := rows.Scan(&id, &text, &img, &cate, &premission, &fromURL, &fromID, &createdBy); err != nil {
			revel.INFO.Println(err)
			return
		}
		createdByUnix := getTimestamp(createdBy)
		result = append(result, &Cell{
			ID:         id,
			Img:        img,
			Text:       text,
			Premission: premission,
			Cate:       cate,
			FromID:     fromID,
			FromURL:    fromURL,
			CreatedBy:  createdByUnix,
		})
	}
	return
}

func FetchOneGirl(db *sql.DB, id int) *Cell {
	var cate, premission int
	var img, text, createdBy, fromURL, fromID string

	row := db.QueryRow("SELECT img ,text, cate, premission, from_url, from_id, createdby FROM cells WHERE id=$1 LIMIT 1", id)
	if err := row.Scan(&img, &text, &cate, &premission, &fromURL, &fromID, &createdBy); err != nil {
		utils.Log("i", err)
		return nil
	}

	createdByUnix := getTimestamp(createdBy)

	return &Cell{
		ID:         id,
		Img:        img,
		Text:       text,
		Cate:       cate,
		Premission: premission,
		FromURL:    fromURL,
		FromID:     fromID,
		CreatedBy:  createdByUnix,
	}
}

func getTimestamp(createdBy string) (createdByUnix int64) {
	timestamp, err := time.Parse(time.RFC3339, createdBy)
	if err != nil {
		createdByUnix = time.Now().Unix()
		utils.Log("i", err)
	} else {
		createdByUnix = timestamp.Unix()
	}
	return

}

func FetchGirls(db *sql.DB, cate, row, offset int) (Cells, error) {
	// 需要保证返回的是最后几条数据，还没想好怎么存 redis 里面
	return fetchGilsFromDatabase(db, cate, row, offset)
}

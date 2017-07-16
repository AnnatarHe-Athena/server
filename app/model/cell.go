package model

import (
	"database/sql"
	"fmt"
)

type Cell struct {
	ID   int    `json:"id"`
	Img  string `json:"img"`
	Text string `json:"text"`
	Cate int    `json:"cate"`
}

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

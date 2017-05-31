package controllers

import (
	"github.com/douban-girls/douban-girls-server/app"
	"github.com/douban-girls/douban-girls-server/app/model"
	"github.com/douban-girls/douban-girls-server/app/utils"
	"github.com/revel/revel"
)

// SELECT * FROM cells WHERE cate=$1 ORDER BY id DESC LIMIT $2 OFFSET $3

type Girls struct {
	*revel.Controller
}

func (g Girls) Get(cate, row, offset int) revel.Result {
	rows, err := app.DB.Query("SELECT * FROM cells WHERE cate=$1 ORDER BY id DESC LIMIT $2 OFFSET $3", cate, row, offset)
	defer rows.Close()

	if err != nil {
		return g.RenderJSON(utils.Response(500, nil, err))
	}

	result := []model.Cell{}
	for rows.Next() {
		var id int
		var text string
		var img string
		var cate int
		if err := rows.Scan(&id, &img, &text, &cate); err != nil {
			return g.RenderError(err)
		}
		result = append(result, model.Cell{
			Id:   id,
			Img:  img,
			Text: text,
			Cate: cate,
		})
	}
	return g.RenderJSON(utils.Response(200, result, nil))
}

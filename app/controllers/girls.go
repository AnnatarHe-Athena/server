package controllers

import (
	"github.com/douban-girls/douban-girls-server/app/initial"
	"github.com/douban-girls/douban-girls-server/app/model"
	"github.com/douban-girls/douban-girls-server/app/utils"
	"github.com/revel/revel"
)

// SELECT * FROM cells WHERE cate=$1 ORDER BY id DESC LIMIT $2 OFFSET $3

// Girls Controller
type Girls struct {
	*revel.Controller
}

// Get will return girls by params
func (g Girls) Get(cate, row, offset int) revel.Result {
	rows, err := initial.DB.Query("SELECT * FROM cells WHERE cate=$1 ORDER BY id DESC LIMIT $2 OFFSET $3", cate, row, offset)
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
			ID:   id,
			Img:  img,
			Text: text,
			Cate: cate,
		})
	}
	return g.RenderJSON(utils.Response(200, result, nil))
}

func (g Girls) GetCategories() revel.Result {
	categories, err := models.FetchAllCategories(initial.DB)
	if err != nil {
		return g.RenderJSON(utils.Response(500, nil, err))
	}
	return g.RenderJSON(utils.Response(200, categories, nil))
}

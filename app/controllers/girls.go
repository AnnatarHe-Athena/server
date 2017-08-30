package controllers

import (
	"github.com/douban-girls/server/app/initial"
	"github.com/douban-girls/server/app/model"
	"github.com/douban-girls/server/app/utils"
	"github.com/revel/revel"
)

// SELECT * FROM cells WHERE cate=$1 ORDER BY id DESC LIMIT $2 OFFSET $3

// Girls Controller
type Girls struct {
	*revel.Controller
}

// Get will return girls by params
func (g Girls) Get(cate, row, offset int) revel.Result {
	girls, err := model.FetchGirls(initial.DB, cate, row, offset)
	if err != nil {
		return g.RenderJSON(utils.Response(500, nil, err))
	}

	return g.RenderJSON(utils.Response(200, girls, nil))
}

func (g Girls) GetCategories() revel.Result {
	categories, err := model.FetchAllCategories(initial.DB)
	if err != nil {
		return g.RenderJSON(utils.Response(500, nil, err))
	}
	return g.RenderJSON(utils.Response(200, categories, nil))
}

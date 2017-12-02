package controllers

import (
	"github.com/douban-girls/server/app/utils"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Hello() revel.Result {
	return c.RenderText("hello")
}

func (c App) GetQiniuToken() revel.Result {
	result := map[string]string{
		"uptoken": utils.GenQiniuToken(),
	}
	return c.RenderJSON(result)
}

package controllers

import (
	"github.com/douban-girls/douban-girls-server/app/utils"
	"github.com/revel/revel"
)

type Collection struct {
	*revel.Controller
}

func (c Collection) T() revel.Result {
	return c.RenderJSON(utils.Response(200, nil, nil))
}

// AddToCollection will add imgID to users collection
func (c Collection) AddToCollection(ids []int) revel.Result {
	return c.RenderJSON(utils.Response(200, nil, nil))
}

// RemoveFromCollection will remove imgID from user collection
func (c Collection) RemoveFromCollection(ids []int) revel.Result {
	return c.RenderJSON(utils.Response(200, nil, nil))
}

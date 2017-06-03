package controllers

import (
	"github.com/douban-girls/douban-girls-server/app/initial"
	"github.com/douban-girls/douban-girls-server/app/utils"
	"github.com/revel/revel"
)

type Collection struct {
	*revel.Controller
}

// AddToCollection will add imgID to users collection
func (c Collection) AddToCollection(ids []int) revel.Result {
	uid := utils.GetUID(c.Request)
	stat, err := initial.DB.Prepare("INSERT INTO collections(cell, owner) VALUES($1, $2)")
	if err != nil {
		return c.RenderJSON(utils.Response(500, nil, err))
	}
	for id := range ids {
		_, err := stat.Exec(id, uid)
		if err != nil {
			return c.RenderJSON(utils.Response(500, nil, err))
		}
	}
	return c.RenderJSON(utils.Response(200, map[string]string{"message": "success"}, nil))
}

// RemoveFromCollection will remove imgID from user collection
func (c Collection) RemoveFromCollection(ids []int) revel.Result {
	stat, err := initial.DB.Prepare("DELETE FROM collections WHERE id=$1")
	if err != nil {
		return c.RenderJSON(utils.Response(500, nil, err))
	}

	for id := range ids {
		_, err := stat.Exec(id)
		if err != nil {
			return c.RenderJSON(utils.Response(500, nil, err))
		}
	}

	return c.RenderJSON(utils.Response(200, nil, nil))
}

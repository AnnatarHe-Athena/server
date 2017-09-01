package controllers

import (
	"github.com/douban-girls/server/app/initial"
	"github.com/douban-girls/server/app/model"
	"github.com/douban-girls/server/app/utils"
	"github.com/revel/revel"
)

type Collection struct {
	*revel.Controller
}

// AddToCollection will add imgID to users collection
func (c Collection) AddToCollection(ids string) revel.Result {
	intIds, err := utils.CutCommaAndTrimStrings(ids)
	if err != nil {
		return c.RenderJSON(utils.Response(400, nil, err))
	}
	uid := utils.GetUID(c.Request)

	var collections model.Collections
	for _, id := range intIds {
		collection := model.NewCollection(id, uid, -1)
		collections = append(collections, collection)
	}
	if err := collections.Save(initial.DB); err != nil {
		return c.RenderJSON(utils.Response(500, nil, err))
	}

	return c.RenderJSON(utils.Response(200, collections, nil))
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

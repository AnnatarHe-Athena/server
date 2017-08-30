package controllers

import (
	"github.com/douban-girls/server/app/utils"
	"github.com/revel/revel"
)

// Exceptionwill catch error
type ExceptionController struct {
	*revel.Controller
}

// TokenMiss return token miss message and it should login first
func (e ExceptionController) TokenMiss() revel.Result {
	return e.RenderJSON(utils.Response(400, "token miss", nil))
}

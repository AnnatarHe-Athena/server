package filters

import (
	"strings"

	"errors"

	"github.com/douban-girls/douban-girls-server/app/initial"
	"github.com/douban-girls/douban-girls-server/app/utils"
	"github.com/revel/revel"
)

// TokenFilter will check token first
func TokenFilter(c *revel.Controller, fc []revel.Filter) {
	token := c.Request.Header.Get("douban-girls-token")
	uri := c.Request.RequestURI
	if strings.HasPrefix(uri, "/api/collection") {
		if token == "" {
			err := errors.New("token miss")
			c.Result = c.RenderJSON(utils.Response(400, nil, err))
			return
		}
		uid := strings.Split(c.Request.Header.Get("douban-girls-token"), "|")[0]
		tokenInRedis := initial.Redis.Get("token:" + uid)
		if err := tokenInRedis.Err(); err != nil {
			c.Result = c.RenderJSON(utils.Response(403, nil, err))
			return
		}
		if tokenInRedis.Val() != token {
			err := errors.New("incorrect token")
			c.Result = c.RenderJSON(utils.Response(403, nil, err))
			return
		}
	}
	revel.INFO.Println(token, uri)
	revel.INFO.Println("token filter")
	fc[0](c, fc[1:])
}

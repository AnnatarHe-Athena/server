package filters

import (
	"github.com/revel/revel"
)

// TokenFilter will check token first
func TokenFilter(c *revel.Controller, fc []revel.Filter) {
	// token := c.Request.Header.Get("douban-girls-token")
	// uri := c.Request.RequestURI
	// if strings.HasPrefix(uri, "/api/collection") {
	// 	if token == "" {
	// 		err := errors.New("token miss")
	// 		c.Result = c.RenderJSON(utils.Response(400, nil, err))
	// 		return
	// 	}
	// 	uid := utils.GetUID(c.Request)
	// 	tokenInRedis := initial.Redis.Get("token:" + strconv.Itoa(uid))
	// 	if err := tokenInRedis.Err(); err != nil {
	// 		c.Result = c.RenderJSON(utils.Response(403, nil, err))
	// 		return
	// 	}
	// 	if tokenInRedis.Val() != token {
	// 		err := errors.New("incorrect token")
	// 		c.Result = c.RenderJSON(utils.Response(403, nil, err))
	// 		return
	// 	}
	// }
	fc[0](c, fc[1:])
}

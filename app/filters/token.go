package filters

import (
	"github.com/revel/revel"
)

func TokenFilter(c *revel.Controller, fc []revel.Filter) {
	token := c.Request.Header.Get("douban-girls-token")
	uri := c.Request.RequestURI
	revel.INFO.Println(token, uri)
	fc[0](c, fc[1:])
}

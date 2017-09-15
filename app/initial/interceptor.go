package initial

import (
	"github.com/revel/revel"
)

func interceptorCloseDB(c *revel.Controller) revel.Result {
	DB.Close()
	return nil
}

func RunInterceptors() {
	revel.InterceptFunc(interceptorCloseDB, revel.AFTER, &revel.Controller{})
}

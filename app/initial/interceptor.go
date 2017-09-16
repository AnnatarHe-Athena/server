package initial

import (
	"github.com/revel/revel"
)

// We should not manually close this for some reason:
// https://stackoverflow.com/questions/29063123/when-should-i-close-the-database-connection-in-this-simple-web-app
// https://github.com/revel/revel/issues/404
func interceptorCloseDB(c *revel.Controller) revel.Result {

	// revel.INFO.Println("will close db resource")
	// defer DB.Close()
	return nil
}

func RunInterceptors() {
	revel.InterceptFunc(interceptorCloseDB, revel.FINALLY, &revel.Controller{})
}

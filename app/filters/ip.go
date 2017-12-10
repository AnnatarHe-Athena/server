package filters

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/douban-girls/server/app/initial"
	"github.com/douban-girls/server/app/utils"

	"github.com/revel/revel"
)

// IPFilter limit api request
func IPFilter(c *revel.Controller, fc []revel.Filter) {
	if strings.Index(c.Request.URL.Path, "/graphql/v1") != 0 {
		fc[0](c, fc[1:])
		return
	}

	ip := strings.Split(c.Request.RemoteAddr, ":")[0]
	redisKey := "ip:" + ip + ":requested"

	count, err := initial.Redis.Get(redisKey).Result()

	if err != nil {
		// 说明可能是第一次登陆，加入 key
		if err := initial.Redis.Set(redisKey, 0, time.Hour*24).Err(); err != nil {
			revel.INFO.Println(err)
			c.Result = c.RenderJSON(utils.Response(500, nil, err))
			return
		}
		count = "0"
	}

	countInt, err := strconv.Atoi(count)
	if err != nil {
		c.Result = c.RenderJSON(utils.Response(500, nil, err))
		return
	}

	// 每个ip每天限制300个请求
	if countInt > 300 {
		err := errors.New("api request out of limit")
		c.Response.SetStatus(http.StatusTooManyRequests)
		c.Result = c.RenderJSON(utils.Response(http.StatusTooManyRequests, nil, err))
		return
	}

	go initial.Redis.Incr(redisKey)

	fc[0](c, fc[1:])
}

package utils

import (
	"errors"
	"strconv"
	"time"

	"golang.org/x/crypto/scrypt"

	"github.com/douban-girls/server/app/initial"
	"github.com/revel/revel"
)

// IsTokenPair will check token in header is pair token in redis or not
// if return false. DO NOT return correct result
// JUST FOR GraphQL arch
func IsTokenPair(c *revel.Controller) (bool, error) {
	token := c.Request.Header.Get("athena-token")

	userID := c.Session["id"]

	innerToken, err := initial.Redis.Get("token:" + userID).Result()

	if err != nil || token != innerToken {
		revel.INFO.Println(err)
		err403 := errors.New("login first please")
		return false, err403
	}

	return true, nil
}

func GenToken(id int) (string, error) {
	idStr := strconv.Itoa(int(id))

	token := idStr + "|" + Md5Encode(time.Now().Format("20060102150405"))
	go func() {
		timeout := time.Until(time.Now().AddDate(1, 0, 0))
		if err := initial.Redis.Set("token:"+idStr, token, timeout).Err(); err != nil {
			revel.INFO.Println("error when set token", err)
		}
	}()

	return token, nil
}

// GenPassword will return a very complex password
func GenPassword(pwd string) string {
	salt := revel.Config.StringDefault("salt", "default")
	realPasword, err := scrypt.Key([]byte(pwd), []byte(salt), 16384, 8, 1, 32)
	if err != nil {
		revel.INFO.Println("error in crypt password", err)
		return ""
	}
	return string(realPasword)
}

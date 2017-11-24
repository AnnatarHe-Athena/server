package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/scrypt"

	"github.com/douban-girls/server/app/initial"
	"github.com/revel/revel"
)

// IsTokenPair will check token in header is pair token in redis or not
// if return false. DO NOT return correct result
// JUST FOR GraphQL arch
func IsTokenPair(c *revel.Controller) (bool, error) {

	if revel.DevMode {
		return true, nil
	}

	token := c.Request.Header.Get("athena-token")

	userID := c.Session["userID"]

	// app 客户端并没有 session， 所以一定会报错。那么这里从 token 取得 userID
	if userID == "" {
		userID = strings.Split(token, "|")[0]
	}

	innerToken, err := initial.Redis.Get("token:" + userID).Result()

	revel.INFO.Println(innerToken, token)

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
		timeout := time.Duration(time.Minute * 60 * 24)
		if err := initial.Redis.Set("token:"+idStr, token, timeout).Err(); err != nil {
			revel.INFO.Println("error when set token", err)
		}
	}()

	return token, nil
}

// GenPassword will return a very complex password
func GenPassword(pwd string) string {
	return sha256Encode(pwd)
}

func sha256Encode(pwd string) string {
	h := sha256.New()
	io.WriteString(h, pwd)
	return hex.EncodeToString(h.Sum(nil))
}

func scryptEncode(pwd string) string {
	salt := revel.Config.StringDefault("salt", "default")
	realPasword, err := scrypt.Key([]byte(pwd), []byte(salt), 16384, 8, 1, 32)
	if err != nil {
		revel.INFO.Println("error in crypt password", err)
		return pwd
	}
	return string(realPasword)

}

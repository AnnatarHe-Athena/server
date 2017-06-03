package controllers

import (
	"database/sql"
	"strconv"
	"strings"
	"time"

	"github.com/douban-girls/douban-girls-server/app/initial"
	"github.com/douban-girls/douban-girls-server/app/utils"
	"github.com/revel/revel"
)

// Profile is user profile controller
type Profile struct {
	*revel.Controller
}

func (p Profile) saveToken(result sql.Result) (sql.Result, error) {
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	idStr := strconv.Itoa(int(id))

	token := idStr + "|" + utils.Md5Encode(time.Now().Format("20060102150405"))

	resultWithToken, err := initial.DB.Exec("UPDATE users SET token=$1 WHERE id=$2", token, id)

	go func() {
		timeout := time.Until(time.Now().AddDate(1, 0, 0))
		if err := initial.Redis.Set("token:"+idStr, token, timeout).Err(); err != nil {
			revel.INFO.Println(err)
		}
	}()
	if err != nil {
		return nil, err
	}
	return resultWithToken, nil
}

// Signup will save a user profile into database
func (p Profile) Signup() revel.Result {
	email := p.Params.Get("email")
	pwd := p.Params.Get("password")
	avatar := p.Params.Get("avatar")
	bio := p.Params.Get("bio")

	result, err := initial.DB.Exec("INSERT INTO users(email, pwd, avatar bio) VALUES($1, $2, $3, $4)", email, pwd, avatar, bio)
	if err != nil {
		return p.RenderJSON(utils.Response(500, nil, err))
	}
	resultWithToken, err := p.saveToken(result)
	if err != nil {
		return p.RenderJSON(utils.Response(500, nil, err))
	}
	return p.RenderJSON(utils.Response(200, resultWithToken, nil))
}

// Signin will return token
func (p Profile) Signin() revel.Result {
	email := p.Params.Get("email")
	pwd := p.Params.Get("password")
	result, err := initial.DB.Exec("SELECT * FROM users WHERE email=$1 AND pwd=$2", email, pwd)
	if err != nil {
		return p.RenderJSON(utils.Response(500, nil, err))
	}
	resultWithToken, err := p.saveToken(result)
	if err != nil {
		return p.RenderJSON(utils.Response(500, nil, err))
	}
	return p.RenderJSON(utils.Response(200, resultWithToken, nil))
}

// Logout will remove session
func (p Profile) Logout() revel.Result {
	uid := strings.Split(p.Request.Header.Get("douban-girls-token"), "|")[0]
	go initial.Redis.Del("token:" + uid)
	return p.RenderJSON(utils.Response(200, map[string]string{"message": "success"}, nil))
}

// Update user profile
func (p Profile) Update(uid int) revel.Result {
	// TODO:

	return p.RenderJSON(utils.Response(200, nil, nil))
}

// UserInfo will get user profile by id
func (p Profile) UserInfo(uid int) revel.Result {
	// TODO:
	return p.RenderJSON(utils.Response(200, nil, nil))
}

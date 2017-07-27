package controllers

import (
	"strconv"
	"time"

	"github.com/douban-girls/douban-girls-server/app/initial"
	"github.com/douban-girls/douban-girls-server/app/model"
	"github.com/douban-girls/douban-girls-server/app/utils"
	"github.com/revel/revel"
)

// Profile is user profile controller
type Profile struct {
	*revel.Controller
}

func (p Profile) getToken(id int) (string, error) {
	idStr := strconv.Itoa(int(id))

	token := idStr + "|" + utils.Md5Encode(time.Now().Format("20060102150405"))
	go func() {
		timeout := time.Until(time.Now().AddDate(1, 0, 0))
		if err := initial.Redis.Set("token:"+idStr, token, timeout).Err(); err != nil {
			revel.INFO.Println(err)
		}
	}()

	return token, nil
}

// Signup will save a user profile into database
func (p Profile) Signup() revel.Result {
	id := 0
	name := p.Params.Get("name")
	email := p.Params.Get("email")
	pwd := utils.Sha256Encode(p.Params.Get("password"))
	avatar := p.Params.Get("avatar")
	bio := p.Params.Get("bio")

	user := model.NewUser(id, email, name, pwd, avatar, bio, "")
	err := user.Save(initial.DB)
	if err != nil {
		return p.RenderJSON(utils.Response(500, nil, err))
	}
	token, err := p.getToken(user.ID)
	user.Token = token
	if err != nil {
		return p.RenderJSON(utils.Response(500, nil, err))
	}
	if err := user.Update(initial.DB); err != nil {
		return p.RenderJSON(utils.Response(500, nil, err))
	}
	return p.RenderJSON(utils.Response(200, user, nil))
}

// Signin will return user with token
func (p Profile) Signin() revel.Result {
	email := p.Params.Get("email")
	pwd := utils.Sha256Encode(p.Params.Get("password"))
	user, err := model.UserAuth(initial.DB, email, pwd)
	if err != nil {
		return p.RenderJSON(utils.Response(404, nil, err))
	}
	token, err := p.getToken(user.ID)
	user.Token = token
	if err := user.Update(initial.DB); err != nil {
		return p.RenderJSON(utils.Response(500, nil, err))
	}
	return p.RenderJSON(utils.Response(200, user, nil))
}

// Logout will remove session
func (p *Profile) Logout() revel.Result {
	uid := utils.GetUID(p.Request)
	go initial.Redis.Del("token:" + strconv.Itoa(uid))
	return p.RenderJSON(utils.Response(200, map[string]string{"message": "success"}, nil))
}

// Update user profile
func (p Profile) Update(uid int) revel.Result {
	// TODO:

	return p.RenderJSON(utils.Response(200, nil, nil))
}

// UserInfo will get user profile by id
func (p Profile) UserInfo(uid int) revel.Result {
	user, err := model.FetchUserBy(initial.DB, uid)
	if err != nil {
		return p.RenderJSON(utils.Response(200, nil, err))
	}
	return p.RenderJSON(utils.Response(200, user, nil))
}

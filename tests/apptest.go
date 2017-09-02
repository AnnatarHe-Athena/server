package tests

import (
	"github.com/revel/revel/testing"
)

type AppTest struct {
	testing.TestSuite
	token string
}

func (t *AppTest) Before() {
	// 登录，获取token
	println("Set up")
}

func (t *AppTest) TestThatIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t *AppTest) After() {
	println("Tear down")
}

package controllers

import (
	"github.com/revel/revel"
)

type Profile struct {
	*revel.Controller
}

func (p Profile) Register() revel.Result {

}

func (p Profile) Login() revel.Result {

}

func (p Profile) Logout() revel.Result {

}

func (p Profile) Update() revel.Result {

}

func (p Profile) AddCollection(ids []int) revel.Result {

}

func (p Profile) RemoveCollection(ids []int) revel.Result {

}

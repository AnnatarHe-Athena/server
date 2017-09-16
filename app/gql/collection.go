package gql

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"errors"
	"strconv"

	"github.com/douban-girls/server/app/initial"
	"github.com/douban-girls/server/app/model"
	"github.com/douban-girls/server/app/utils"
	"github.com/graphql-go/graphql"
	"github.com/revel/revel"
)

// QueryCollectionResolver will return collection by user
func QueryCollectionResolver(params graphql.ResolveParams) (interface{}, error) {
	isPair, err := utils.IsTokenPair(utils.GetController(params))
	if !isPair || err != nil {
		return nil, errors.New("token not pair")
	}

	userID := params.Args["id"].(int)
	from := params.Args["from"].(int)
	size := params.Args["size"].(int)
	if size > 50 {
		return nil, errors.New("max size is 50")
	}
	collections, err := model.FetchUserCollectionBy(initial.DB, userID, from, size)

	if err != nil {
		revel.INFO.Println(err)
		return nil, err
	}

	return collections, nil
}

func AddCollection(params graphql.ResolveParams) (interface{}, error) {
	controller := utils.GetController(params)
	isPair, err := utils.IsTokenPair(controller)
	if !isPair || err != nil {
		return nil, errors.New("token not pair")
	}

	userID, err := strconv.Atoi(controller.Session["userID"])
	if err != nil {
		return nil, err
	}

	cellsStructed := bytes.Buffer{}
	var cellIDs []int
	gob.NewDecoder(&cellsStructed).Decode(params.Args["cells"])
	if err := json.Unmarshal(cellsStructed.Bytes(), &cellIDs); err != nil {
		revel.INFO.Println("error when parse the girls collection list", err)
		return nil, err
	}

	var userIDs []int
	var fakeIDs []int

	for i := 0; i < 11; i++ {
		userIDs = append(userIDs, userID)
		fakeIDs = append(fakeIDs, i)
	}

	if err := model.NewCollections(fakeIDs, cellIDs, userIDs).Save(initial.DB); err != nil {
		revel.INFO.Println("error occean in save collection step")
		return nil, err
	}
	return true, nil
}

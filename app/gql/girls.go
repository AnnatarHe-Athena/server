package gql

import (
	"encoding/json"
	"errors"

	"github.com/revel/revel"

	"github.com/douban-girls/server/app/initial"
	"github.com/douban-girls/server/app/model"
	"github.com/douban-girls/server/app/utils"
	"github.com/graphql-go/graphql"
)

// GirlsResolver is graphql resolver
func GirlsResolver(params graphql.ResolveParams) (interface{}, error) {
	revel.INFO.Println("in girls resolver")

	isPair, err := utils.IsTokenPair(utils.GetController(params))
	if !isPair || err != nil {
		return nil, errors.New("token not pair")
	}
	// TODO: add redis cache here
	from := params.Args["from"].(int)
	take := params.Args["take"].(int)
	offset := params.Args["offset"].(int)
	girls, err := model.FetchGirls(initial.DB, from, take, offset)

	return girls, err
}

func getGirlArgument(params graphql.ResolveParams, keys []string) (result map[string]string) {
	for _, val := range keys {
		result[val] = params.Args[val].(string)
	}
	return result
}

type customCell struct {
	Cells []struct {
		Data map[string]string
		// img  string
		// id   int
		// text string
	}
}

// CreateGirl will set a girl to database
func CreateGirl(params graphql.ResolveParams) (interface{}, error) {

	// FIXME: 数据结构有问题，需要重新做 array 分析
	var cellsData customCell
	cells, e := json.Marshal(map[string]interface{}{"cells": params.Args["cells"]})
	if e != nil {
		revel.INFO.Println(e)
	}
	revel.INFO.Println(string(cells))
	if err := json.Unmarshal(cells, &cellsData); err != nil {
		revel.INFO.Println(params.Args["cells"])
		revel.INFO.Println("error when parse the girls cell list", err)
		return nil, err
	}

	var resolvedCells model.Cells
	revel.INFO.Println(cellsData)

	for val := range cellsData.Cells {
		revel.INFO.Println(val, cellsData.Cells, cellsData.Cells[val], cellsData.Cells[val].Data)
		cell := &model.Cell{
		// Img:  cellsData.Cells[val].data["img"].(string),
		// Text: cellsData.Cells[val].data["text"].(string),
		// Cate: cellsData.Cells[val].data["cate"].(int),
		}
		resolvedCells = append(resolvedCells, cell)
	}

	if err := resolvedCells.Save(initial.DB); err != nil {
		revel.INFO.Println("error when save girls cell list:", err)
		return nil, err
	}
	return cellsData, nil
}

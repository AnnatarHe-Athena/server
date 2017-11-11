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

	revel.INFO.Println(params.Args["cells"])

	var resolvedCells model.Cells
	cellsInterface := params.Args["cells"].([]interface{})
	cellsByte, err := json.Marshal(cellsInterface)
	if err != nil {
		utils.Log("marshal the input json", err)
		return resolvedCells, err
	}

	if err := json.Unmarshal(cellsByte, &resolvedCells); err != nil {
		utils.Log("unmarshal the input json to Data.Cells", err)
		return resolvedCells, err
	}

	if err := resolvedCells.Save(initial.DB); err != nil {
		utils.Log("error occean when save cells", err)
		return nil, err
	}

	var cellIDs []int
	var uid = utils.GetUserIDFromSession(params)

	for _, cell := range resolvedCells {
		cellIDs = append(cellIDs, cell.ID)
	}

	// can be set to goroutine
	if err := model.NewCollectionJustCell(cellIDs, uid).Save(initial.DB); err != nil {
		utils.Log("error when user save to collection", err)
		return nil, err
	}

	return resolvedCells, nil
}

func RemoveGirl(params graphql.ResolveParams) (interface{}, error) {

	revel.INFO.Println(params.Args["cells"])

	controller := utils.GetController(params)
	isPair, err := utils.IsTokenPair(controller)
	if !isPair || err != nil {
		return nil, errors.New("token not pair")
	}

	if err != nil {
		return nil, err
	}

	var cellIDs []int

	cellsInterface := params.Args["cells"].([]interface{})
	revel.INFO.Println("cellInterface", cellsInterface)
	cellsByte, err := json.Marshal(cellsInterface)
	if err != nil {
		utils.Log("marshal the input json", err)
		return cellIDs, err
	}

	revel.AppLog.Info(string(cellsByte))

	if err := json.Unmarshal(cellsByte, &cellIDs); err != nil {
		utils.Log("unmarshal the input json to Data.Cells", err)
		return cellIDs, err
	}

	go func() {
		for _, cellID := range cellIDs {
			revel.INFO.Println(cellID)
			model.CellHide(cellID)
		}
	}()

	isOk := okReturn{IsOk: true}

	return isOk, nil
}

type okReturn struct {
	IsOk bool `json:"isOk"`
}

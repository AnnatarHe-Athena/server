package gql

import (
	"bytes"
	"encoding/gob"
	"encoding/json"

	"github.com/revel/revel"

	"github.com/douban-girls/server/app/initial"
	"github.com/douban-girls/server/app/model"
	"github.com/graphql-go/graphql"
)

// GirlsResolver is graphql resolver
func GirlsResolver(params graphql.ResolveParams) (interface{}, error) {
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

// CreateGirl will set a girl to database
func CreateGirl(params graphql.ResolveParams) (interface{}, error) {

	paramsStructed := bytes.Buffer{}
	cellsData := model.Cells{}
	gob.NewDecoder(&paramsStructed).Decode(params.Args["cells"])
	if err := json.Unmarshal(paramsStructed.Bytes(), &cellsData); err != nil {
		revel.INFO.Println("error when parse the girls cell list", err)
		return nil, err
	}

	if err := cellsData.Save(initial.DB); err != nil {
		revel.INFO.Println("error when save girls cell list:", err)
		return nil, err
	}
	return cellsData, nil
}

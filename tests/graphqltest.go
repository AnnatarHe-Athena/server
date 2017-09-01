package tests

import (
	"bytes"
	"encoding/json"

	"github.com/revel/revel"
)

// 写post内容太麻烦了，不写了
func (t AppTest) TestGraphqlCategories() {
	postDataStr := `{
		"query": "query fetchCategories{ categories { id, name, src }}",
		"variables": null,
		"operationName": ""
	}`

	postData := new(bytes.Buffer)
	json.NewEncoder(postData).Encode(postDataStr)

	t.Post("/graphql/v1", "application/json", postData)
	revel.INFO.Println(string(t.ResponseBody))
	t.AssertOk()
}

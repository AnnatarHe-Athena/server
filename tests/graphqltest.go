package tests

import (
	"bytes"
	"encoding/json"

	"github.com/revel/revel"
)

func renderPostData(body string) *bytes.Buffer {
	postDataStr := []byte(body)

	postData := new(bytes.Buffer)
	json.NewEncoder(postData).Encode(postDataStr)
	return bytes.NewBuffer(postDataStr)
}

// 写post内容太麻烦了，不写了
func (t AppTest) TestGraphqlGetGirlsNeedToken() {
	// postDataStr := renderPostData("" +
	// 	"{\"query\":\"query fetchGirls($from: Int!, $take: Int!, $offset: Int!) {" +
	// 	"girls(from: $from, take: $take, offset: $offset) {    id    img    text    __typename  }}\"," +
	// 	"\"variables\":{\"from\":-1,\"take\":20,\"offset\":0},\"operationName\":\"fetchGirls\"}")
	postDataStr := renderPostData(`
		{
			"query": "query fetchGirls($from:Int!, $take:Int!, $offset:Int) {
				girls($from:Int!, $take:Int!, $offset:Int) { id, img, text }
			}",
			"variable": {"from": -1, "take": 20, "offset": 0}
		}`)

	t.Post("/graphql/v1", "application/json", postDataStr)
	revel.INFO.Println(string(t.ResponseBody))
	t.AssertNotContains("\"errors\": ")
	t.AssertOk()
}

func (t AppTest) TestGraphqlGetGirlsCorrect() {
	postDataStr := renderPostData("" +
		"{\"query\":\"query fetchCategories {" +
		"categories {    id    name		src 	count    __typename  }}\"," +
		"\"operationName\":\"fetchCategories\"}")
	t.Post("/graphql/v1", "application/json", postDataStr)
	revel.INFO.Println(string(t.ResponseBody))
	t.AssertNotContains("\"errors\": ")
	t.AssertOk()
}

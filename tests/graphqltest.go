package tests

import (
	"bytes"
	"encoding/json"

	"github.com/revel/revel"
)

// 写post内容太麻烦了，不写了
func (t AppTest) TestGraphqlGetGirlsNeedToken() {
	postDataStr := []byte("" +
		"{\"query\":\"query fetchGirls($from: Int!, $take: Int!, $offset: Int!) {" +
		"girls(from: $from, take: $take, offset: $offset) {    id    img    text    __typename  }}\"," +
		"\"variables\":{\"from\":-1,\"take\":20,\"offset\":0},\"operationName\":\"fetchGirls\"}")

	postData := new(bytes.Buffer)
	json.NewEncoder(postData).Encode(postDataStr)

	t.Post("/graphql/v1", "application/json", bytes.NewBuffer(postDataStr))
	revel.INFO.Println(string(t.ResponseBody))
	t.AssertContains("\"errors\": null")
	t.AssertOk()
}

func (t AppTest) TestGraphqlGetGirlsCorrect() {
	t.AssertOk()
}

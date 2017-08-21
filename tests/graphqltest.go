package tests

import (
	"strings"

	"github.com/revel/revel"
)

// 	`"query":"query fetchGirls($from: Int!, $take: Int!, $offset: Int!) {\n  gi
// rls(from: $from, take: $take, offset: $offset) {\n    id\n    img\n    text\n    __typename\n  }\n  categories {\n
//   id\n    name\n    src\n    __typename\n  }\n}\n","variables":{"from":0,"take":0,"offset":0},"operationName":"fetch
// Girls"`

// 写post内容太麻烦了，不写了
func (t AppTest) TestGraphqlCategories() {
	postData := strings.NewReader("\"query\"=\"query fetchCategories{ categories { id, name, src }}\"")
	t.Post("/graphql/v1", "text/json", postData)
	revel.INFO.Println(string(t.ResponseBody))
	t.AssertOk()
}

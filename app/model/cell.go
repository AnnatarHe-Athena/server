package model

type Cell struct {
	ID   int    `json:"id"`
	Img  string `json:"img"`
	Text string `json:"text"`
	Cate int    `json:"cate"`
}

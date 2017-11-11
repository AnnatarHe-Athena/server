package model

import (
	"database/sql"

	"github.com/revel/revel"

	"github.com/graphql-go/graphql"
)

type Collection struct {
	ID    int `json:"id"`
	Cell  int `json:"cell"`
	Owner int `json:"owner"`
}

var CollectionGraphQLSchema = graphql.NewObject(graphql.ObjectConfig{
	Name: "collectionItem",
	Fields: graphql.Fields{
		"id":    &graphql.Field{Type: graphql.ID},
		"cell":  &graphql.Field{Type: graphql.Int},
		"owner": &graphql.Field{Type: graphql.Int},
	},
})

var CollectionAddedReturnSchemaType = graphql.NewObject(graphql.ObjectConfig{
	Name: "addCollectionReturnType",
	Fields: graphql.Fields{
		"isOk": &graphql.Field{Type: graphql.Boolean},
	},
})

type Collections []*Collection

func NewCollection(cell, owner, id int) *Collection {
	return &Collection{
		ID:    id,
		Cell:  cell,
		Owner: owner,
	}
}

func NewCollectionJustCell(cells []int, owner int) Collections {
	var collections Collections
	for index := range cells {
		collection := &Collection{Cell: cells[index], Owner: owner}
		collections = append(collections, collection)
	}
	return collections

}

func NewCollections(ids, cells, owners []int) Collections {
	var collections Collections
	for index := range ids {
		collection := &Collection{ID: ids[index], Cell: cells[index], Owner: owners[index]}
		collections = append(collections, collection)
	}
	return collections
}

func (cs Collections) Save(db *sql.DB) error {
	stat, err := db.Prepare("INSERT INTO collections(cell, owner) VALUES($1, $2) RETURNING id")
	if err != nil {
		return err
	}
	for _, collection := range cs {
		var id int
		err := stat.QueryRow(collection.Cell, collection.Owner).Scan(&id)
		collection.ID = id
		if err != nil {
			return err
		}
	}
	return nil
}

func FetchUserCollectionBy(db *sql.DB, id, from, size int) (Cells, error) {
	rows, err := db.Query("SELECT cells.id, cells.text, cells.img, cells.cate FROM cells, collections WHERE collections.cell = cells.id AND collections.owner = $1 AND cells.id > $2 LIMIT $3", id, from, size)

	if err != nil {
		revel.INFO.Println(err)
		return nil, err
	}

	collections := GetCellsFromRows(rows)

	return collections, err
}

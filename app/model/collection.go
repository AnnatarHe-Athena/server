package model

import (
	"database/sql"

	"github.com/graphql-go/graphql"
)

type Collection struct {
	ID    int `json:"id"`
	Cell  int `json:"cell"`
	Owner int `json:"owner"`
}

var CollectionGraphQLSchema = graphql.NewObject(graphql.ObjectConfig{
	Name: "collection Item",
	Fields: graphql.Fields{
		"id":    &graphql.Field{Type: graphql.ID},
		"cell":  &graphql.Field{Type: graphql.Int},
		"owner": &graphql.Field{Type: graphql.Int},
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

func FetchUserCollectionBy(db *sql.DB, id int) (Collections, error) {
	rows, err := db.Query("SELECT DISTINCT ON (collections.cell) id, cell, owner FROM users LEFT JOIN collections ON users.id=collections.owner WHERE users.id=$1", id)

	var collections Collections

	for rows.Next() {
		var id, cell, owner int
		if err := rows.Scan(&id, &cell, &owner); err != nil {
			return collections, err
		}

		collections = append(collections, NewCollection(cell, owner, id))
	}

	return collections, err
}

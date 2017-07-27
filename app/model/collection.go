package model

import (
	"database/sql"
)

type Collection struct {
	ID    int `json:"id"`
	Cell  int `json:"cell"`
	Owner int `json:"owner"`
}

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

package model

import (
	"database/sql"

	"github.com/graphql-go/graphql"
)

// id SERIAL PRIMARY KEY,
// platform VARCHAR(32) NOT NULL DEFAULT '',
// version INTEGER NOT NULL DEFAULT 0,
// published_by VARCHAR(32) NOT NULL DEFAULT '',
// link VARCHAR(255) NOT NULL DEFAULT '',
// description TEXT NOT NULL DEFAULT '',
// title VARCHAR(32) NOT NULL DEFAULT ''
type MobAppVersion struct {
	ID          int    `json:"id"`
	Platform    string `json:"platform"`
	Version     int    `json:"version"`
	PublishedBy string `json:"publishedBy"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Title       string `json:"title"`
}

// MobAppVersionGraphQLSchema for mobile version check
var MobAppVersionGraphQLSchema = graphql.NewObject(graphql.ObjectConfig{
	Name: "versions",
	Fields: graphql.Fields{
		"id":          &graphql.Field{Type: graphql.ID},
		"platform":    &graphql.Field{Type: graphql.String},
		"version":     &graphql.Field{Type: graphql.Int},
		"publishedBy": &graphql.Field{Type: graphql.String},
		"link":        &graphql.Field{Type: graphql.String},
		"description": &graphql.Field{Type: graphql.String},
		"title":       &graphql.Field{Type: graphql.String},
	},
})

// MobAppVersionGraphQLArgs is arguments of version
var MobAppVersionGraphQLArgs = graphql.FieldConfigArgument{
	"platform":   &graphql.ArgumentConfig{Type: graphql.String},
	"getLastOne": &graphql.ArgumentConfig{Type: graphql.Boolean},
}

func getVersionFromDatabase(rows *sql.Rows) (versions []MobAppVersion, err error) {
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id, version int
		var platform, publishedBy, link, description, title string
		rows.Scan(&id, &platform, &version, &publishedBy, &link, &description, &title)
		versionInstance := MobAppVersion{
			ID:          id,
			Platform:    platform,
			Version:     version,
			PublishedBy: publishedBy,
			Link:        link,
			Description: description,
			Title:       title,
		}
		versions = append(versions, versionInstance)
	}

	return versions, nil

}

const baseFetchSQL = "SELECT id, platform, version, published_by, link, descriptions, title from versions"

// FetchAllVersions : fetch all versions from database
func FetchAllVersions(db *sql.DB) (versions []MobAppVersion, err error) {
	rows, e := db.Query(baseFetchSQL)
	if e != nil {
		return nil, e
	}
	return getVersionFromDatabase(rows)
}

// FetchPlatformSpecialOne : has where conditions
func FetchPlatformSpecialOne(db *sql.DB, platform string, getLastOne bool) (versions []MobAppVersion, err error) {
	fetchSQL := baseFetchSQL + " where platform=? ORDER BY createdat DESC LIMIT ?"
	limit := 1
	if !getLastOne {
		limit = 100
	}
	rows, e := db.Query(fetchSQL, platform, limit)
	if e != nil {
		return nil, e
	}
	return getVersionFromDatabase(rows)

}

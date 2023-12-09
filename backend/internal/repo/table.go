package repo

import (
	"fmt"

	"github.com/koki-develop/lgtmgen/backend/internal/env"
)

type dynamoDBTables struct {
	Categories categoriesTable
}

type categoriesTable struct {
	Name    string
	Indexes categoriesTableIndexes
	Columns categoriesTableColumns
}

type categoriesTableIndexes struct {
	IndexByLang indexByLang
}

type indexByLang struct {
	Name    string
	HashKey string
}

type categoriesTableColumns struct {
	Name  string
	Lang  string
	Count string
}

var tables = dynamoDBTables{
	Categories: categoriesTable{
		Name: fmt.Sprintf("lgtmgen-%s-categories", env.Vars.Stage),
		Indexes: categoriesTableIndexes{
			IndexByLang: indexByLang{
				Name:    "index_by_lang",
				HashKey: "lang",
			},
		},
	}}

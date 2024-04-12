package db

import (
	"context"
	"log"

	"github.com/ilcm96/dku-aegis-library/ent"
)

var Client *ent.Client

func InitDB() {
	var err error
	Client, err = ent.Open("sqlite3", "file:sqlite3.db?cache=shared&_fk=1")
	if err != nil {
		log.Panic(err)
	}
	err = Client.Schema.Create(context.Background())
	if err != nil {
		log.Panic(err)
	}
}

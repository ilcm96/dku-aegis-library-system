package db

import (
	"context"
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/ilcm96/dku-aegis-library/ent"
	"log"
)

var Client *ent.Client

func open(databaseUrl string) *ent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

func InitDB() {
	var err error
	Client = open("postgresql://dku:dku@dku-postgresql:5432/dku?sslmode=disable")
	err = Client.Schema.Create(context.Background())
	if err != nil {
		log.Panic(err)
	}
}

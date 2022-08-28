package store

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/tgs266/neon/neon/store/entities"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type postgresStore struct {
	db *bun.DB
}

var store = &postgresStore{}

func CreateStore(host string, username string, password string) {

	dsn := fmt.Sprintf("postgres://%s:%s@%s/neon?sslmode=disable", username, password, host)
	fmt.Println(dsn)
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())

	store.db = db

	err := db.ResetModel(context.TODO(), (*entities.Product)(nil))
	fmt.Println(err)
}

func IsConnected() bool {
	return store.db != nil
}

func InsertProduct(product entities.Product) error {
	_, err := store.db.NewInsert().Model(&product).Exec(context.TODO())
	return err
}

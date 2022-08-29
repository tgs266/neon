package store

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/tgs266/neon/neon/store/entities"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

type postgresStore struct {
	db *bun.DB
}

var store = &postgresStore{}

func CreateStore(host string, username string, password string) {

	dsn := fmt.Sprintf("postgres://%s:%s@%s/neon?sslmode=disable", username, password, host)
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())

	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	store.db = db

	if err := db.ResetModel(context.TODO(), (*entities.Product)(nil)); err != nil {
		panic(err)
	}
	if err := db.ResetModel(context.TODO(), (*entities.Release)(nil)); err != nil {
		panic(err)
	}
	if err := db.ResetModel(context.TODO(), (*entities.Install)(nil)); err != nil {
		panic(err)
	}
	if err := db.ResetModel(context.TODO(), (*entities.App)(nil)); err != nil {
		panic(err)
	}

}

func IsConnected() bool {
	return store.db != nil
}

func Insert[T any](item T) error {
	_, err := store.db.NewInsert().Model(&item).Exec(context.TODO())
	return err
}

func List[T any]() ([]T, error) {
	items := []T{}
	err := store.db.NewSelect().
		Model(&items).
		Scan(context.TODO())
	return items, err
}

func GetProduct(query string, args ...interface{}) (entities.Product, error) {
	var item entities.Product
	err := store.db.NewSelect().
		Model(&item).
		Where(query, args...).
		Relation("Releases", func(sq *bun.SelectQuery) *bun.SelectQuery {
			return sq.OrderExpr("string_to_array(\"r\".\"product_version\", '.')::int[] DESC")
		}).
		Scan(context.TODO())
	return item, err
}

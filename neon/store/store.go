package store

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/tgs266/neon/neon/store/entities"
	"github.com/tgs266/neon/neon/store/repositories"
	"github.com/tgs266/neon/neon/utils"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type postgresStore struct {
	db *bun.DB
}

var RELEASE_CHANNELS = []entities.ReleaseChannel{
	{Name: "DEV", Value: 0},
	{Name: "RELEASE_CANDIDATE", Value: 1},
	{Name: "RELEASE", Value: 2},
}

var store = &postgresStore{}

func CreateStore(host string, username string, password string, reset bool) {

	dsn := fmt.Sprintf("postgres://%s:%s@%s/neon?sslmode=disable", username, password, host)
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())

	// db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	store.db = db

	if reset {
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
		if err := db.ResetModel(context.TODO(), (*entities.ReleaseChannel)(nil)); err != nil {
			panic(err)
		}
		if err := db.ResetModel(context.TODO(), (*entities.QueuedChange)(nil)); err != nil {
			panic(err)
		}
		if err := db.ResetModel(context.TODO(), (*entities.StoredChange)(nil)); err != nil {
			panic(err)
		}
		if err := db.ResetModel(context.TODO(), (*entities.Credentials)(nil)); err != nil {
			panic(err)
		}
		db.NewInsert().Model(&RELEASE_CHANNELS).Exec(context.TODO())

		utils.WriteKey()
	}
	// just make sure we can read the key - this will panic
	utils.ReadKey()
}

func IsConnected() bool {
	return store.db != nil
}

func Count[T any](query string, args ...interface{}) (int, error) {
	res, err := store.db.NewSelect().Model(new(T)).Where(query, args...).Count(context.TODO())
	return res, err
}

func Insert[T any](item T) error {
	_, err := store.db.NewInsert().Model(&item).Exec(context.TODO())
	return err
}

func StoredChangeRepository() repositories.StoredChangeRepository {
	return repositories.StoredChangeRepository{DB: store.db}
}

func QueuedChangeRepository() repositories.QueuedChangeRepository {
	return repositories.QueuedChangeRepository{DB: store.db}
}

func AppRepository() repositories.AppRepository {
	return repositories.AppRepository{DB: store.db}
}

func ProductRepository() repositories.ProductRepository {
	return repositories.ProductRepository{DB: store.db}
}

func ReleaseRepository() repositories.ReleaseRepository {
	return repositories.ReleaseRepository{DB: store.db}
}

func ReleaseChannelRepository() repositories.ReleaseChannelRepository {
	return repositories.ReleaseChannelRepository{DB: store.db}
}

func InstallRepository() repositories.InstallRepository {
	return repositories.InstallRepository{DB: store.db}
}

func CredentialsRepository() repositories.CredentialsRepository {
	return repositories.CredentialsRepository{DB: store.db}
}

func List[T any]() ([]T, error) {
	items := []T{}
	err := store.db.NewSelect().
		Model(&items).
		Scan(context.TODO())
	return items, err
}

func PullProducts(productNames []string, releaseChannel int) ([]entities.Product, error) {
	items := []entities.Product{}
	err := store.db.NewSelect().
		Model(&items).
		Where("name IN (?)", bun.In(productNames)).
		Relation("Releases").
		// Relation("Releases", func(sq *bun.SelectQuery) *bun.SelectQuery {
		// 	return sq.Where("release_channel >= (?)", releaseChannel).OrderExpr("string_to_array(\"r\".\"product_version\", '.')::int[] DESC")
		// }).
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

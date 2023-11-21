package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/landucci/go-exagonal/adapters/db"
	"github.com/landucci/go-exagonal/application"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products (
		"id" string,
		"name" srting,
		"price" float,
		"status" string
	);`

	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `INSERT INTO products VALUES ("1", "Master Sword", 5000.0, "enabled")`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)
	product, err := productDb.Get("1")

	require.Nil(t, err)
	require.Equal(t, "Master Sword", product.GetName())
	require.Equal(t, 5000.0, product.GetPrice())
	require.Equal(t, "enabled", product.GetStatus())
}

func TestProductDb_Save(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)

	product := application.NewProduct()
	product.Name = "Hylian Shield"
	product.Price = 3000.0
	product.Status = "enabled"
	productResult, err := productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())

	product.Name = "Deku Seed"
	product.Price = 0.0
	product.Status = "desabled"
	productResult, err = productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())

}

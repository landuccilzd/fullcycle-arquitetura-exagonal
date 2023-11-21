package db

import (
	"database/sql"

	"github.com/landucci/go-exagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db: db}
}

func (p *ProductDb) Get(id string) (application.ProductInterface, error) {
	var product application.Product
	stmt, err := p.db.Prepare("select id, name, price, status from products where id = ?")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *ProductDb) Save(product application.ProductInterface) (application.ProductInterface, error) {
	var rows int
	var err error
	p.db.QueryRow("select id from products where id = ?", product.GetId()).Scan(&rows)

	if rows == 0 {
		_, err = p.create(product)
	} else {
		_, err = p.update(product)
	}

	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductDb) create(product application.ProductInterface) (application.ProductInterface, error) {
	insert := `INSERT INTO products (id, name, price, status) VALUES (?, ?, ?, ?)`
	stmt, err := p.db.Prepare(insert)
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(product.GetId(), product.GetName(), product.GetPrice(), product.GetStatus())
	if err != nil {
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductDb) update(product application.ProductInterface) (application.ProductInterface, error) {
	update := `UPDATE products SET name = ?, price = ?, status = ? WHERE id = ?`
	stmt, err := p.db.Prepare(update)
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(product.GetName(), product.GetPrice(), product.GetStatus(), product.GetId())
	if err != nil {
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		return nil, err
	}

	return product, nil
}

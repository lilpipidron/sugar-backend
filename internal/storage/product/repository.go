package product

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/lilpipidron/sugar-backend/internal/models/products"
)

type Repository interface {
	AddProduct(product products.Product) error
	GetProductsWithValueInName(value string) ([]*products.Product, error)
	GetBreadUnitAmount(name string) (int, error)
}

type repository struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) *repository {
	return &repository{DB: db}
}

func (db *repository) AddProduct(product products.Product) error {
	const op = "storage.product.AddProduct"

	query := "INSERT INTO products (product_name, carbs) VALUES ($1, $2)"
	_, err := db.DB.Exec(query, product.Name, product.BreadUnits)
	if err != nil {
		return fmt.Errorf("%s: failed add product: %w", op, err)
	}

	return nil
}

func (db *repository) GetProductsWithValueInName(value string) ([]*products.Product, error) {
	const op = "storage.Product.GetProductsWithValueInName"

	query := "SELECT * FROM products WHERE product_name ILIKE '%' || $1 || '%'"
	rows, err := db.DB.Query(query, value)
	if err != nil {
		return nil, fmt.Errorf("%s: failed get products which contains value: %w", op, err)
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println(fmt.Errorf("%s: failed close product's rows: %w", op, err))
		}
	}(rows)

	var product []*products.Product

	for rows.Next() {
		p := &products.Product{}
		err = rows.Scan(&p.ProductID, &p.Name, &p.BreadUnits)
		if err != nil {
			return nil, fmt.Errorf("%s: failed scan product's rows: %w", op, err)
		}
		product = append(product, p)
	}

	return product, nil
}

func (db *repository) GetAllProducts() ([]*products.Product, error) {
	const op = "storage.Product.GetAllProducts"

	query := "SELECT * FROM products order by product_name"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("%s: failed get all products: %w", op, err)
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println(fmt.Errorf("%s: failed close product's rows: %w", op, err))
		}
	}(rows)

	var product []*products.Product

	for rows.Next() {
		p := &products.Product{}
		err = rows.Scan(&p.ProductID, &p.Name, &p.BreadUnits)
		if err != nil {
			return nil, fmt.Errorf("%s: failed scan product's rows: %w", op, err)
		}
		product = append(product, p)
	}

	return product, nil
}

func (db *repository) GetBreadUnitAmount(name string) (int, error) {
	const op = "storage.product.GetBreadUnitAmount"

	query := "SELECT carbs FROM products WHERE product_name = $1"
	row, err := db.DB.Query(query, name)
	if err != nil {
		return -1, fmt.Errorf("%s: failed get carbs amount: %w", op, err)
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println(fmt.Errorf("%s: failed close product's row: %w", op, err))
		}
	}(row)

	var carbs int
	row.Next()
	err = row.Scan(&carbs)
	if err != nil {
		return -1, fmt.Errorf("%s: failed scan product's row: %w", op, err)
	}

	return carbs, nil
}

package repo

import (
	"database/sql"
	"fmt"
	"product/entity"
	"strings"
	"time"
)

type ProductRepo interface {
	CreateProduct(product *entity.Product) (*entity.Product, error)
	ListProducts() ([]entity.Product, error)
	GetProductById(id int) (*entity.Product, error)
	GetProductsByCategoryId(id int) ([]entity.Product, error)
	GetProductsByCategoryName(name string) ([]entity.Product, error)
	UpdateProductById(id int, params map[string]any) error
}

type productRepoImpl struct {
	db *sql.DB
}

func NewProductRepo(db *sql.DB) productRepoImpl {
	return productRepoImpl{
		db: db,
	}
}

func (repo productRepoImpl) CreateProduct(product *entity.Product) (*entity.Product, error) {
	const FormatDateOnly = "2006-01-02"
	productDate, err := time.Parse(FormatDateOnly, product.ProductDate)
	if err != nil {
		return nil, err
	}

	sql := `INSERT INTO 
				products (name, stock, price, category_id, product_date, created_at, updated_at) 
			VALUES 
				($1, $2, $3, $4, $5, NOW(), NOW()) 
			RETURNING 
				id, created_at, updated_at`

	err = repo.db.QueryRow(
		sql,
		product.Name,
		product.Stock,
		product.Price,
		product.ProductCategoryId,
		productDate,
	).Scan(
		&product.Id,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (repo productRepoImpl) ListProducts() ([]entity.Product, error) {
	products := []entity.Product{}

	sql := `SELECT id, category_id, name, stock, price, product_date, created_at, updated_at FROM products WHERE deleted_at IS NULL`

	rows, err := repo.db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product entity.Product
		err := rows.Scan(
			&product.Id,
			&product.ProductCategoryId,
			&product.Name,
			&product.Stock,
			&product.Price,
			&product.ProductDate,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (repo productRepoImpl) GetProductById(id int) (*entity.Product, error) {
	var product entity.Product

	sql := `SELECT id, category_id, name, stock, price, product_date, created_at, updated_at FROM products WHERE id = $1;`

	err := repo.db.QueryRow(sql, id).Scan(
		&product.Id,
		&product.ProductCategoryId,
		&product.Name,
		&product.Stock,
		&product.Price,
		&product.ProductDate,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (repo productRepoImpl) GetProductsByCategoryId(id int) ([]entity.Product, error) {
	products := []entity.Product{}

	sql := `SELECT id, category_id, name, stock, price, product_date, created_at, updated_at FROM products WHERE category_id = $1 AND deleted_at IS NULL;`

	rows, err := repo.db.Query(sql, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product entity.Product
		err := rows.Scan(
			&product.Id,
			&product.ProductCategoryId,
			&product.Name,
			&product.Stock,
			&product.Price,
			&product.ProductDate,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (repo productRepoImpl) GetProductsByCategoryName(name string) ([]entity.Product, error) {
	products := []entity.Product{}

	sql := `SELECT 
				id, category_id, name, stock, price, product_date, created_at, updated_at 
			FROM 
				products p
			WHERE 
				category_id = (SELECT id FROM categories c WHERE name = $1) AND p.deleted_at IS NULL;`

	rows, err := repo.db.Query(sql, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product entity.Product
		err := rows.Scan(
			&product.Id,
			&product.ProductCategoryId,
			&product.Name,
			&product.Stock,
			&product.Price,
			&product.ProductDate,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (repo productRepoImpl) checkProductById(id int) error {
	var productId int
	sql := `SELECT id FROM products WHERE id = $1;`
	err := repo.db.QueryRow(sql, id).Scan(&productId)
	if err != nil {
		return err
	}
	return nil
}

func (repo productRepoImpl) UpdateProductById(id int, params map[string]any) error {
	err := repo.checkProductById(id)
	if err != nil {
		return err
	}

	var sql strings.Builder
	sql.WriteString("UPDATE products SET ")

	queryParams := []any{}
	counter := 0
	for field, val := range params {
		if field == "id" {
			continue
		}
		fmt.Fprintf(&sql, "%s = $%d, ", field, counter+1)
		queryParams = append(queryParams, val)
		counter++
	}

	fmt.Fprintf(&sql, "updated_at = NOW() WHERE id = $%d", counter+1)
	queryParams = append(queryParams, params["id"])

	_, err = repo.db.Exec(sql.String(), queryParams...)
	if err != nil {
		return err
	}

	return nil
}

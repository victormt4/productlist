package productrepo

import (
	"database/sql"
	"productlist/model"
	"productlist/utils/errorutils"
)

type Repo interface {
	GetAll() []model.Product
	Insert(product model.Product) model.Product
}

type ProductRepo struct {
	db *sql.DB
}

func GetRepo(db *sql.DB) Repo {
	return ProductRepo{db: db}
}

func (r ProductRepo) GetAll() []model.Product {

	db := r.db

	products := make([]model.Product, 0)

	rows, err := db.Query("SELECT * FROM produtos")

	errorutils.PanicOnError(err)

	for rows.Next() {

		var id int64
		var quantidade int
		var nome, descricao string
		var preco float64

		err := rows.Scan(&id, &nome, &descricao, &preco, &quantidade)

		errorutils.PanicOnError(err)

		products = append(products, model.Product{
			Id:          id,
			Name:        nome,
			Description: descricao,
			Price:       preco,
			Quantity:    quantidade,
		})
	}

	return products
}

func (r ProductRepo) Insert(product model.Product) model.Product {

	db := r.db

	statement, err := db.Prepare("INSERT INTO produtos (nome, descricao, preco, quantidade) VALUES ($1, $2, $3, $4) RETURNING id")
	errorutils.PanicOnError(err)
	defer statement.Close()

	var productId int64

	err = statement.QueryRow(product.Name, product.Description, product.Price, product.Quantity).Scan(&productId)
	errorutils.PanicOnError(err)

	product.Id = productId

	return product
}

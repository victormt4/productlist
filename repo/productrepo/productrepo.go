package productrepo

import (
	"log"
	"productlist/db"
	"productlist/model"
	"productlist/utils/errorutils"
)

func GetAll() []model.Product {

	conn := db.GetConnection()
	defer conn.Close()

	products := make([]model.Product, 0)

	rows, err := conn.Query("SELECT * FROM produtos")

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

func Insert(product model.Product) model.Product {

	conn := db.GetConnection()
	defer conn.Close()

	statement, err := conn.Prepare("INSERT INTO produtos (nome, descricao, preco, quantidade) VALUES ($1, $2, $3, $4) RETURNING id")
	if err != nil {
		log.Panic(err)
	}
	defer statement.Close()

	var productId int64

	err = statement.QueryRow(product.Name, product.Description, product.Price, product.Quantity).Scan(&productId)
	errorutils.PanicOnError(err)

	product.Id = productId

	return product
}

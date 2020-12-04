package productrepo

import (
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

		var id, quantidade int
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

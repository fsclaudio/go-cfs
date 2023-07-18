package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"https://github.com/fsclaudio/go-cfs/infra/database"
	"https://github.com/fsclaudio/go-cfs/internal/usecase"
)

func main() {
	insertOrders("12345", 120.00, 2.00)
}

func insertOrders(id string, price, tax float64) {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}
	defer db.Close() // espera tudo rodar e depois executa o close
	orderRepository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPrice(orderRepository)

	input := usecase.OrderInput{
		ID:    id,
		Price: price,
		Tax:   tax,
	}
	output, err := uc.Execute(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}

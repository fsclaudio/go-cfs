package database

import (
	"database/sql"

	"https://github.com/fsclaudio/go-cfs/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		Db: db,
	}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	// Create table
	statement, err := r.DB.Prepare("CREATE TABLE IF NOT EXISTS orders (id varchar(255) PRIMARY KEY NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL)")
	if err != nil {
		log.Println("Error in creating table")
	} else {
		log.Println("Successfully created table orders!")
	}
	statement.Exec()

	_, err := r.Db.Exec("INSERT INTO orders (id, price, tax, final_price) VALUES (?,?,?,?)",
		order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetTotalTransactions() (int, error) {
	var total int
	err := r.Db.QueryRow("SELECT COUNT(*) FROM orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

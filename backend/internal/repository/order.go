package repository

import (
	"fmt"

	"github.com/TSI-Projects/group-project/internal/db"
)

type Order struct {
	ID            int     `db:"id"              json:"id"                  validate:"omitempty"`
	OrderStatusID int     `db:"order_status_id" json:"order_status_id"     validate:"omitempty"`
	OrderTypeID   int     `db:"order_type_id"   json:"order_type_id"       validate:"required"`
	WorkerID      int     `db:"worker_id"       json:"worker_id"           validate:"required"`
	CustomerID    int     `db:"customer_id"     json:"customer_id"         validate:"omitempty"`
	Reason        string  `db:"reason"          json:"reason"              validate:"required"`
	Defect        string  `db:"defect"          json:"defect"              validate:"required"`
	TotalPrice    float64 `db:"total_price"     json:"total_price"         validate:"required"`
	Prepayment    float64 `db:"prepayment"      json:"prepayment"          validate:"required"`

	Status   *OrderStatus `db:"order_statuses"  json:"status"              validate:"omitempty"`
	Type     *OrderType   `db:"order_types"     json:"type"                validate:"omitempty"`
	Customer *Customer    `db:"customers"       json:"customer"            validate:"omitempty"`
	Worker   *Worker      `db:"workers"         json:"worker"              validate:"omitempty"`
}

type OrderRepo struct {
	DBClient db.IDatabase
}

func NewOrderRepo(database db.IDatabase) IRepository[Order] {
	return &OrderRepo{
		DBClient: database,
	}
}

func (o *OrderRepo) Create(order *Order) error {
	var orderStatusID int
	db := o.DBClient.GetConn()

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	if err = tx.QueryRow("INSERT INTO order_statuses DEFAULT VALUES RETURNING id").Scan(&orderStatusID); err != nil {
		return fmt.Errorf("failed to insert order status item into order_statuses table: %v", err)
	}

	if order.CustomerID <= 0 {
		if err := tx.QueryRow(
			`INSERT INTO customers
				(phone_number, language_id)
			VALUES
				($1, $2)
			RETURNING id`,
			order.Customer.PhoneNumber, order.Customer.LanguageID,
		).Scan(&order.CustomerID); err != nil {
			return fmt.Errorf("failed to insert customer: %v", err)
		}
	}

	if _, err = tx.Exec(
		`INSERT INTO orders 
				(reason, defect, total_price_eur, prepayment_eur, worker_id, customer_id, order_status_id, order_type_id)
			VALUES 
				($1, $2, $3, $4, $5, $6, $7, $8)`,
		order.Reason, order.Defect, order.TotalPrice, order.Prepayment, order.WorkerID, order.CustomerID, orderStatusID, order.OrderTypeID); err != nil {
		return fmt.Errorf("failed to insert order item into orders table: %v", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit create order changes: %v", err)
	}

	return nil
}

func (o *OrderRepo) Delete(id int) error {
	db := o.DBClient.GetConn()

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	if _, err = tx.Exec("DELETE FROM orders WHERE id = $1", id); err != nil {
		return fmt.Errorf("failed to execute delete command for orders table, with id '%d': %v", id, err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}

func (o *OrderRepo) GetByID(id int) (*Order, error) {
	return nil, nil
}

func (o *OrderRepo) GetAll() ([]*Order, error) {
	orders := make([]*Order, 0)

	selectQuery := `
	SELECT
		o.id AS order_id,
		o.order_status_id,
		o.order_type_id,
		o.worker_id,
		o.customer_id,
		o.reason,
		o.defect,
		o.total_price_eur,
		o.prepayment_eur,
		os.id AS order_status_id,
		os.ready_at,
		os.returned_at,
		os.customer_notified_at,
		os.is_outsourced,
		os.is_receipt_lost,
		ot.id AS order_types_id,
		ot.full_name,
		c.id AS customer_id,
		c.language_id,
		c.phone_number,
		w.id AS worker_id,
		w.first_name,
		w.last_name
	FROM
		orders o
	JOIN
		order_statuses os
		ON o.order_status_id = os.id
	JOIN
		order_types ot
		ON o.order_type_id = ot.id
	JOIN
		customers c
		ON o.customer_id = c.id
	JOIN
		workers w
		ON o.worker_id = w.id;`

	rows, err := o.DBClient.GetConn().Query(selectQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to query orders: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		order := &Order{
			Status:   &OrderStatus{},
			Worker:   &Worker{},
			Customer: &Customer{},
			Type:     &OrderType{},
		}

		if err := rows.Scan(
			&order.ID,
			&order.OrderStatusID,
			&order.OrderTypeID,
			&order.WorkerID,
			&order.CustomerID,
			&order.Reason,
			&order.Defect,
			&order.TotalPrice,
			&order.Prepayment,
			&order.Status.ID,
			&order.Status.ReadyAt,
			&order.Status.ReturnedAt,
			&order.Status.CustomerNotifiedAt,
			&order.Status.IsOutsourced,
			&order.Status.IsReceiptLost,
			&order.Type.ID,
			&order.Type.FullName,
			&order.Customer.ID,
			&order.Customer.LanguageID,
			&order.Customer.PhoneNumber,
			&order.Worker.ID,
			&order.Worker.FirstName,
			&order.Worker.LastName,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row in orders table: %v", err)
		}
		orders = append(orders, order)
	}
	fmt.Printf("Fetched Order: %+v\n", orders)
	return orders, nil
}

func (o *OrderRepo) Update(order *Order) error {
	return nil
}

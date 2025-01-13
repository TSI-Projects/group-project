package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/TSI-Projects/group-project/internal/db"
	"github.com/TSI-Projects/group-project/utils"
)

type Order struct {
	ID uint `db:"id"              json:"id"                    validate:"omitempty"`

	OrderStatusID uint `db:"order_status_id" json:"order_status_id,omitempty"       validate:"omitempty"`
	OrderTypeID   uint `db:"order_type_id"   json:"order_type_id,omitempty"         validate:"required"`
	WorkerID      uint `db:"worker_id"       json:"worker_id,omitempty"             validate:"required"`
	CustomerID    uint `db:"customer_id"     json:"customer_id,omitempty"           validate:"omitempty"`

	Reason     string     `db:"reason"          json:"reason"                validate:"required"`
	Defect     string     `db:"defect"          json:"defect"                validate:"required"`
	ItemName   string     `db:"item_name"       json:"item_name"             validate:"required"`
	TotalPrice float64    `db:"total_price"     json:"total_price"           validate:"required"`
	Prepayment float64    `db:"prepayment"      json:"prepayment"            validate:"required"`
	CreatedAt  *time.Time `db:"created_at"      json:"created_at,omitempty"  validate:"omitempty"`

	Status   *OrderStatus `db:"order_statuses"  json:"status"              validate:"omitempty"`
	Type     *OrderType   `db:"order_types"     json:"type"                validate:"omitempty"`
	Customer *Customer    `db:"customers"       json:"customer"            validate:"omitempty"`
	Worker   *Worker      `db:"workers"         json:"worker"              validate:"omitempty"`
}

type OrderRepo struct {
	DBClient db.IDatabase
}

func NewOrderRepo(database db.IDatabase) IOrderRepository[Order] {
	return &OrderRepo{
		DBClient: database,
	}
}

func (o *OrderRepo) Create(order *Order) error {
	var orderStatusID int

	tx, err := o.DBClient.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	if err := tx.QueryRow(`
		SELECT id 
		FROM customers
		WHERE phone_number = $1`,
		order.Customer.PhoneNumber,
	).Scan(&order.CustomerID); err != nil {
		if err != sql.ErrNoRows {
			return fmt.Errorf("failed to query customer id: %v", err)
		}

		if err := tx.QueryRow(`
			INSERT INTO customers (phone_number, language_id)
			VALUES ($1, $2)
			RETURNING id`,
			order.Customer.PhoneNumber,
			order.Customer.LanguageID,
		).Scan(&order.CustomerID); err != nil {
			return fmt.Errorf("failed to create new customer: %v", err)
		}
	}

	if err = tx.QueryRow("INSERT INTO order_statuses DEFAULT VALUES RETURNING id").Scan(&orderStatusID); err != nil {
		return fmt.Errorf("failed to insert order status item into order_statuses table: %v", err)
	}

	if _, err = tx.Exec(
		`INSERT INTO orders 
				(reason,
				 defect,
				 item_name,
				 total_price_eur,
				 prepayment_eur,
				 created_at,
				 worker_id,
				 customer_id,
				 order_status_id,
				 order_type_id)
			VALUES 
				($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		order.Reason,
		order.Defect,
		order.ItemName,
		order.TotalPrice,
		order.Prepayment,
		utils.GetTimestamp(),
		order.WorkerID,
		order.CustomerID,
		orderStatusID,
		order.OrderTypeID,
	); err != nil {
		return fmt.Errorf("failed to insert order item into orders table: %v", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit create order changes: %v", err)
	}

	return nil
}

func (o *OrderRepo) Delete(id uint) error {
	if _, err := o.DBClient.Exec("DELETE FROM orders WHERE id = $1", id); err != nil {
		return fmt.Errorf("failed to execute delete command for orders table, with id '%d': %v", id, err)
	}
	return nil
}

func (o *OrderRepo) GetAll() ([]*Order, error) {
	selectQuery := `
    SELECT
        o.id AS order_id,
        o.item_name,
        o.reason,
        o.defect,
        o.total_price_eur,
        o.prepayment_eur,
        o.created_at,
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
        order_statuses os ON o.order_status_id = os.id
    JOIN
        order_types ot ON o.order_type_id = ot.id
    JOIN
        customers c ON o.customer_id = c.id
    JOIN
        workers w ON o.worker_id = w.id;`

	rows, err := o.DBClient.Query(selectQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to query orders: %v", err)
	}
	defer rows.Close()

	orders, err := scanOrders(rows)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (o *OrderRepo) GetByID(id uint) (*Order, error) {
	order := &Order{
		ID:       id,
		Status:   &OrderStatus{},
		Worker:   &Worker{},
		Customer: &Customer{},
		Type:     &OrderType{},
	}

	if err := o.DBClient.QueryRow(
		`SELECT
			o.id AS order_id,
			o.item_name,
			o.reason,
			o.defect,
			o.total_price_eur,
			o.prepayment_eur,
			o.created_at,
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
			ON o.worker_id = w.id
		WHERE o.id = $1;`, id,
	).Scan(
		&order.ID,
		&order.ItemName,
		&order.Reason,
		&order.Defect,
		&order.TotalPrice,
		&order.Prepayment,
		&order.CreatedAt,
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
		return order, fmt.Errorf("failed to make query row request: %w", err)
	}

	return order, nil
}

func (o *OrderRepo) Update(order *Order) error {
	if _, err := o.DBClient.Exec(
		`UPDATE orders 
		SET order_type_id   = $1,
			worker_id       = $2,
			item_name       = $3,
			reason          = $4,
			defect          = $5,
			total_price_eur = $6,
			prepayment_eur  = $7
		WHERE id = $8;`,
		order.OrderTypeID,
		order.WorkerID,
		order.ItemName,
		order.Reason,
		order.Defect,
		order.TotalPrice,
		order.Prepayment,
		order.ID,
	); err != nil {
		return fmt.Errorf("failed to make exec request: %w", err)
	}
	return nil
}

func (o *OrderRepo) GetActiveOrders() ([]*Order, error) {
	selectQuery := `
    SELECT
        o.id AS order_id,
        o.item_name,
        o.reason,
        o.defect,
        o.total_price_eur,
        o.prepayment_eur,
        o.created_at,
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
        order_statuses os ON o.order_status_id = os.id
    JOIN
        order_types ot ON o.order_type_id = ot.id
    JOIN
        customers c ON o.customer_id = c.id
    JOIN
        workers w ON o.worker_id = w.id
    WHERE
        os.ready_at IS NULL;`

	rows, err := o.DBClient.Query(selectQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to query active orders: %v", err)
	}
	defer rows.Close()

	orders, err := scanOrders(rows)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (o *OrderRepo) GetCompletedOrders() ([]*Order, error) {
	selectQuery := `
    SELECT
        o.id AS order_id,
        o.item_name,
        o.reason,
        o.defect,
        o.total_price_eur,
        o.prepayment_eur,
        o.created_at,
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
        order_statuses os ON o.order_status_id = os.id
    JOIN
        order_types ot ON o.order_type_id = ot.id
    JOIN
        customers c ON o.customer_id = c.id
    JOIN
        workers w ON o.worker_id = w.id
    WHERE
        os.ready_at IS NOT NULL;`
	rows, err := o.DBClient.Query(selectQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to query completed orders: %v", err)
	}
	defer rows.Close()

	orders, err := scanOrders(rows)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func scanOrders(rows *sql.Rows) ([]*Order, error) {
	orders := make([]*Order, 0)
	for rows.Next() {
		order := &Order{
			Status:   &OrderStatus{},
			Worker:   &Worker{},
			Customer: &Customer{},
			Type:     &OrderType{},
		}
		if err := rows.Scan(
			&order.ID,
			&order.ItemName,
			&order.Reason,
			&order.Defect,
			&order.TotalPrice,
			&order.Prepayment,
			&order.CreatedAt,
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
	return orders, nil
}

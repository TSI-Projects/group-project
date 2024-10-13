package repository

import (
	"fmt"
	"time"

	"github.com/TSI-Projects/group-project/internal/db"
)

type OrderStatus struct {
	ID                 int        `db:"id"                       json:"id"                     validate:"omitempty"`
	ReadyAt            *time.Time `db:"ready_at"                 json:"ready_at"               validate:"omitempty"`
	ReturnedAt         *time.Time `db:"returned_at"              json:"returned_at"            validate:"omitempty"`
	CustomerNotifiedAt *time.Time `db:"customer_notified_at"     json:"customer_notified_at"   validate:"omitempty"`
	IsOutsourced       bool       `db:"is_outsourced"            json:"is_outsourced"          validate:"omitempty"`
	IsReceiptLost      bool       `db:"is_recipient_lost"        json:"is_recipient_lost"      validate:"omitempty"`
}

type OrderStatusRepo struct {
	DBClient *db.DBClient
}

func NewOrderStatusRepo(dbClient db.DBClient) IRepository[OrderStatus] {
	return &OrderStatusRepo{
		DBClient: &dbClient,
	}
}

func (o *OrderStatusRepo) Create(orderStatus *OrderStatus) error {
	if _, err := o.DBClient.Exec(
		`INSERT INTO order_statuses
			(ready_at, returned_at, customer_notified_at, is_outsourced, is_receipt_lost)
		 VALUES 
		 	($1, $2, $3, $4, $5)
		`,
		orderStatus.ReadyAt,
		orderStatus.ReturnedAt,
		orderStatus.CustomerNotifiedAt,
		orderStatus.IsOutsourced,
		orderStatus.IsReceiptLost,
	); err != nil {
		return fmt.Errorf("failed to make exec request: %v", err)
	}
	return nil
}

func (o *OrderStatusRepo) Delete(id int) error {
	if _, err := o.DBClient.Exec(
		`DELETE FROM order_statuses WHERE id = $1`, id); err != nil {
		return fmt.Errorf("failed to make exec request: %v", err)
	}
	return nil
}

func (o *OrderStatusRepo) GetAll() ([]*OrderStatus, error) {
	orderStatuses := make([]*OrderStatus, 0)

	rows, err := o.DBClient.Query(
		`SELECT
			id
		 	ready_at,
		 	returned_at,
		 	customer_notified_at,
		 	is_outsourced,
			is_receipt_lost
		FROM
			order_statuses`,
	)
	if err != nil {
		return orderStatuses, fmt.Errorf("failed to make query request: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		orderStatus := &OrderStatus{}

		if err := rows.Scan(
			&orderStatus.ID,
			&orderStatus.ReadyAt,
			&orderStatus.ReturnedAt,
			&orderStatus.CustomerNotifiedAt,
			&orderStatus.IsOutsourced,
			&orderStatus.IsReceiptLost,
		); err != nil {
			return orderStatuses, fmt.Errorf("failed to scan rows: %v", err)
		}

		orderStatuses = append(orderStatuses, orderStatus)
	}

	return orderStatuses, nil
}

func (o *OrderStatusRepo) GetByID(id int) (*OrderStatus, error) {
	orderStatus := &OrderStatus{ID: id}

	if err := o.DBClient.QueryRow(
		`SELECT
		 	ready_at,
		 	returned_at,
		 	customer_notified_at,
		 	is_outsourced,
			is_receipt_lost
		FROM order_statuses
		WHERE id = $1`, id,
	).Scan(
		&orderStatus.ReadyAt,
		&orderStatus.ReturnedAt,
		&orderStatus.CustomerNotifiedAt,
		&orderStatus.IsOutsourced,
		&orderStatus.IsReceiptLost,
	); err != nil {
		return orderStatus, fmt.Errorf("failed to make query row request: %v", err)
	}

	return orderStatus, nil
}

func (o *OrderStatusRepo) Update(orderStatus *OrderStatus) error {
	if _, err := o.DBClient.Exec(
		`UPDATE order_statuses
		 SET
			ready_at = $1,
		 	returned_at = $2,
		 	customer_notified_at = $3,
		 	is_outsourced = $4,
			is_receipt_lost = $5
		 WHERE id = $6`,
		orderStatus.ReadyAt,
		orderStatus.ReturnedAt,
		orderStatus.CustomerNotifiedAt,
		orderStatus.IsOutsourced,
		orderStatus.IsReceiptLost,
		orderStatus.ID,
	); err != nil {
		return fmt.Errorf("failed to make exec request: %v", err)
	}

	return nil
}

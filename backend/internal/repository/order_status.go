package repository

import "time"

type OrderStatus struct {
	ID                 int        `db:"id"`
	ReadyAt            *time.Time `db:"ready_at"`
	ReturnedAt         *time.Time `db:"returned_at"`
	CustomerNotifiedAt *time.Time `db:"customer_notified_at"`
	IsOutsourced       bool       `db:"is_outsourced"`
	IsReceiptLost      bool       `db:"is_recipient"`
}

package repository

import "time"

type OrderStatus struct {
	ID                 int        `db:"id"                       json:"id"`
	ReadyAt            *time.Time `db:"ready_at"                 json:"ready_at"`
	ReturnedAt         *time.Time `db:"returned_at"              json:"returned_at"`
	CustomerNotifiedAt *time.Time `db:"customer_notified_at"     json:"customer_notified_at"`
	IsOutsourced       bool       `db:"is_outsourced"            json:"is_outsourced"`
	IsReceiptLost      bool       `db:"is_recipient_lost"        json:"is_recipient_lost"`
}

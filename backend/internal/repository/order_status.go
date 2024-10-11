package repository

import "time"

type OrderStatus struct {
	ID                 int        `db:"id"                       json:"id"                     validate:"omitempty"`
	ReadyAt            *time.Time `db:"ready_at"                 json:"ready_at"               validate:"omitempty"`
	ReturnedAt         *time.Time `db:"returned_at"              json:"returned_at"            validate:"omitempty"`
	CustomerNotifiedAt *time.Time `db:"customer_notified_at"     json:"customer_notified_at"   validate:"omitempty"`
	IsOutsourced       bool       `db:"is_outsourced"            json:"is_outsourced"          validate:"omitempty"`
	IsReceiptLost      bool       `db:"is_recipient_lost"        json:"is_recipient_lost"      validate:"omitempty"`
}

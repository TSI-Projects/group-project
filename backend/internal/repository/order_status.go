package repository

type OrderStatus struct {
	ID                 int    `db:"id"`
	ReadyAt            string `db:"ready_at"`
	ReturnedAt         string `db:"returned_at"`
	CustomerNotifiedAt string `db:"customer_notified_at"`
	IsOutsourced       bool   `db:"is_outsourced"`
	IsReceiptLost      bool   `db:"is_recipient"`
}

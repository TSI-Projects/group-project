package repository

type Order struct {
	ID            int     `db:"id"`
	OrderStatusID int     `db:"order_status_id"`
	OrderTypeID   int     `db:"order_type_id"`
	WorkerID      int     `db:"worker_id"`
	Reason        string  `db:"reason"`
	Defect        string  `db:"defect"`
	TotalPrice    float64 `db:"total_price"`
	Prepayment    float64 `db:"prepayment"`
}

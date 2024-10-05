package repository

type OrderType struct {
	ID       int    `db:"id"`
	FullName string `db:"full_name"`
}

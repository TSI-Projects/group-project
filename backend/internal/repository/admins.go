package repository

type Admin struct {
	ID       int    `db:"id"          json:"id"`
	Username string `db:"username"    json:"username"`
	Password string `db:"password"    json:"password"`
}

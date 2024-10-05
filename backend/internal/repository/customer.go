package repository

type Customer struct {
	ID          int    `db:"id"`
	LanguageID  int    `db:"language_id"`
	PhoneNumber string `db:"phone_number"`
}

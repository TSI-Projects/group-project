package repository

type Customer struct {
	ID          int    `db:"id"              json:"id"`
	LanguageID  int    `db:"language_id"     json:"language_id"`
	PhoneNumber string `db:"phone_number"    json:"phone_number"`
}

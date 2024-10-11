package repository

type Customer struct {
	ID          int    `db:"id"              json:"id"               validate:"omitempty"`
	LanguageID  int    `db:"language_id"     json:"language_id"      validate:"required"`
	PhoneNumber string `db:"phone_number"    json:"phone_number"     validate:"required"`
}

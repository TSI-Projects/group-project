package repository

type Language struct {
	ID        int    `db:"id"`
	ShortName string `db:"short_name"`
	FullName  string `db:"full_name"`
}

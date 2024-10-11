package repository

import (
	"fmt"

	"github.com/TSI-Projects/group-project/internal/db"
)

type Language struct {
	ID        int    `db:"id"          json:"id,omitempty"  validate:"omitempty"`
	ShortName string `db:"short_name"  json:"short_name"    validate:"required"`
	FullName  string `db:"full_name"   json:"full_name"     validate:"required"`
}

type LanguageRepo struct {
	DBClient db.IDatabase
}

func NewLanguageRepo(database db.IDatabase) IRepository[Language] {
	return &LanguageRepo{
		DBClient: database,
	}
}

func (l *LanguageRepo) Create(language *Language) error {
	db := l.DBClient.GetConn()

	if _, err := db.Query(
		`INSERT INTO languages
			(short_name, full_name)
		VALUES
			($1, $2)`,
		language.ShortName,
		language.FullName,
	); err != nil {
		return fmt.Errorf("failed to make insert query request: %v", err)
	}

	return nil
}

func (l *LanguageRepo) Delete(id int) error {
	db := l.DBClient.GetConn()

	if _, err := db.Query("DELETE FROM languages WHERE id = $1", id); err != nil {
		return fmt.Errorf("failed to make delete query request: %v", err)
	}

	return nil
}

func (l *LanguageRepo) GetAll() ([]*Language, error) {
	db := l.DBClient.GetConn()
	var languages []*Language

	rows, err := db.Query(
		`SELECT
			id,
			short_name,
			full_name
		FROM
			languages
		`,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to make select query request: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		language := &Language{}

		if err := rows.Scan(
			&language.ID,
			&language.ShortName,
			&language.FullName,
		); err != nil {
			return nil, fmt.Errorf("failed to scan rows in languages table: %v", err)
		}
		languages = append(languages, language)
	}

	return languages, nil
}

func (l *LanguageRepo) GetByID(int) (*Language, error) {
	return nil, nil
}

func (l *LanguageRepo) Update(*Language) error {
	return nil
}

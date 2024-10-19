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
	if _, err := l.DBClient.Exec(
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
	if _, err := l.DBClient.Exec("DELETE FROM languages WHERE id = $1", id); err != nil {
		return fmt.Errorf("failed to make delete query request: %v", err)
	}
	return nil
}

func (l *LanguageRepo) GetAll() ([]*Language, error) {
	var languages []*Language

	rows, err := l.DBClient.Query(
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

func (l *LanguageRepo) GetByID(id int) (*Language, error) {
	language := &Language{ID: id}

	if err := l.DBClient.QueryRow(
		`SELECT full_name, short_name
		 FROM languages
		 WHERE id = $1`, id,
	).Scan(
		&language.FullName,
		&language.ShortName,
	); err != nil {
		return language, fmt.Errorf("failed to make query row request: %w", err)
	}

	return language, nil
}

func (l *LanguageRepo) Update(language *Language) error {
	if _, err := l.DBClient.Exec(
		`UPDATE languages
         SET full_name = $1, short_name = $2
         WHERE id = $3`,
		language.FullName,
		language.ShortName,
		language.ID,
	); err != nil {
		return fmt.Errorf("failed to make exec request: %w", err)
	}

	return nil
}

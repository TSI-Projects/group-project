package repository

import (
	"fmt"

	"github.com/TSI-Projects/group-project/internal/db"
)

type Language struct {
	ID        int    `db:"id"          json:"id"`
	ShortName string `db:"short_name"  json:"short_name"`
	FullName  string `db:"full_name"   json:"full_name"`
}

type LanguageRepo struct {
	DBClient db.IDatabase
}

type ILanguageRepo interface {
	CreateLanguage(*Language) error
	GetLanguage() ([]*Language, error)
	DeleteLanguage(id int) error
}

func NewLanguageRepo(database db.IDatabase) ILanguageRepo {
	return &LanguageRepo{
		DBClient: database,
	}
}

func (l *LanguageRepo) CreateLanguage(language *Language) error {
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

func (l *LanguageRepo) DeleteLanguage(id int) error {
	db := l.DBClient.GetConn()

	if _, err := db.Query("DELETE FROM languages WHERE id = $1", id); err != nil {
		return fmt.Errorf("failed to make delete query request: %v", err)
	}

	return nil
}

func (l *LanguageRepo) GetLanguage() ([]*Language, error) {
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

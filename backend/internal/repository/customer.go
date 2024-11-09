package repository

import (
	"fmt"
	"log"

	"github.com/TSI-Projects/group-project/internal/db"
)

type Customer struct {
	ID          uint   `db:"id"              json:"id"               validate:"omitempty"`
	LanguageID  int    `db:"language_id"     json:"language_id"      validate:"required"`
	PhoneNumber string `db:"phone_number"    json:"phone_number"     validate:"required"`

	Language *Language `db:"language"        json:"language"`
}

type CustomerRepo struct {
	DBClient db.IDatabase
}

func NewCustomerRepo(dbClient db.IDatabase) IRepository[Customer] {
	return &CustomerRepo{
		DBClient: dbClient,
	}
}

func (c *CustomerRepo) Create(customer *Customer) error {
	if _, err := c.DBClient.Exec(
		`INSERT INTO customers
		 	(language_id, phone_number)
		 VALUES ($1, $2)`,
		customer.LanguageID,
		customer.PhoneNumber,
	); err != nil {
		return fmt.Errorf("failed to make exec request: %w", err)
	}
	return nil
}

func (c *CustomerRepo) Delete(id uint) error {
	if _, err := c.DBClient.Exec("DELETE FROM customers WHERE id = $1", id); err != nil {
		return fmt.Errorf("failed to make exec request: %w", err)
	}
	return nil
}

func (c *CustomerRepo) GetAll() ([]*Customer, error) {
	customers := make([]*Customer, 0)

	rows, err := c.DBClient.Query(
		`SELECT 
			c.id, 
			c.phone_number,
			l.id,
			l.full_name,
			l.short_name
		 FROM customers c
		 JOIN 
		 	languages l
			ON c.language_id = l.id
		`,
	)
	if err != nil {
		return customers, fmt.Errorf("failed to make query request: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		customer := &Customer{
			Language: &Language{},
		}

		if err := rows.Scan(
			&customer.ID,
			&customer.PhoneNumber,
			&customer.Language.ID,
			&customer.Language.FullName,
			&customer.Language.ShortName,
		); err != nil {
			return nil, fmt.Errorf("failed to scan rows: %v", err)
		}

		customers = append(customers, customer)
	}
	log.Println(customers)
	return customers, nil
}

func (c *CustomerRepo) GetByID(id uint) (*Customer, error) {
	customer := &Customer{
		ID:       id,
		Language: &Language{},
	}

	if err := c.DBClient.QueryRow(
		`SELECT 
			c.phone_number,
			l.id,
			l.full_name,
			l.short_name
		 FROM customers c
		 JOIN languages l
			ON c.language_id = l.id
		WHERE c.id = $1;
		`, id,
	).Scan(
		&customer.PhoneNumber,
		&customer.Language.ID,
		&customer.Language.FullName,
		&customer.Language.ShortName,
	); err != nil {
		return customer, fmt.Errorf("failed to scan row: %v", err)
	}

	return customer, nil
}

func (c *CustomerRepo) Update(customer *Customer) error {
	if _, err := c.DBClient.Exec(
		`UPDATE customers
		 SET phone_number = $1, language_id = $2
		 WHERE id = $3`,
		customer.PhoneNumber,
		customer.LanguageID,
		customer.ID,
	); err != nil {
		return fmt.Errorf("failed to make exec request: %v", err)
	}

	return nil
}

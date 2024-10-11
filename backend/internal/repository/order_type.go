package repository

import (
	"fmt"

	"github.com/TSI-Projects/group-project/internal/db"
)

type OrderType struct {
	ID       int    `db:"id"        json:"id"           validate:"omitempty"`
	FullName string `db:"full_name" json:"full_name"    validate:"required"`
}

type OrderTypeRepo struct {
	DBClient db.IDatabase
}

func NewOrderTypeRepo(dbClient db.IDatabase) IRepository[OrderType] {
	return &OrderTypeRepo{
		DBClient: dbClient,
	}
}

func (o *OrderTypeRepo) Create(orderType *OrderType) error {
	queryCommand := `
	INSERT INTO order_types
		(full_name)
	VALUES
		($1)
	`

	if _, err := o.DBClient.Exec(queryCommand, orderType.FullName); err != nil {
		return fmt.Errorf("failed to create order type: %v", err)
	}

	return nil
}

func (o *OrderTypeRepo) Delete(id int) error {
	if _, err := o.DBClient.Exec("DELETE FROM order_types WHERE id = $1", id); err != nil {
		return fmt.Errorf("failed to make query: %v", err)
	}
	return nil
}

func (o *OrderTypeRepo) GetAll() ([]*OrderType, error) {
	orderTypes := make([]*OrderType, 0)

	queryCommand := `
	SELECT
		id,
		full_name
	FROM
		order_types
	`

	rows, err := o.DBClient.Query(queryCommand)
	if err != nil {
		return nil, fmt.Errorf("failed to make a query command: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		orderType := &OrderType{}

		if err := rows.Scan(
			&orderType.ID,
			&orderType.FullName,
		); err != nil {
			return nil, fmt.Errorf("failed to scan rows in order_types table: %v", err)
		}
		orderTypes = append(orderTypes, orderType)
	}

	return orderTypes, nil
}

func (l *OrderTypeRepo) GetByID(int) (*OrderType, error) {
	panic("unimplemented")
}

func (l *OrderTypeRepo) Update(*OrderType) error {
	panic("unimplemented")
}

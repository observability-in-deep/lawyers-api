package customer

import (
	"context"
	"time"

	models "github.com/observability-in-deep/lawyers-api/src/model"
	"github.com/observability-in-deep/lawyers-api/src/pkg/pool"
)

func Get(ctx context.Context, customerCPF string) (*models.Customers, error) {

	Customers := &models.Customers{}

	now := time.Now().UTC()

	Customers.CreateAt = &now
	Customers.UpdateAt = &now

	query := `SELECT name, email, folder, cpf, phone, updated_at, created_at, lawyer_id FROM customers WHERE cpf = $1`

	conn, err := pool.GetConnection()
	if err != nil {
		return nil, err
	}

	err = conn.QueryRow(ctx, query, customerCPF).Scan(&Customers.Name, &Customers.Email, &Customers.Folder, &Customers.Cpf, &Customers.Phone, &Customers.UpdateAt, &Customers.CreateAt, &Customers.LawyerName)

	if err != nil {
		return nil, err
	}

	return Customers, nil

}

func Create(ctx context.Context, costumer *models.Customers) (*models.Customers, error) {

	now := time.Now().UTC()

	costumer.CreateAt = &now
	costumer.UpdateAt = &now

	query := `
		INSERT INTO customers (
			name, email, folder, cpf, phone, updated_at, created_at, lawyer_id, id
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, (SELECT id FROM lawyers WHERE name=$8), gen_random_uuid()
		) RETURNING id
	`

	conn, err := pool.GetConnection()
	if err != nil {
		return nil, err
	}

	var id string
	err = conn.QueryRow(ctx, query, costumer.Name, costumer.Email, costumer.Folder, costumer.Cpf, costumer.Phone, costumer.UpdateAt, costumer.CreateAt, costumer.LawyerName).Scan(&id)
	if err != nil {
		return nil, err
	}

	return costumer, nil

}

func Update(ctx context.Context, customerId string, costumer *models.Customers) (*models.Customers, error) {
	now := time.Now().UTC()

	costumer.UpdateAt = &now

	query := `UPDATE customers SET name = $1, email = $2, folder = $3, cpf = $4, phone = $5, updated_at = $6 WHERE id = $7 RETURNING id`

	conn, err := pool.GetConnection()
	if err != nil {
		return nil, err
	}

	err = conn.QueryRow(ctx, query, costumer.Name, costumer.Email, costumer.Folder, costumer.Cpf, costumer.Phone, costumer.UpdateAt, customerId).Scan(&costumer.ID)

	if err != nil {
		return nil, err
	}

	return costumer, nil

}

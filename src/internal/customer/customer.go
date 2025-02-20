package customer

import (
	"context"
	"time"

	models "github.com/observability-in-deep/lawyers-api/src/model"
	"github.com/observability-in-deep/lawyers-api/src/pkg/pool"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

func Get(ctx context.Context, customerCPF string) (*models.Customers, error) {

	tracer := otel.Tracer("customer")
	ctx, span := tracer.Start(ctx, "Get")
	defer span.End()

	Customers := &models.Customers{}

	now := time.Now().UTC()
	Customers.CreateAt = &now
	Customers.UpdateAt = &now

	query := `SELECT name, email, folder, cpf, phone, updated_at, created_at, lawyer_id FROM customers WHERE cpf = $1`

	conn, err := pool.GetConnection()
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	err = conn.QueryRow(ctx, query, customerCPF).Scan(&Customers.Name, &Customers.Email, &Customers.Folder, &Customers.Cpf, &Customers.Phone, &Customers.UpdateAt, &Customers.CreateAt, &Customers.LawyerName)

	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	span.SetAttributes(
		attribute.String("customerCPF", customerCPF),
		attribute.String("customerName", Customers.Name),
		attribute.String("customerEmail", Customers.Email),
		attribute.String("customerFolder", Customers.Folder),
		attribute.String("db.query", query),
	)

	return Customers, nil

}

func Create(ctx context.Context, customer *models.Customers) (*models.Customers, error) {

	tracer := otel.Tracer("customer")
	ctx, span := tracer.Start(ctx, "Create")
	defer span.End()

	now := time.Now().UTC()

	customer.CreateAt = &now
	customer.UpdateAt = &now

	query := `
		INSERT INTO customers (
			name, email, folder, cpf, phone, updated_at, created_at, lawyer_id, id
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, (SELECT id FROM lawyers WHERE name=$8), gen_random_uuid()
		) RETURNING id
	`

	conn, err := pool.GetConnection()
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	var id string
	err = conn.QueryRow(ctx, query, customer.Name, customer.Email, customer.Folder, customer.Cpf, customer.Phone, customer.UpdateAt, customer.CreateAt, customer.LawyerName).Scan(&id)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	span.SetAttributes(
		attribute.String("customerCPF", customer.Cpf),
		attribute.String("customerName", customer.Name),
		attribute.String("customerEmail", customer.Email),
		attribute.String("customerFolder", customer.Folder),
		attribute.String("db.query", query),
	)

	return customer, nil

}

func Update(ctx context.Context, customerId string, customer *models.Customers) (*models.Customers, error) {
	now := time.Now().UTC()

	customer.UpdateAt = &now

	query := `UPDATE customers SET name = $1, email = $2, folder = $3, cpf = $4, phone = $5, updated_at = $6 WHERE id = $7 RETURNING id`

	conn, err := pool.GetConnection()
	if err != nil {
		return nil, err
	}

	err = conn.QueryRow(ctx, query, customer.Name, customer.Email, customer.Folder, customer.Cpf, customer.Phone, customer.UpdateAt, customerId).Scan(&customer.ID)

	if err != nil {
		return nil, err
	}

	return customer, nil

}

func Delete(ctx context.Context, customerId string) error {

	query := `DELETE FROM customers WHERE id = $1`

	conn, err := pool.GetConnection()
	if err != nil {
		return err
	}

	_, err = conn.Exec(ctx, query, customerId)

	if err != nil {
		return err
	}

	return nil

}

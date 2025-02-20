package lawyers

import (
	"context"
	"time"

	models "github.com/observability-in-deep/lawyers-api/src/model"
	"github.com/observability-in-deep/lawyers-api/src/pkg/pool"
)

func Get(ctx context.Context, oab string) (*models.Lawyer, error) {

	lawyer := &models.Lawyer{}

	now := time.Now().UTC()

	lawyer.CreateAt = &now
	lawyer.UpdateAt = &now

	query := `SELECT name, email, oab, phone, updated_at, created_at FROM lawyers WHERE oab = $1`

	conn, err := pool.GetConnection()
	if err != nil {
		return nil, err
	}

	err = conn.QueryRow(ctx, query, oab).Scan(&lawyer.Name, &lawyer.Email, &lawyer.OAB, &lawyer.Phone, &lawyer.UpdateAt, &lawyer.CreateAt)

	if err != nil {
		return nil, err
	}

	return lawyer, nil
}

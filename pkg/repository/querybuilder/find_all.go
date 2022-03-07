package querybuilder

import (
	"context"
	"database/sql"
	"github.com/huandu/go-sqlbuilder"
	"simple-customer-service/db"
)

func FindAll(ctx context.Context, dbtx db.DBTX, sb *sqlbuilder.SelectBuilder) (rows *sql.Rows, err error) {
	query, args := sb.Build()
	rows, err = db.Query(ctx, dbtx, query, args)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

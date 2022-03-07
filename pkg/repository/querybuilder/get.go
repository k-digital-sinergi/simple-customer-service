package querybuilder

import (
	"context"
	"database/sql"
	"github.com/huandu/go-sqlbuilder"
	"simple-customer-service/db"
)

func Get(ctx context.Context, dbtx db.DBTX, sb *sqlbuilder.SelectBuilder) (row *sql.Row, err error) {
	query, args := sb.Build()
	row, err = db.QueryRow(ctx, dbtx, query, args)
	if err != nil {
		return nil, err
	}

	return row, nil
}

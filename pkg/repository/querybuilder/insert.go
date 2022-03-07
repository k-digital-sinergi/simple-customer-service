package querybuilder

import (
	"context"
	"github.com/huandu/go-sqlbuilder"
	"simple-customer-service/db"
)

func Insert(ctx context.Context, dbtx db.DBTX, sb *sqlbuilder.InsertBuilder) (insertID, rowsAffected int64, err error) {
	query, args := sb.Build()
	insertID, rowsAffected, err = db.Exec(ctx, dbtx, query, args)
	if err != nil {
		return
	}

	return
}

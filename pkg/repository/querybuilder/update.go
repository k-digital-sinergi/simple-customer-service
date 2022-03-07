package querybuilder

import (
	"context"
	"github.com/huandu/go-sqlbuilder"
	"simple-customer-service/db"
)

func Update(ctx context.Context, dbtx db.DBTX, sb *sqlbuilder.UpdateBuilder) (insertID, rowsAffected int64, err error) {
	query, args := sb.Build()
	insertID, rowsAffected, err = db.Exec(ctx, dbtx, query, args)
	if err != nil {
		return
	}

	return
}

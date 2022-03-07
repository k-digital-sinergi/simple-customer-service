package repo_customer

import (
	"context"
	"github.com/huandu/go-sqlbuilder"
	"simple-customer-service/db"
	"simple-customer-service/pkg/model"
	"simple-customer-service/pkg/repository/querybuilder"
	"strconv"
)

func Delete(ctx context.Context, dbtx db.DBTX, custID int64) (insertID, rowsAffected int64, err error) {
	sb := sqlbuilder.NewDeleteBuilder()
	sb.DeleteFrom(model.TblCustomer)
	sb.Where(
		sb.Equal("id", strconv.FormatInt(custID, 10)),
	)

	insertID, rowsAffected, err = querybuilder.Delete(ctx, dbtx, sb)
	if err != nil {
		return
	}

	return
}

package repo_customer

import (
	"context"
	"github.com/huandu/go-sqlbuilder"
	"simple-customer-service/db"
	"simple-customer-service/pkg/model"
	"simple-customer-service/pkg/repository/querybuilder"
)

func Create(ctx context.Context, dbtx db.DBTX, customer model.Customer) (insertID, rowsAffected int64, err error) {
	sb := sqlbuilder.NewInsertBuilder()
	sb.InsertInto(model.TblCustomer)
	sb.Cols("phone", "name", "balance")
	sb.Values(customer.Phone, customer.Name, customer.Balance)

	insertID, rowsAffected, err = querybuilder.Insert(ctx, dbtx, sb)
	if err != nil {
		return
	}

	return
}

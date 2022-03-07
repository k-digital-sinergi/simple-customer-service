package repo_customer

import (
	"context"
	"github.com/huandu/go-sqlbuilder"
	"simple-customer-service/db"
	"simple-customer-service/pkg/model"
	"simple-customer-service/pkg/repository/querybuilder"
)

func List(ctx context.Context, dbtx db.DBTX) (data []model.Customer, err error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.From(model.TblCustomer)
	sb.Select(
		"id",
		"phone",
		"name",
		"balance",
	)

	sqlRows, err := querybuilder.FindAll(ctx, dbtx, sb)
	if err != nil {
		return nil, err
	}

	for sqlRows.Next() {
		var r model.Customer
		sqlRows.Scan(
			&r.ID,
			&r.Phone,
			&r.Name,
			&r.Balance,
		)
		data = append(data, r)
	}

	return data, nil

}

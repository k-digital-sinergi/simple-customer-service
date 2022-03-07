package repo_customer

import (
	"context"
	"database/sql"
	"github.com/huandu/go-sqlbuilder"
	"simple-customer-service/db"
	"simple-customer-service/pkg/model"
	"simple-customer-service/pkg/repository/querybuilder"
	"strconv"
)

func Get(ctx context.Context, dbtx db.DBTX, custID int64) (*model.Customer, error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.From(model.TblCustomer)
	sb.Select(
		"id",
		"phone",
		"name",
		"balance",
	)
	sb.Where(
		sb.Equal("id", strconv.FormatInt(custID, 10)))

	sqlRow, err := querybuilder.Get(ctx, dbtx, sb)
	if err != nil {
		return nil, err
	}

	var r model.Customer
	err = sqlRow.Scan(
		&r.ID,
		&r.Phone,
		&r.Name,
		&r.Balance,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &r, nil

}

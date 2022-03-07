package repo_customer

import (
	"context"
	"errors"
	"fmt"
	"github.com/huandu/go-sqlbuilder"
	"simple-customer-service/db"
	"simple-customer-service/pkg/model"
	"simple-customer-service/pkg/repository/querybuilder"
	"strconv"
)

func Update(ctx context.Context, dbtx db.DBTX, customer model.Customer) (insertID, rowsAffected int64, err error) {
	sb := sqlbuilder.NewUpdateBuilder()
	sb.Update(model.TblCustomer)

	var str []string

	if customer.Phone != "" {
		str = append(str, sb.Assign("phone", customer.Phone))
	}

	if customer.Name != "" {
		str = append(str, sb.Assign("name", customer.Name))
	}

	if customer.Balance != 0 {
		str = append(str, sb.Assign("balance", fmt.Sprintf("%.4f", customer.Balance)))
	}

	if len(str) == 0 {
		return
	}
	sb.Set(str...)

	if customer.ID == 0 {
		return 0, 0, errors.New("customer ID cannot be empty")
	}
	sb.Where(
		sb.Equal("id", strconv.FormatInt(customer.ID, 10)))

	//log.Println(sb.String(), sb.Args)

	insertID, rowsAffected, err = querybuilder.Update(ctx, dbtx, sb)
	if err != nil {
		return
	}

	return
}

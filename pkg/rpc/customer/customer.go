package customer

import (
	"context"
	"errors"
	"simple-customer-service/db"
	"simple-customer-service/pkg/model"
	"simple-customer-service/pkg/repository/repo_customer"
)

type Customer struct{}

func (c *Customer) List(ctx context.Context) ([]model.Customer, error) {
	customers, err := repo_customer.List(ctx, db.GetDatabaseConnection())
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (c *Customer) Get(ctx context.Context, request *model.CustomerGetReq) (*model.Customer, error) {
	customer, err := repo_customer.Get(ctx, db.GetDatabaseConnection(), request.CustID)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (c *Customer) Create(ctx context.Context, request *model.CustomerCUReq) (int64, error) {
	insertID, _, err := repo_customer.Create(ctx, db.GetDatabaseConnection(), request.Customer)
	if err != nil {
		return 0, err
	}

	return insertID, nil
}

func (c *Customer) Update(ctx context.Context, request *model.CustomerCUReq) error {
	_, rowsAffected, err := repo_customer.Update(ctx, db.GetDatabaseConnection(), request.Customer)
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("data tidak ditemukan")
	}

	return nil
}

func (c *Customer) Delete(ctx context.Context, request *model.CustomerDeleteReq) error {
	_, rowsAffected, err := repo_customer.Delete(ctx, db.GetDatabaseConnection(), request.CustID)
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("data tidak ditemukan")
	}

	return nil
}

func New() *Customer {
	return &Customer{}
}

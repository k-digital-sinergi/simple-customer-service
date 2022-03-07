package customer

import (
	"context"
	"simple-customer-service/pkg/model"
)

type Service interface {
	List(ctx context.Context) ([]model.Customer, error)
	Create(ctx context.Context, request *model.CustomerCUReq) (int64, error)
	Update(ctx context.Context, request *model.CustomerCUReq) error
	Delete(ctx context.Context, request *model.CustomerDeleteReq) error
}

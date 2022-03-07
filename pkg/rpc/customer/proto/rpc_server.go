package customer

import (
	"context"
	"encoding/json"
	"simple-customer-service/pkg/model"
	service "simple-customer-service/pkg/rpc/customer"
)

type RpcServer struct {
	svc service.Service
}

func (r *RpcServer) List(ctx context.Context, empty *Empty) (*Response, error) {
	customers, err := r.svc.List(ctx)
	if err != nil {
		return nil, err
	}

	var response Response
	response.Body, err = json.Marshal(customers)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (r *RpcServer) Create(ctx context.Context, request *Request) (*Response, error) {
	var req model.CustomerCUReq
	if err := json.Unmarshal(request.Body, &req); err != nil {
		return nil, err
	}

	custID, err := r.svc.Create(ctx, &req)
	if err != nil {
		return nil, err
	}

	var response Response
	response.Body, err = json.Marshal(custID)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (r *RpcServer) Update(ctx context.Context, request *Request) (*Response, error) {
	var req model.CustomerCUReq
	if err := json.Unmarshal(request.Body, &req); err != nil {
		return nil, err
	}

	if err := r.svc.Update(ctx, &req); err != nil {
		return nil, err
	}

	var response Response

	return &response, nil
}

func (r *RpcServer) Delete(ctx context.Context, request *Request) (*Response, error) {
	var req model.CustomerDeleteReq
	if err := json.Unmarshal(request.Body, &req); err != nil {
		return nil, err
	}

	if err := r.svc.Delete(ctx, &req); err != nil {
		return nil, err
	}

	var response Response

	return &response, nil
}

func (r *RpcServer) mustEmbedUnimplementedCustomerServiceServer() {
	//TODO implement me
	panic("implement me")
}

func NewRpcServer() CustomerServiceServer {
	return &RpcServer{svc: service.New()}
}

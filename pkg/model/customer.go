package model

import "simple-customer-service/db"

const (
	TblCustomer = "customer"
)

type Customer struct {
	ID      int64     `json:"id" uri:"id" binding:"required,min=1"`
	Phone   db.AES256 `json:"phone"`
	Name    string    `json:"name"`
	Balance float64   `json:"balance"`
}

type CustomerGetReq struct {
	CustID int64 `json:"cust_id" uri:"id" binding:"required,min=1"`
}

type CustomerCUReq struct {
	Customer
}

type CustomerDeleteReq struct {
	CustID int64 `json:"cust_id" uri:"id" binding:"required,min=1"`
}

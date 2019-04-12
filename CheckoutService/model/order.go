package model

import (
	"github.com/ajanthan/go-ecommerce-demo/CartService/model"
	payment "github.com/ajanthan/go-ecommerce-demo/PaymentService/model"
	shipping "github.com/ajanthan/go-ecommerce-demo/ShippingService/model"
)

type Order struct {
	UserID         string
	Email          string
	Address        shipping.Address
	CreditCardInfo payment.CreditCard
}

type OrderResult struct {
	OrderID    string
	TrackingID string
	Address    shipping.Address
	Cart       model.Cart
	Cost       float64
}

package dto

import (
	"service-hf-payment-p5/internal/core/domain/entity"
)

type (
	RequestClient struct {
		UUID      string `json:"uuid,omitempty"`
		Name      string `json:"name,omitempty"`
		CPF       string `json:"cpf,omitempty"`
		Email     string `json:"email,omitempty"`
		CreatedAt string ` json:"createdAt,omitempty"`
	}

	RequestPayment struct {
		Price  float64       `json:"price,omitempty"`
		Client RequestClient `json:"payment,omitempty"`
	}

	OutputPayment struct {
		PaymentStatus string              `json:"paymentStatus,omitempty"`
		HTTPStatus    int                 `json:"httpStatus,omitempty"`
		Error         *OutputPaymentError `json:"error,omitempty"`
	}

	OutputPaymentError struct {
		Message string `json:"message,omitempty"`
		Code    string `json:"code,omitempty"`
	}
)

func (r RequestPayment) Payment() entity.Payment {
	return entity.Payment{
		Price: r.Price,
		Client: entity.Client{
			UUID:      r.Client.UUID,
			Name:      r.Client.Name,
			CPF:       r.Client.CPF,
			Email:     r.Client.Email,
			CreatedAt: r.Client.CreatedAt,
		},
	}
}

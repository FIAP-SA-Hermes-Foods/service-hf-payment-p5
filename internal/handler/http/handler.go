package http

import (
	"bytes"
	"encoding/json"
	l "service-hf-payment-p5/external/logger"
	ps "service-hf-payment-p5/external/strings"

	"service-hf-payment-p5/internal/core/domain/entity/dto"
	"fmt"
	"net/http"
)

type ClientHandler interface {
	Handler(rw http.ResponseWriter, req *http.Request)
	HealthCheck(rw http.ResponseWriter, req *http.Request)
}

type paymentHandler struct{}

func NewHandler() ClientHandler { return paymentHandler{} }

func (h paymentHandler) Handler(rw http.ResponseWriter, req *http.Request) {

	var routesClient = map[string]http.HandlerFunc{
		"post hermes_foods/mercado_pago_api": h.proccess,
	}

	handler, err := router(req.Method, req.URL.Path, routesClient)

	if err == nil {
		handler(rw, req)
		return
	}

	rw.WriteHeader(http.StatusNotFound)
	rw.Write([]byte(`{"error": "route ` + req.Method + " " + req.URL.Path + ` not found"} `))
}

func (h paymentHandler) HealthCheck(rw http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		rw.Write([]byte(`{"error": "method not allowed"}`))
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(`{"status": "OK"}`))
}

func (h paymentHandler) proccess(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	if req.Method != http.MethodPost {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		rw.Write([]byte(`{"error": "method not allowed"}`))
		return
	}

	var (
		buff  bytes.Buffer
		input dto.RequestPayment
	)

	if _, err := buff.ReadFrom(req.Body); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to read data body: %v"} `, err)
		return
	}

	if err := json.Unmarshal(buff.Bytes(), &input); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to Unmarshal: %v"} `, err)
		return
	}

	l.Infof("PaymentHandler input: ", " | ", ps.MarshalString(input))

	rw.WriteHeader(http.StatusOK)
	resp := `{"paymentStatus":"Paid","httpStatus":200}`
	l.Infof("PaymentHandler out: ", " | ", resp)
	rw.Write([]byte(resp))
}

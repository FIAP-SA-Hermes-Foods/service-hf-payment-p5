package http

import (
	"io"
	"net/http/httptest"
	"strings"
	"testing"
)

// go test -v -count=1 -failfast -run ^Test_Handler$
func Test_Handler(t *testing.T) {
	type args struct {
		method string
		url    string
		body   io.Reader
	}
	tests := []struct {
		name        string
		args        args
		wantOut     string
		isWantedErr bool
	}{
		{
			name: "success",
			args: args{
				method: "POST",
				url:    "hermes_foods/mercado_pago_api",
				body:   strings.NewReader(`{"name":"Marty", "cpf":"051119995", "email": "martybttf@bttf.com"}`),
			},
			wantOut:     `{"paymentStatus":"Paid","httpStatus":200}`,
			isWantedErr: false,
		},
	}
	for _, tc := range tests {
		h := NewHandler()
		t.Run(tc.name, func(*testing.T) {

			req := httptest.NewRequest(tc.args.method, "/", tc.args.body)
			req.URL.Path = tc.args.url
			rw := httptest.NewRecorder()

			h.Handler(rw, req)

			response := rw.Result()
			defer response.Body.Close()

			b, err := io.ReadAll(response.Body)

			if (!tc.isWantedErr) && err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if strings.TrimSpace(string(b)) != strings.TrimSpace(tc.wantOut) {
				t.Errorf("expected: %s\ngot: %s", tc.wantOut, string(b))

			}
		})
	}
}

// go test -v -count=1 -failfast -run ^Test_HealthCheck$
func Test_HealthCheck(t *testing.T) {
	type args struct {
		method string
		url    string
		body   io.Reader
	}
	tests := []struct {
		name        string
		args        args
		wantOut     string
		isWantedErr bool
	}{
		{
			name: "success",
			args: args{
				method: "GET",
				url:    "/",
				body:   nil,
			},
			wantOut:     `{"status": "OK"}`,
			isWantedErr: false,
		},
		{
			name: "error_method_not_allowed",
			args: args{
				method: "POST",
				url:    "/",
				body:   nil,
			},
			wantOut:     `{"error": "method not allowed"}`,
			isWantedErr: true,
		},
	}

	for _, tc := range tests {
		h := NewHandler()
		t.Run(tc.name, func(*testing.T) {

			req := httptest.NewRequest(tc.args.method, tc.args.url, tc.args.body)
			rw := httptest.NewRecorder()

			h.HealthCheck(rw, req)

			response := rw.Result()
			defer response.Body.Close()

			b, err := io.ReadAll(response.Body)

			if (!tc.isWantedErr) && err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if strings.TrimSpace(string(b)) != strings.TrimSpace(tc.wantOut) {
				t.Errorf("expected: %s\ngot: %s", tc.wantOut, string(b))

			}
		})
	}
}

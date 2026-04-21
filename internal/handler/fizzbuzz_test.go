package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFizzBuzzHandler(t *testing.T) {
	tests := []struct {
		name        string
		method      string
		url         string
		wantStatus  int
		wantCT      string
		wantBody    map[string]string
		wantErrBody string
	}{
		{
			name:       "retorna fizzbuzz",
			method:     http.MethodGet,
			url:        "/fizzbuzz?n=15",
			wantStatus: http.StatusOK,
			wantCT:     "application/json",
			wantBody: map[string]string{
				"message": "FizzBuzz",
			},
		},
		{
			name:       "retorna fizz",
			method:     http.MethodGet,
			url:        "/fizzbuzz?n=9",
			wantStatus: http.StatusOK,
			wantCT:     "application/json",
			wantBody: map[string]string{
				"message": "Fizz",
			},
		},
		{
			name:       "retorna buzz",
			method:     http.MethodGet,
			url:        "/fizzbuzz?n=10",
			wantStatus: http.StatusOK,
			wantCT:     "application/json",
			wantBody: map[string]string{
				"message": "Buzz",
			},
		},
		{
			name:       "retorna numero como string",
			method:     http.MethodGet,
			url:        "/fizzbuzz?n=7",
			wantStatus: http.StatusOK,
			wantCT:     "application/json",
			wantBody: map[string]string{
				"message": "7",
			},
		},
		{
			name:       "retorna mensagem erro número negativo",
			method:     http.MethodGet,
			url:        "/fizzbuzz?n=-7",
			wantStatus: http.StatusOK,
			wantCT:     "application/json",
			wantBody: map[string]string{
				"message": "The number must be positive.",
			},
		},
		{
			name:        "metodo invalido",
			method:      http.MethodPost,
			url:         "/fizzbuzz?n=15",
			wantStatus:  http.StatusMethodNotAllowed,
			wantErrBody: "Method Not Allowed\n",
		},
		{
			name:        "query ausente",
			method:      http.MethodGet,
			url:         "/fizzbuzz",
			wantStatus:  http.StatusBadRequest,
			wantCT:      "text/plain; charset=utf-8",
			wantErrBody: "Attribute error\n",
		},
		{
			name:        "query invalida",
			method:      http.MethodGet,
			url:         "/fizzbuzz?n=abc",
			wantStatus:  http.StatusBadRequest,
			wantCT:      "text/plain; charset=utf-8",
			wantErrBody: "Attribute error\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.url, nil)
			rr := httptest.NewRecorder()

			handler := http.HandlerFunc(FizzBuzzHandler)
			handler.ServeHTTP(rr, req)

			resp := rr.Result()

			if resp.StatusCode != tt.wantStatus {
				t.Errorf("got status %d, want %d", resp.StatusCode, tt.wantStatus)
			}

			if tt.wantCT != "" {
				gotCT := resp.Header.Get("Content-Type")
				if gotCT != tt.wantCT {
					t.Errorf("got Content-Type %q, want %q", gotCT, tt.wantCT)
				}
			}

			if tt.wantBody != nil {
				var got map[string]string
				err := json.NewDecoder(resp.Body).Decode(&got)
				if err != nil {
					t.Fatalf("error decoding response body: %v", err)
				}

				if got["message"] != tt.wantBody["message"] {
					t.Errorf("got body message %q, want %q", got["message"], tt.wantBody["message"])
				}
			}

			if tt.wantErrBody != "" {
				gotBody := rr.Body.String()
				if gotBody != tt.wantErrBody {
					t.Errorf("got body %q, want %q", gotBody, tt.wantErrBody)
				}
			}
		})
	}
}

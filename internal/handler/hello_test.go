package handler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	w := httptest.NewRecorder()

	HelloHandler(w, req)

	resp := w.Result()
	body, err := io.ReadAll((resp.Body))
	if err != nil {
		t.Fatal("failed to read response body: ", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("got status %d, want %d", resp.StatusCode, http.StatusOK)

	}

	wantContentType := "application/json"
	gotContentType := resp.Header.Get("Content-Type")
	if gotContentType != wantContentType {
		t.Errorf("Got Content-Type %q, want %q", gotContentType, wantContentType)
	}

	wantBody := `{"message":"Hello, World!"}` + "\n"
	if string(body) != wantBody {
		t.Errorf("Got body %q, want %q", string(body), wantBody)
	}

}

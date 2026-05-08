package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheckHandler(t *testing.T) {
	// Tạo request giả
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	// ResponseRecorder dùng để ghi lại response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckHandler)

	// Gọi handler
	handler.ServeHTTP(rr, req)

	// Kiểm tra HTTP Status Code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Mã lỗi trả về sai: nhận được %v, mong đợi %v", status, http.StatusOK)
	}

	// Kiểm tra Body
	expected := "OK\n"
	if rr.Body.String() != expected {
		t.Errorf("Body trả về sai: nhận được %v, mong đợi %v", rr.Body.String(), expected)
	}
}

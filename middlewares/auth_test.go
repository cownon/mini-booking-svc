// middlewares/auth_test.go
package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cownon/mini-booking-svc/utils"
)

func TestJWTMiddleware(t *testing.T) {
	// Tạo một handler giả để mô phỏng API được bảo vệ
	dummyHandler := func(w http.ResponseWriter, r *http.Request) {
		// Lấy userID từ context xem middleware có truyền vào đúng không
		userID := r.Context().Value(UserIDKey)
		if userID == nil {
			t.Errorf("Mong đợi userID trong context nhưng không có")
		}
		w.WriteHeader(http.StatusOK)
	}

	// Bọc handler giả bằng middleware của chúng ta
	protectedHandler := JWTMiddleware(dummyHandler)

	t.Run("Thiếu Header", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/", nil)
		rr := httptest.NewRecorder()
		protectedHandler.ServeHTTP(rr, req)

		if rr.Code != http.StatusUnauthorized {
			t.Errorf("Mong đợi 401, nhận được %d", rr.Code)
		}
	})

	t.Run("Token hợp lệ", func(t *testing.T) {
		// Tạo token thật từ utils
		validToken, _ := utils.GenerateToken("user-123")

		req := httptest.NewRequest("GET", "/", nil)
		req.Header.Set("Authorization", "Bearer "+validToken)
		rr := httptest.NewRecorder()

		protectedHandler.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("Mong đợi 200 OK, nhận được %d", rr.Code)
		}
	})
}

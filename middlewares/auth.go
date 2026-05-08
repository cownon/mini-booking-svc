// middlewares/auth.go
package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/cownon/mini-booking-svc/utils"
)

// Tạo một type riêng cho key của context để tránh đụng độ (collision)
type contextKey string

const UserIDKey contextKey = "user_id"

// JWTMiddleware là hàm bao bọc một HTTP Handler khác
func JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1. Lấy header Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Thiếu Authorization header", http.StatusUnauthorized)
			return
		}

		// 2. Kiểm tra định dạng "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Định dạng token không hợp lệ", http.StatusUnauthorized)
			return
		}

		// 3. Xác thực Token
		tokenString := parts[1]
		userID, err := utils.VerifyToken(tokenString)
		if err != nil {
			http.Error(w, "Token không hợp lệ hoặc đã hết hạn", http.StatusUnauthorized)
			return
		}

		// 4. Gắn userID vào Context của Request để Handler phía sau có thể dùng được
		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		reqWithCtx := r.WithContext(ctx)

		// 5. Cho phép đi tiếp vào Handler chính
		next.ServeHTTP(w, reqWithCtx)
	}
}

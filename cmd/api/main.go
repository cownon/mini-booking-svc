// cmd/api/main.go
package api

import (
	"fmt"
	"net/http"

	"github.com/cownon/mini-booking-svc/middlewares"
)

// API không cần bảo vệ
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}

// API cần bảo vệ (Yêu cầu JWT)
func ProtectedProfileHandler(w http.ResponseWriter, r *http.Request) {
	// Trích xuất userID từ context (đã được Middleware nhét vào)
	userID := r.Context().Value(middlewares.UserIDKey)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Xin chào user có ID: %v\n", userID)
}

func main() {
	// Route Public
	http.HandleFunc("/health", HealthCheckHandler)

	// Route Protected (Bọc qua JWTMiddleware)
	http.HandleFunc("/api/profile", middlewares.JWTMiddleware(ProtectedProfileHandler))

	fmt.Println("Server đang chạy tại http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Lỗi khởi động server:", err)
	}
}

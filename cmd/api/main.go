package api

import (
	"fmt"
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}

func main() {
	http.HandleFunc("/health", HealthCheckHandler)
	fmt.Println("Server đang chạy tại http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Lỗi khởi động server:", err)
	}
}

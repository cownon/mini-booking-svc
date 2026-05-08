// db/postgres_test.go
package db

import (
	"context"
	"testing"
	"time"

	"github.com/cownon/mini-booking-svc/models"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestPostgresIntegration(t *testing.T) {
	ctx := context.Background()

	// 1. Dùng Testcontainers khởi tạo một container PostgreSQL thật
	dbName := "booking_test_db"
	dbUser := "user"
	dbPassword := "password"

	postgresContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:15-alpine"),
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second),
		),
	)
	if err != nil {
		t.Fatalf("Không thể khởi động postgres container: %v", err)
	}

	// Đảm bảo container sẽ bị dọn dẹp sau khi test xong
	defer func() {
		if err := postgresContainer.Terminate(ctx); err != nil {
			t.Fatalf("Lỗi khi tắt container: %v", err)
		}
	}()

	// 2. Lấy chuỗi kết nối (Connection String) từ container vừa tạo
	connStr, err := postgresContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		t.Fatalf("Lỗi lấy chuỗi kết nối: %v", err)
	}

	// 3. Chạy hàm kết nối DB của chúng ta
	err = InitDB(connStr)
	if err != nil {
		t.Fatalf("Lỗi InitDB: %v", err)
	}

	// 4. Test logic nghiệp vụ: Thử insert một record vào DB thật
	newBooking := &models.Booking{
		ID:        "b-001",
		UserID:    "user-123", // Dữ liệu hoàn toàn cô lập
		EventID:   "evt-999",
		Status:    "Pending",
		CreatedAt: time.Now(),
	}

	err = CreateBooking(newBooking)
	if err != nil {
		t.Errorf("Lỗi khi tạo booking: %v", err)
	}

	// Lấy ra kiểm tra lại
	var savedBooking models.Booking
	err = DB.First(&savedBooking, "id = ?", "b-001").Error
	if err != nil {
		t.Errorf("Không tìm thấy booking vừa tạo: %v", err)
	}

	if savedBooking.UserID != "user-123" {
		t.Errorf("Dữ liệu lưu sai, mong đợi user-123, nhận được %s", savedBooking.UserID)
	}
}

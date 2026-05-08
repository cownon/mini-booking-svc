// db/postgres.go
package db

import (
	"fmt"
	"log"

	"github.com/cownon/mini-booking-svc/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Hàm khởi tạo kết nối
func InitDB(dsn string) error {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("không thể kết nối database: %v", err)
	}

	// AutoMigrate: Tự động tạo bảng dựa trên struct Booking
	// Đảm bảo tính cô lập: Bảng này tự quản lý data của nó, không foreign keys sang service khác
	err = DB.AutoMigrate(&models.Booking{})
	if err != nil {
		return fmt.Errorf("không thể migrate database: %v", err)
	}

	log.Println("Kết nối và Migrate Database thành công!")
	return nil
}

// Hàm tạo một booking mới
func CreateBooking(booking *models.Booking) error {
	return DB.Create(booking).Error
}

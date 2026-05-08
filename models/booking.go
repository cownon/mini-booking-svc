// models/booking.go
package models

import "time"

type Booking struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`  // Chỉ lưu định danh, giữ trạng thái cô lập hoàn toàn
	EventID   string    `json:"event_id"` // Chỉ lưu định danh
	Status    string    `json:"status"`   // Ví dụ: Pending, Confirmed, Cancelled
	CreatedAt time.Time `json:"created_at"`
}

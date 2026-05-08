package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Trong thực tế, chuỗi secret này phải được giấu trong file .env
// Tuyệt đối không hardcode thế này trên Production nhé!
var secretKey = []byte("my-super-secret-key-for-booking-svc")

// Hàm tạo Token khi user đăng nhập thành công
func GenerateToken(userID string) (string, error) {
	// Khởi tạo thông tin (Claims) muốn lưu trong token
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Thời hạn 24 tiếng
		"iat":     time.Now().Unix(),                     // Thời điểm tạo
	}

	// Tạo token với thuật toán mã hóa HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Ký token bằng Secret Key và trả về chuỗi string
	return token.SignedString(secretKey)
}

// Hàm xác thực và giải mã Token
func VerifyToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Kiểm tra xem thuật toán mã hóa có đúng là HMAC không
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("thuật toán mã hóa không hợp lệ")
		}
		return secretKey, nil
	})

	if err != nil {
		return "", err
	}

	// Trích xuất user_id từ payload nếu token hợp lệ
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := claims["user_id"].(string)
		return userID, nil
	}

	return "", errors.New("token không hợp lệ")
}

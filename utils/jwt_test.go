package utils

import (
	"testing"
)

func TestGenerateAndVerifyToken(t *testing.T) {
	userID := "user-999"

	// 1. Test sinh token
	tokenString, err := GenerateToken(userID)
	if err != nil {
		t.Fatalf("Lỗi khi tạo token: %v", err)
	}
	if tokenString == "" {
		t.Error("Token sinh ra bị rỗng")
	}

	// 2. Test giải mã token thành công
	extractedUserID, err := VerifyToken(tokenString)
	if err != nil {
		t.Fatalf("Lỗi khi xác thực token hợp lệ: %v", err)
	}
	if extractedUserID != userID {
		t.Errorf("UserID giải mã sai: mong đợi %s, nhận được %s", userID, extractedUserID)
	}

	// 3. Test giải mã token sai
	fakeToken := "header.fake_payload.fake_signature"
	_, err = VerifyToken(fakeToken)
	if err == nil {
		t.Error("Mong đợi lỗi khi dùng token giả mạo nhưng không thấy lỗi")
	}
}

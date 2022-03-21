package entity

import "time"

// ========================= request =========================

type UploadFileRequest struct {
	AccountID uint64   `json:"account_id" binding:"required,gt=0"`
	Addresses []string `json:"address"`
	// todo ...
}

type GetFileListRequest struct {
	AccountID uint64 `json:"account_id" binding:"required,gt=0"`
	Address   string `json:"address"`
	// todo ...
}

// ========================= response =========================

type GetFileListResponse struct {
	AccountID uint64    `json:"account_id"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
}

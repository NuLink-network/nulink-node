package entity

import "time"

// ========================= request =========================

type File struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
}

type UploadFileRequest struct {
	Files     []File `json:"file" binding:"required"`
	AccountID string `json:"account_id" binding:"required"`
	FileOwner string `json:"file_owner" binding:"required"`
	PolicyID  string `json:"policy_id" binding:"required"`
	Signature string `json:"signature" binding:"required"`
}

type Paginate struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type GetFileListRequest struct {
	AccountID string `json:"account_id" binding:"required"`
	FileName  string `json:"file_name"`
	//UploadTime uint64   `json:"upload_time"`
	Paginate Paginate `json:"paginate"`
}

type GetOthersFileListRequest struct {
	AccountID uint64 `json:"account_id" binding:"required,gt=0"`
	// todo ...
}

type DeleteFileRequest struct {
	FileID    uint64 `json:"file_id" binding:"required,gt=0"`
	AccountID uint64 `json:"account_id" binding:"required,gt=0"`
	Signature string `json:"signature"`
}

// ========================= response =========================

type GetFileListResponse struct {
	AccountID string    `json:"account_id"`
	FileName  string    `json:"file_name"`
	Address   string    `json:"address"`
	Thumbnail string    `json:"thumbnail"`
	CreatedAt time.Time `json:"created_at"`
}

type GetOthersFileListResponse struct {
	AccountID uint64    `json:"account_id"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
}

package entity

import "time"

// ========================= request =========================

type File struct {
	ID      string `json:"id" binding:"required"`
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
}

type UploadFileRequest struct {
	Files     []File `json:"file" binding:"required"`
	AccountID string `json:"account_id" binding:"required"`
	PolicyID  string `json:"policy_id" binding:"required"`
	Signature string `json:"signature" binding:"required"`
}

type CreatePolicyAndUploadFileRequest struct {
	Files       []File `json:"file" binding:"required"`
	AccountID   string `json:"account_id" binding:"required"`
	FileOwner   string `json:"file_owner" binding:"required"`
	PolicyID    string `json:"policy_id" binding:"required"`
	PolicyLabel string `json:"policy_label" binding:"required"`
	EncryptedPK string `json:"encrypted_pk" binding:"required"`
	Signature   string `json:"signature" binding:"required"`
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
	AccountID string `json:"account_id" binding:"required"`
	FileName  string `json:"file_name"`
	//UploadTime uint64   `json:"upload_time"`
	Paginate Paginate `json:"paginate"`
}

type DeleteFileRequest struct {
	FileIDs   []string `json:"file_id" binding:"required"`
	AccountID string   `json:"account_id" binding:"required"`
	Signature string   `json:"signature"`
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
	AccountID string    `json:"account_id"`
	FileName  string    `json:"file_name"`
	Address   string    `json:"address"`
	Thumbnail string    `json:"thumbnail"`
	CreatedAt time.Time `json:"created_at"`
}

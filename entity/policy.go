package entity

import "time"

// ========================= request =========================

type CreatePolicyRequest struct {
	AccountID uint64 `json:"account_id" binding:"required,gt=0"`
	//Name        string `json:"name" binding:"required"`
	PolicyID    uint64 `json:"policy_id" binding:"required,gt=0"`
	Label       string `json:"label"`
	EncryptedPK string `json:"encrypted_pk" binding:"required"`
	VerifyPK    string `json:"verify_pk" binding:"required"`
	Signature   string `json:"signature" binding:"required"`
}

type RevokePolicyRequest struct {
	AccountID string `json:"account_id" binding:"required"`
	PolicyID  uint64 `json:"policy_id" binding:"required"`
	Signature string `json:"signature" binding:"required"`
}

type PolicyListRequest struct {
	PolicyID    uint64   `json:"policy_id"`
	PolicyLabel string   `json:"policy_label"`
	CreatorID   string   `json:"account_id"`
	ConsumerID  string   `json:"consumer_id"`
	Status      uint8    `json:"status" binding:"gt=0,lt=2"`
	Paginate    Paginate `json:"paginate"`
}

type FileDetailListRequest struct {
	PolicyID   uint64   `json:"policy_id"`
	CreatorID  string   `json:"creator_id"`
	ConsumerID string   `json:"consumer_id"`
	Paginate   Paginate `json:"paginate"`
}

// ========================= response =========================

type PolicyListResponse struct {
	Hrac       string    `json:"hrac"`
	Label      string    `json:"label"`
	PolicyID   uint64    `json:"policy_id"`
	Creator    string    `json:"creator"`
	CreatorID  string    `json:"creator_id"`
	Consumer   string    `json:"consumer"`
	ConsumerID string    `json:"consumer_id"`
	Gas        string    `json:"gas"`
	TxHash     string    `json:"tx_hash"`
	StartAt    int64     `json:"start_at"`
	EndAt      int64     `json:"end_at"`
	CreatedAt  time.Time `json:"created_at"`
}

type FileDetailListResponse struct {
	FileID        string `json:"file_id"`
	FileName      string `json:"file_name"`
	Owner         string `json:"owner"`
	OwnerID       string `json:"owner_id"`
	Address       string `json:"address"`
	Thumbnail     string `json:"thumbnail"`
	CreatedAt     int64  `json:"created_at"`
	PolicyID      uint64 `json:"policy_id"`
	PolicyHrac    string `json:"policy_hrac"`
	PolicyStartAt int64  `json:"policy_start_at"`
	PolicyEndAt   int64  `json:"policy_end_at"`
}

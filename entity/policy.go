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
	PolicyID  string `json:"policy_id" binding:"required,"`
	Signature string `json:"signature" binding:"required"`
}

type PolicyListRequest struct {
	PolicyID   string   `json:"policy_id"`
	CreatorID  string   `json:"account_id"`
	ConsumerID string   `json:"consumer_id"`
	Status     uint8    `json:"status" binding:"gt=0 lt=2"`
	Paginate   Paginate `json:"paginate"`
}

// ========================= response =========================

type PolicyListResponse struct {
	Hrac             string    `json:"hrac"`
	Label            string    `json:"label"`
	PolicyID         string    `json:"policy_id"`
	Creator          string    `json:"creator"`
	CreatorID        string    `json:"creator_id"`
	Consumer         string    `json:"consumer"`
	ConsumerID       string    `json:"consumer_id"`
	EncryptedPK      string    `json:"encryptedPK"`
	EncryptedAddress string    `json:"encrypted_address"`
	Status           uint8     `json:"status"`
	Gas              string    `json:"gas"`
	TxHash           string    `json:"tx_hash"`
	CreatedAt        time.Time `json:"created_at"`
}

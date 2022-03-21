package entity

import "time"

// ========================= request =========================

type CreateAccountRequest struct {
	Name         string `json:"name" binding:"required"`
	EthereumAddr string `json:"ethereum_addr" binding:"required"`
	EncryptedPK  string `json:"encrypted_pk" binding:"required"`
	VerifyPK     string `json:"verify_pk" binding:"required"`
	Signature    string `json:"signature" binding:"required"`
}

type GetAccountRequest struct {
	AccountID uint64 `json:"account_id" binding:"required,gt=0"`
}

// ========================= response =========================

type CreateAccountResponse struct {
	AccountID uint64 `json:"account_id"`
}

type GetAccountResponse struct {
	ID           uint64    `json:"id"`
	Name         string    `json:"name"`
	EthereumAddr string    `json:"ethereum_addr"`
	EncryptedPK  string    `json:"encrypted_pk"`
	VerifyPK     string    `json:"verify_pk"`
	Signature    string    `json:"signature"`
	CreatedAt    time.Time `json:"created_at"`
}

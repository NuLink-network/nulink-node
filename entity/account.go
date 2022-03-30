package entity

import "time"

// ========================= request =========================

type CreateAccountRequest struct {
	Name         string `json:"name" binding:"required"`
	Account      string `json:"account" binding:"required"`
	EthereumAddr string `json:"ethereum_addr" binding:"required"`
	EncryptedPK  string `json:"encrypted_pk" binding:"required"`
	VerifyPK     string `json:"verify_pk" binding:"required"`
}

type GetAccountRequest struct {
	AccountID uint64 `json:"account_id" binding:"required,gt=0"`
}

type AccountIsExistRequest struct {
	AccountID    uint64 `json:"account_id" binding:"required,gt=0"`
	Name         string `json:"name" binding:"required"`
	EthereumAddr string `json:"ethereum_addr" binding:"required"`
	EncryptedPK  string `json:"encrypted_pk" binding:"required"`
	VerifyPK     string `json:"verify_pk" binding:"required"`
}

// ========================= response =========================

type GetAccountResponse struct {
	ID           uint64    `json:"id"`
	Name         string    `json:"name"`
	EthereumAddr string    `json:"ethereum_addr"`
	EncryptedPK  string    `json:"encrypted_pk"`
	VerifyPK     string    `json:"verify_pk"`
	Signature    string    `json:"signature"`
	CreatedAt    time.Time `json:"created_at"`
}

type AccountIsExistResponse struct {
	IsExist bool `json:"is_exist"`
}

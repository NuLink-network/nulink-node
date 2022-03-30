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
	Account string `json:"account" binding:"required"`
}

type AccountIsExistRequest struct {
	Name         string `json:"name" binding:"required"`
	Account      string `json:"account"`
	EthereumAddr string `json:"ethereum_addr" binding:"required"`
	EncryptedPK  string `json:"encrypted_pk" binding:"required"`
	VerifyPK     string `json:"verify_pk" binding:"required"`
}

// ========================= response =========================

type GetAccountResponse struct {
	Name         string    `json:"name"`
	Account      string    `json:"account"`
	EthereumAddr string    `json:"ethereum_addr"`
	EncryptedPK  string    `json:"encrypted_pk"`
	VerifyPK     string    `json:"verify_pk"`
	Status       int8      `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
}

type AccountIsExistResponse struct {
	IsExist bool `json:"is_exist"`
}

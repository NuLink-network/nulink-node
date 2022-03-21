package entity

// ========================= request =========================

type CreateAccountRequest struct {
	Name         string `json:"name" binding:"required"`
	EthereumAddr string `json:"ethereum_addr" binding:"required"`
	EncryptedPK  string `json:"encrypted_pk" binding:"required"`
	VerifyPK     string `json:"verify_pk" binding:"required"`
	Signature    string `json:"signature" binding:"required"`
}

// ========================= response =========================

type CreateAccountResponse struct {
	AccountID uint64 `json:"account_id"`
}

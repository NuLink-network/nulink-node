package entity

// ========================= request =========================

type CreateAccountRequest struct {
	Name         string `json:"name" binding:"required"`
	AccountID    string `json:"account_id" binding:"required"`
	EthereumAddr string `json:"ethereum_addr" binding:"required"`
	EncryptedPK  string `json:"encrypted_pk" binding:"required"`
	VerifyPK     string `json:"verify_pk" binding:"required"`
}

type GetAccountRequest struct {
	AccountID string `json:"account_id" binding:"required"`
}

type AccountIsExistRequest struct {
	Name         string `json:"name" binding:"required"`
	AccountID    string `json:"account_id" binding:"required"`
	EthereumAddr string `json:"ethereum_addr" binding:"required"`
	EncryptedPK  string `json:"encrypted_pk" binding:"required"`
	VerifyPK     string `json:"verify_pk" binding:"required"`
}

type UpdateAccountRequest struct {
	AccountID string `json:"account_id" binding:"required"`
	Avatar    string `json:"avatar"`
	Name      string `json:"name"`
	UserSite  string `json:"user_site"`
	Twitter   string `json:"twitter"`
	Instagram string `json:"instagram"`
	Facebook  string `json:"facebook"`
	Profile   string `json:"profile"`
}

// ========================= response =========================

type GetAccountResponse struct {
	Name         string `json:"name"`
	AccountID    string `json:"account_id"`
	EthereumAddr string `json:"ethereum_addr"`
	EncryptedPK  string `json:"encrypted_pk"`
	VerifyPK     string `json:"verify_pk"`
	Status       int8   `json:"status"`
	Avatar       string `json:"avatar"`
	UserSite     string `json:"user_site"`
	Twitter      string `json:"twitter"`
	Instagram    string `json:"instagram"`
	Facebook     string `json:"facebook"`
	Profile      string `json:"profile"`
	CreatedAt    int64  `json:"created_at"`
}

type AccountIsExistResponse struct {
	IsExist bool `json:"is_exist"`
}

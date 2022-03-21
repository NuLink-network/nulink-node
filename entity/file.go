package entity

// ========================= request =========================

type UploadFileRequest struct {
	AccountID uint64   `json:"account_id" binding:"required,gt=0"`
	Addresses []string `json:"address"`
	// todo ...
}

// ========================= response =========================

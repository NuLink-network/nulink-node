package entity

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

// ========================= response =========================

//type CreatePolicyListResponse struct {
//	PolicyID  uint64    `json:"policy_id"`
//}

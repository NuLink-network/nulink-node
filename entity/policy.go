package entity

// ========================= request =========================

type RevokePolicyRequest struct {
	AccountID string `json:"account_id" binding:"required"`
	PolicyID  uint64 `json:"policy_id" binding:"required"`
	Signature string `json:"signature" binding:"required"`
}

type PolicyListRequest struct {
	PolicyID      uint64   `json:"policy_id"`
	PolicyLabelID string   `json:"policy_label_id"`
	CreatorID     string   `json:"creator_id"`
	ConsumerID    string   `json:"consumer_id"`
	Paginate      Paginate `json:"paginate"`
}

type FileDetailListRequest struct {
	PolicyID uint64   `json:"policy_id" binding:"required"`
	Paginate Paginate `json:"paginate"`
}

// ========================= response =========================

type PolicyListResponse struct {
	Hrac            string `json:"hrac"`
	PolicyID        uint64 `json:"policy_id"`
	Creator         string `json:"creator"`
	CreatorID       string `json:"creator_id"`
	CreatorAddress  string `json:"creator_address"`
	Consumer        string `json:"consumer"`
	ConsumerID      string `json:"consumer_id"`
	ConsumerAddress string `json:"consumer_address"`
	Gas             string `json:"gas"`
	TxHash          string `json:"tx_hash"`
	EncryptedPK     string `json:"encrypted_pk"`
	StartAt         int64  `json:"start_at"`
	EndAt           int64  `json:"end_at"`
	CreatedAt       int64  `json:"created_at"`
}

type FileDetailListResponse struct {
	FileID        string `json:"file_id"`
	FileName      string `json:"file_name"`
	Owner         string `json:"owner"`
	OwnerID       string `json:"owner_id"`
	OwnerAvatar   string `json:"owner_avatar"`
	Address       string `json:"address"`
	Thumbnail     string `json:"thumbnail"`
	CreatedAt     int64  `json:"created_at"`
	PolicyID      uint64 `json:"policy_id"`
	PolicyHrac    string `json:"policy_hrac"`
	PolicyStartAt int64  `json:"policy_start_at"`
	PolicyEndAt   int64  `json:"policy_end_at"`
}

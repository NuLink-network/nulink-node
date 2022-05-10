package entity

// ========================= request =========================

type ApplyFileRequest struct {
	FileIDs    []string `json:"file_ids" binding:"required"`
	ProposerID string   `json:"proposer_id" binding:"required"`
	StartAt    int64    `json:"start_at" binding:"required"`
	EndAt      int64    `json:"end_at" binding:"required"`
	Signature  string   `json:"signature" binding:"required"`
}

type ApplyFileListRequest struct {
	FileID      string   `json:"file_id"`
	Status      uint8    `json:"status" binding:"gte=0,lte=3"`
	ProposerID  string   `json:"proposer_id"`
	FileOwnerID string   `json:"file_owner_id"`
	Paginate    Paginate `json:"paginate"`
}

type RevokeApplyRequest struct {
	ProposerID string   `json:"proposer_id"`
	ApplyIDs   []uint64 `json:"apply_ids" binding:"required"`
}

type Policy struct {
	Hrac             string `json:"hrac" binding:"required"`
	Gas              string `json:"gas" binding:"required"`
	TxHash           string `json:"tx_hash" binding:"required"`
	ConsumerID       string `json:"consumer_id" binding:"required"`
	EncryptedPK      string `json:"encrypted_pk" binding:"required"`
	EncryptedAddress string `json:"encrypted_address" binding:"required"`
}

type ApproveApplyRequest struct {
	AccountID string `json:"account_id" binding:"required"`
	ApplyID   uint64 `json:"apply_id" binding:"required,gt=0"`
	Policy    Policy `json:"policy" binding:"required"`
	Signature string `json:"signature" binding:"required"`
}

type RejectApplyRequest struct {
	AccountID string `json:"account_id" binding:"required"`
	ApplyID   uint64 `json:"apply_id" binding:"required,gt=0"`
	Signature string `json:"signature" binding:"required"`
}

// ========================= response =========================

type ApplyFileListResponse struct {
	ApplyID     uint64 `json:"apply_id"`
	FileID      string `json:"file_id"`
	Proposer    string `json:"proposer"`
	ProposerID  string `json:"proposer_id"`
	FileOwner   string `json:"file_owner"`
	FileOwnerID string `json:"file_owner_id"`
	StartAt     int64  `json:"start_at"`
	EndAt       int64  `json:"end_at"`
	CreatedAt   int64  `json:"created_at"`
}

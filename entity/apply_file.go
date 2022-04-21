package entity

import "time"

// ========================= request =========================

type ApplyFileRequest struct {
	FileIDs    []string  `json:"file_ids" binding:"required"`
	ProposerID string    `json:"proposer_id" binding:"required"`
	StartAt    time.Time `json:"start_at" binding:"required"`
	FinishAt   time.Time `json:"finish_at" binding:"required"`
	Signature  string    `json:"signature" binding:"required"`
}

type ApplyFileListRequest struct {
	FileID      string   `json:"file_id"`
	Status      uint8    `json:"status" binding:"gt=0 lt=3"`
	ProposerID  string   `json:"proposer_id"`
	FileOwnerID string   `json:"file_owner_id"`
	Paginate    Paginate `json:"paginate"`
}

type RevokeApplyRequest struct {
	ProposerID string   `json:"proposer_id"`
	ApplyIDs   []uint64 `json:"apply_ids" binding:"required"`
}

type Policy struct {
	ID               string `json:"id" binding:"required"`
	Hrac             string `json:"hrac" binding:"required"`
	Gas              string `json:"Gas" binding:"required"`
	TxHash           string `json:"tx_hash" binding:"required"`
	CreatorID        string `json:"publisher_id" binding:"required"`
	EncryptedAddress string `json:"address" binding:"required"`
}

type ApproveApplyRequest struct {
	ApplyID uint64 `json:"apply_id" binding:"required,gt=0"`
	Policy  Policy `json:"policy" binding:"required"`
}

type RejectApplyRequest struct {
	AccountID string `json:"accountID" binding:"required"`
	ApplyID   uint64 `json:"apply_id" binding:"required,gt=0"`
	Signature string `json:"signature" binding:"required"`
}

// ========================= response =========================

type ApplyFileListResponse struct {
	ApplyID     uint64    `json:"apply_id"`
	FileID      string    `json:"file_id"`
	Proposer    string    `json:"proposer"`
	ProposerID  string    `json:"proposer_id"`
	FileOwner   string    `json:"file_owner"`
	FileOwnerID string    `json:"file_owner_id"`
	StartAt     time.Time `json:"start_at"`
	FinishAt    time.Time `json:"finish_at"`
	CreatedAt   time.Time `json:"created_at"`
}

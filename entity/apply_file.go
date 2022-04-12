package entity

import "time"

// ========================= request =========================

type ApplyFileRequest struct {
	FileIDs    []uint64  `json:"file_ids" binding:"required"`
	ProposerID uint64    `json:"proposer_id" binding:"required,gt=0"`
	Signature  string    `json:"signature" binding:"required"`
	StartAt    time.Time `json:"start_at" binding:"required"`
	FinishAt   time.Time `json:"finish_at" binding:"required"`
}

type ApplyFileListRequest struct {
	FileID              uint64   `json:"file_id" binding:"required,gt=0"`
	Status              uint8    `json:"status" binding:"required,gt=0 lt=3"`
	ProposerAccountID   string   `json:"proposer_account_id"`
	ProprietorAccountID string   `json:"proprietor_account_id"`
	Paginate            Paginate `json:"paginate"`
}

type RevokeApplyRequest struct {
	ProposerAccountID string   `json:"proposer_account_id"`
	ApplyIDs          []uint64 `json:"apply_ids" binding:"required"`
}

type ApproveApplyRequest struct {
	ApplyID uint64 `json:"apply_id" binding:"required,gt=0"`
}

type RejectApplyRequest struct {
	AccountID string `json:"accountID" binding:"required,gt=0"`
	ApplyID   uint64 `json:"apply_id" binding:"required,gt=0"`
	Signature string `json:"signature" binding:"required"`
}

// ========================= response =========================

type ApplyFileListResponse struct {
	ApplyID             uint64    `json:"apply_id"`
	FileID              uint64    `json:"file_id"`
	Proposer            string    `json:"proposer"`
	ProposerAccountID   string    `json:"proposer_account_id"`
	Proprietor          string    `json:"proprietor"`
	ProprietorAccountID string    `json:"proprietor_account_id"`
	StartAt             time.Time `json:"start_at"`
	FinishAt            time.Time `json:"finish_at"`
	CreatedAt           time.Time `json:"created_at"`
}

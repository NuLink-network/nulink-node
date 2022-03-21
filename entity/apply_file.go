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
	ProposerID uint64 `json:"proposer_id" binding:"required,gt=0"`
	Status     int8   `json:"status" binding:"required,gt=-2lt=3"`
}

type RevokeApplyRequest struct {
	ApplyIDs []uint64 `json:"apply_ids" binding:"required"`
}

// ========================= response =========================

type ApplyFileListResponse struct {
	ApplyID    uint64 `json:"apply_id"`
	FileID     uint64 `json:"file_ids"`
	ProposerID uint64 `json:"proposer_id"`
	//Signature string    `json:"signature"`
	StartAt   time.Time `json:"start_at"`
	FinishAt  time.Time `json:"finish_at"`
	CreatedAt time.Time `json:"created_at"`
}

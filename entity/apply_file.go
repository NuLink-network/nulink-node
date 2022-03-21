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

// ========================= response =========================

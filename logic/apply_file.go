package logic

import (
	"github.com/NuLink-network/nulink-node/entity"
	"time"

	"github.com/NuLink-network/nulink-node/dao"
)

func ApplyFile(fileIDs []uint64, proposerID uint64, signature string, startAt, finishAt time.Time) error {
	af := &dao.AppleFile{}
	afs := make([]*dao.AppleFile, 0, len(fileIDs))
	for _, fid := range fileIDs {
		afs = append(afs, &dao.AppleFile{
			FileID:     fid,
			ProposerID: proposerID,
			Signature:  signature,
			StartAt:    startAt,
			FinishAt:   finishAt,
		})
	}
	return af.BatchCreate(afs)
}

func ApplyFileList(proposerID uint64, status int8) ([]*entity.ApplyFileListResponse, error) {
	af := &dao.AppleFile{ProposerID: proposerID}
	if status != -1 {
		af.Status = status
	}
	afs, err := af.Find()
	if err != nil {
		return nil, err
	}

	resp := make([]*entity.ApplyFileListResponse, 0, len(afs))
	for _, af := range afs {
		resp = append(resp, &entity.ApplyFileListResponse{
			ApplyID:    af.ID,
			FileID:     af.FileID,
			ProposerID: af.ProposerID,
			StartAt:    af.StartAt,
			FinishAt:   af.FinishAt,
			CreatedAt:  af.CreatedAt,
		})
	}
	return resp, nil
}

func RevokeApply(applyID []uint64) error {
	af := &dao.AppleFile{}
	return af.BatchDelete(applyID)
}

func ApproveApply(applyID uint64) error {
	af := &dao.AppleFile{ID: applyID}
	newAf := &dao.AppleFile{Status: dao.StatusApprove}
	return af.Updates(newAf)
}

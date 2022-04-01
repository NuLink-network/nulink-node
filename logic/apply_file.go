package logic

import (
	"github.com/NuLink-network/nulink-node/entity"
	"time"

	"github.com/NuLink-network/nulink-node/dao"
)

func ApplyFile(fileIDs []uint64, proposerID uint64, signature string, startAt, finishAt time.Time) error {
	af := &dao.ApplyFile{}
	afs := make([]*dao.ApplyFile, 0, len(fileIDs))
	for _, fid := range fileIDs {
		afs = append(afs, &dao.ApplyFile{
			FileID:     fid,
			ProposerID: proposerID,
			StartAt:    startAt,
			FinishAt:   finishAt,
		})
	}
	return af.BatchCreate(afs)
}

func ApplyFileList(proposerID uint64, status int8) ([]*entity.ApplyFileListResponse, error) {
	af := &dao.ApplyFile{ProposerID: proposerID}
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
	af := &dao.ApplyFile{}
	return af.BatchDelete(applyID)
}

func ApproveApply(applyID uint64) error {
	af := &dao.ApplyFile{ID: applyID}
	newAf := &dao.ApplyFile{Status: dao.StatusApprove}
	return af.Updates(newAf)
}

func RejectApply(applyID uint64) error {
	af := &dao.ApplyFile{ID: applyID}
	newAf := &dao.ApplyFile{Status: dao.StatusReject}
	return af.Updates(newAf)
}

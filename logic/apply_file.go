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

func ApplyFileList(fileID uint64, status uint8, proposerAccountID, proprietorAccountID string, page, pageSize int) ([]*entity.ApplyFileListResponse, error) {
	af := &dao.ApplyFile{
		FileID:              fileID,
		ProposerAccountID:   proposerAccountID,
		ProprietorAccountID: proprietorAccountID,
		Status:              status,
	}
	if status != dao.StatusAll {
		af.Status = status
	}
	afs, err := af.Find(page, pageSize)
	if err != nil {
		return nil, err
	}

	resp := make([]*entity.ApplyFileListResponse, 0, len(afs))
	for _, af := range afs {
		resp = append(resp, &entity.ApplyFileListResponse{
			ApplyID:             af.ID,
			FileID:              af.FileID,
			Proposer:            af.Proposer,
			ProposerAccountID:   af.ProposerAccountID,
			Proprietor:          af.Proprietor,
			ProprietorAccountID: af.ProprietorAccountID,
			StartAt:             af.StartAt,
			FinishAt:            af.FinishAt,
			CreatedAt:           af.CreatedAt,
		})
	}
	return resp, nil
}

func RevokeApply(proposerAccountID string, applyIDs []uint64) error {
	af := &dao.ApplyFile{
		ProposerAccountID: proposerAccountID,
	}
	return af.BatchDelete(applyIDs)
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

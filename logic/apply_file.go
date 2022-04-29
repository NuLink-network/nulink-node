package logic

import (
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/NuLink-network/nulink-node/dao"
	"github.com/NuLink-network/nulink-node/entity"
	"github.com/NuLink-network/nulink-node/resp"
)

func ApplyFile(fileIDs []string, proposerID string, startAt, finishAt int64) (code int) {
	files, err := dao.NewFile().FindAny("file_id in ?", fileIDs)
	if err != nil {
		return resp.CodeInternalServerError
	}
	fileOwner := make(map[string]string, len(files))
	for _, file := range files {
		fileOwner[file.FileID] = file.OwnerID
	}

	afs := make([]*dao.ApplyFile, 0, len(fileIDs))
	for _, fid := range fileIDs {
		afs = append(afs, &dao.ApplyFile{
			FileID:      fid,
			ProposerID:  proposerID,
			FileOwnerID: fileOwner[fid],
			StartAt:     time.Unix(startAt, 0),
			FinishAt:    time.Unix(finishAt, 0),
		})
	}
	if err := dao.NewAppleFile().BatchCreate(afs); err != nil {
		return resp.CodeInternalServerError
	}
	return resp.CodeSuccess
}

func ApplyFileList(fileID string, status uint8, proposerID, fileOwnerID string, page, pageSize int) ([]*entity.ApplyFileListResponse, error) {
	af := &dao.ApplyFile{
		FileID:      fileID,
		ProposerID:  proposerID,
		FileOwnerID: fileOwnerID,
		Status:      status,
	}
	if status != dao.ApplyStatusAll {
		af.Status = status
	}
	afs, err := af.Find(page, pageSize)
	if err != nil {
		return nil, err
	}

	ret := make([]*entity.ApplyFileListResponse, 0, len(afs))
	for _, af := range afs {
		ret = append(ret, &entity.ApplyFileListResponse{
			ApplyID:     af.ID,
			FileID:      af.FileID,
			Proposer:    af.Proposer,
			ProposerID:  af.ProposerID,
			FileOwner:   af.FileOwner,
			FileOwnerID: af.FileOwnerID,
			StartAt:     af.StartAt.Unix(),
			FinishAt:    af.FinishAt.Unix(),
			CreatedAt:   af.CreatedAt.Unix(),
		})
	}
	return ret, nil
}

func RevokeApply(proposerID string, applyIDs []uint64) (code int) {
	af := &dao.ApplyFile{
		ProposerID: proposerID,
	}

	apply, err := af.Get()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp.CodeApplyNotExist
		}
		return resp.CodeInternalServerError
	}
	if apply.Status != dao.ApplyStatusApplying {
		return resp.CodeApplyIrrevocable
	}

	if err := af.BatchDelete(applyIDs); err != nil {
		return resp.CodeInternalServerError
	}

	return resp.CodeSuccess
}

func ApproveApply(accountID string, applyID uint64, policy entity.Policy) (code int) {
	af := &dao.ApplyFile{
		ID:          applyID,
		FileOwnerID: accountID,
	}
	apply, err := af.Get()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp.CodeApplyNotExist
		}
		return resp.CodeInternalServerError
	}
	if apply.Status == dao.ApplyStatusRejected {
		return resp.CodeApplyRejected
	}
	if apply.Status == dao.ApplyStatusApproved {
		return resp.CodeSuccess
	}

	p := &dao.Policy{
		PolicyID: policy.ID,
	}
	pl, err := p.Get()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp.CodeApplyNotExist
		}
		return resp.CodeInternalServerError
	}
	if pl.CreatorID != accountID {

	}

	newPolicy := &dao.Policy{
		PolicyID:         policy.ID,
		Hrac:             policy.Hrac,
		Gas:              policy.Gas,
		TxHash:           policy.TxHash,
		ConsumerID:       policy.ConsumerID,
		EncryptedAddress: policy.EncryptedAddress,
	}
	if err := dao.NewPolicy().Updates(newPolicy); err != nil {
		return resp.CodeInternalServerError
	}

	newAf := &dao.ApplyFile{Status: dao.ApplyStatusApproved}
	if err := af.Updates(newAf); err != nil {
		return resp.CodeInternalServerError
	}

	return resp.CodeSuccess
}

func RejectApply(accountID string, applyID uint64) (code int) {
	af := &dao.ApplyFile{
		ID:          applyID,
		FileOwnerID: accountID,
	}

	apply, err := af.Get()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp.CodeApplyNotExist
		}
		return resp.CodeInternalServerError
	}
	if apply.Status == dao.ApplyStatusApproved {
		return resp.CodeApplyApproved
	}
	if apply.Status == dao.ApplyStatusRejected {
		return resp.CodeSuccess
	}

	newAf := &dao.ApplyFile{Status: dao.ApplyStatusRejected}
	if err := af.Updates(newAf); err != nil {
		return resp.CodeInternalServerError
	}
	return resp.CodeSuccess
}

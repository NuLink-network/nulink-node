package logic

import (
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/NuLink-network/nulink-node/dao"
	"github.com/NuLink-network/nulink-node/entity"
	"github.com/NuLink-network/nulink-node/resource/log"
	"github.com/NuLink-network/nulink-node/resp"
)

func ApplyFile(fileIDs []string, proposerID string, startAt, finishAt int64) (code int) {
	fp := &dao.FilePolicy{
		ConsumerID: proposerID,
	}
	query := &dao.QueryExtra{
		Conditions: map[string]interface{}{
			"file_id in ?": fileIDs,
			"end_at > ":    finishAt,
		},
	}
	filePolicyList, err := fp.FindAny(query, nil)
	if err != nil {
		log.Logger().WithField("file policy", fp).WithField("query", query).WithField("error", err).Error("find file policy list failed")
		return resp.CodeInternalServerError
	}
	ignoreFileIDs := make(map[string]struct{}, len(filePolicyList))
	for _, fp := range filePolicyList {
		ignoreFileIDs[fp.FileID] = struct{}{}
	}

	acc := &dao.Account{Account: proposerID}
	account, err := acc.Get()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp.CodeAccountNotExist
		}
		log.Logger().WithField("account", acc).WithField("error", err).Error("get account failed")
		return resp.CodeInternalServerError
	}

	query = &dao.QueryExtra{
		Conditions: map[string]interface{}{
			"file_id in ?": fileIDs,
		},
	}
	files, err := dao.NewFile().FindAny(query, nil)
	if err != nil {
		return resp.CodeInternalServerError
	}
	fileOwner := make(map[string]map[string]string, len(files))
	for _, file := range files {
		fileOwner[file.FileID] = map[string]string{
			"owner":    file.Owner,
			"owner_id": file.OwnerID,
		}
	}

	afs := make([]*dao.ApplyFile, 0, len(fileIDs))
	for _, fid := range fileIDs {
		_, ok := ignoreFileIDs[fid]
		if ok {
			continue
		}
		afs = append(afs, &dao.ApplyFile{
			FileID:      fid,
			Proposer:    account.Account,
			ProposerID:  proposerID,
			FileOwner:   fileOwner[fid]["owner"],
			FileOwnerID: fileOwner[fid]["owner_id"],
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
	afs, err := af.Find(dao.Paginate(page, pageSize))
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

	f := &dao.File{
		FileID: af.FileID,
	}
	file, err := f.Get()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp.CodeFileNotExist
		}
		log.Logger().WithField("file", f).WithField("error", err).Error("get file failed")
		return resp.CodeInternalServerError
	}

	// 如果策略详情已存在更新
	// 如果策略详情不存在创建
	// todo 如何查找策略 ？？？

	fp := &dao.FilePolicy{
		FileID:     af.FileID,
		CreatorID:  af.FileOwnerID,
		ConsumerID: af.ProposerID,
	}
	filePolicy, err := fp.Get()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			//
		} else {
			log.Logger().WithField("file policy", fp).WithField("error", err).Error("get file policy failed")
			return resp.CodeInternalServerError
		}
	}

	if filePolicy.ID == 0 {
		// 文件和策略的关系不存在，创建策略记录和文件策略关系记录
		p := &dao.Policy{
			Hrac:             policy.Hrac,
			PolicyLabelID:    file.PolicyLabelID,
			Creator:          af.FileOwner,
			CreatorID:        af.FileOwnerID,
			Consumer:         af.Proposer,
			ConsumerID:       af.ProposerID,
			EncryptedPK:      policy.EncryptedPK,
			EncryptedAddress: policy.EncryptedAddress,
			Gas:              policy.Gas,
			TxHash:           policy.TxHash,
			StartAt:          af.StartAt,
			EndAt:            af.FinishAt,
		}
		policyID, err := p.Create()
		if err != nil {
			log.Logger().WithField("policy", p).WithField("error", err).Error("create policy failed")
			return resp.CodeInternalServerError
		}

		fp := &dao.FilePolicy{
			FileID:     af.FileID,
			PolicyID:   policyID,
			CreatorID:  af.FileOwnerID,
			ConsumerID: af.ProposerID,
			CreatedAt:  time.Time{},
			UpdatedAt:  time.Time{},
			DeletedAt:  gorm.DeletedAt{},
		}
		if _, err := fp.Create(); err != nil {
			log.Logger().WithField("file policy", fp).WithField("error", err).Error("create file policy failed")
			return resp.CodeInternalServerError
		}

	} else {
		// 文件和策略的关系存在，更新策略记录和文件策略关系记录
		p := &dao.Policy{
			ID: filePolicy.PolicyID,
		}
		newPolicy := &dao.Policy{
			EndAt: af.FinishAt,
		}
		if err := p.Updates(newPolicy); err != nil {
			log.Logger().WithField("policy", p).WithField("new policy", newPolicy).WithField("error", err).Error("update policy failed")
			return resp.CodeInternalServerError
		}

		newFilePolicy := &dao.FilePolicy{
			EndAt: af.FinishAt,
		}
		if err := fp.Updates(newFilePolicy); err != nil {
			log.Logger().WithField("file policy", p).WithField("new file policy", newFilePolicy).WithField("error", err).Error("update file policy failed")
			return resp.CodeInternalServerError
		}
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

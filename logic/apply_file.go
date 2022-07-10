package logic

import (
	"errors"
	"github.com/NuLink-network/nulink-node/utils"
	"time"

	"gorm.io/gorm"

	"github.com/NuLink-network/nulink-node/dao"
	"github.com/NuLink-network/nulink-node/entity"
	"github.com/NuLink-network/nulink-node/resource/log"
	"github.com/NuLink-network/nulink-node/resp"
)

func ApplyFile(fileIDs []string, proposerID string, startAt, endAt int64) (code int) {
	// todo
	fp := &dao.ApplyFile{
		ProposerID: proposerID,
	}
	query := &dao.QueryExtra{
		Conditions: map[string]interface{}{
			"file_id in ?": fileIDs,
			"status != ?":  dao.ApplyStatusRejected,
			// todo 不判断结束时间
			//"finish_at >= ?":    time.Unix(endAt, 0),
		},
	}
	applyFileList, err := fp.FindAny(query, nil)
	if err != nil {
		log.Logger().WithField("file policy", fp).WithField("query", query).WithField("error", err).Error("find file policy list failed")
		return resp.CodeInternalServerError
	}

	ignoreFileIDs := make(map[string]struct{}, len(applyFileList))
	for _, fp := range applyFileList {
		ignoreFileIDs[fp.FileID] = struct{}{}
	}
	if len(ignoreFileIDs) == len(fileIDs) {
		return resp.CodeFileApplied
	}

	acc := &dao.Account{AccountID: proposerID}
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
	if len(files) == 0 {
		return resp.CodeFileNotExist
	}

	fileInfo := make(map[string]map[string]string, len(files))
	for _, file := range files {
		// 自己的文件不需要申请
		if file.OwnerID == proposerID {
			ignoreFileIDs[file.FileID] = struct{}{}
			continue
		}
		fileInfo[file.FileID] = map[string]string{
			"name":          file.Name,
			"owner":         file.Owner,
			"owner_id":      file.OwnerID,
			"owner_address": file.OwnerAddress,
		}
	}
	if len(fileInfo) == 0 {
		log.Logger().Debug("ignored all file apply")
		return resp.CodeSuccess
	}

	afs := make([]*dao.ApplyFile, 0, len(fileIDs))
	for _, fid := range fileIDs {
		_, ok := ignoreFileIDs[fid]
		if ok {
			continue
		}
		afs = append(afs, &dao.ApplyFile{
			FileID:           fid,
			FileName:         fileInfo[fid]["name"],
			Proposer:         account.Name,
			ProposerID:       proposerID,
			ProposerAddress:  account.EthereumAddr,
			FileOwner:        fileInfo[fid]["owner"],
			FileOwnerID:      fileInfo[fid]["owner_id"],
			FileOwnerAddress: fileInfo[fid]["owner_address"],
			StartAt:          time.Unix(startAt, 0),
			FinishAt:         time.Unix(endAt, 0),
		})
	}
	if err := dao.NewAppleFile().BatchCreate(afs); err != nil {
		return resp.CodeInternalServerError
	}
	return resp.CodeSuccess
}

func ApplyFileList(fileID string, status uint8, proposerID, fileOwnerID string, page, pageSize int) ([]*entity.ApplyFileListResponse, int) {
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
		log.Logger().WithField("applyFile", af).WithField("error", err).Error("find apply file list failed")
		return nil, resp.CodeInternalServerError
	}
	if len(afs) == 0 {
		// 不存在申请信息
		return []*entity.ApplyFileListResponse{}, resp.CodeSuccess
	}

	fileIDs := make([]string, 0, len(afs))
	approvedFileIDs := make([]string, 0, 0)
	for _, af := range afs {
		fileIDs = append(fileIDs, af.FileID)
		if af.Status == dao.ApplyStatusApproved {
			approvedFileIDs = append(approvedFileIDs, af.FileID)
		}
	}

	query := &dao.QueryExtra{
		Conditions: map[string]interface{}{
			"file_id in ?": fileIDs,
		},
	}
	files, err := dao.NewFile().FindAny(query, nil)
	if err != nil {
		log.Logger().WithField("ext", query).WithField("error", err).Error("get file list failed")
		return nil, resp.CodeInternalServerError
	}
	fileID2Thumbnail := make(map[string]string, 0)
	for _, f := range files {
		fileID2Thumbnail[f.FileID] = f.Thumbnail
	}

	if len(approvedFileIDs) == 0 {
		// 所有申请未通过仅返回申请信息
		ret := make([]*entity.ApplyFileListResponse, 0, len(afs))
		for _, af := range afs {
			ret = append(ret, &entity.ApplyFileListResponse{
				FileID:           af.FileID,
				FileName:         af.FileName,
				Thumbnail:        fileID2Thumbnail[af.FileID],
				ApplyID:          af.ID,
				Proposer:         af.Proposer,
				ProposerID:       af.ProposerID,
				ProposerAddress:  af.ProposerAddress,
				FileOwner:        af.FileOwner,
				FileOwnerID:      af.FileOwnerID,
				FileOwnerAddress: af.FileOwnerAddress,
				Status:           af.Status,
				Remark:           af.Remark,
				StartAt:          af.StartAt.Unix(),
				EndAt:            af.FinishAt.Unix(),
				CreatedAt:        af.CreatedAt.Unix(),
			})
		}
		return ret, resp.CodeSuccess
	}

	fp := dao.FilePolicy{
		CreatorID:  fileOwnerID,
		ConsumerID: proposerID,
	}
	query = &dao.QueryExtra{
		Conditions: map[string]interface{}{
			"file_id in ?": approvedFileIDs,
		},
	}
	filePolicyList, err := fp.FindAny(query, nil)
	if err != nil {
		log.Logger().WithField("filePolicy", utils.JSON(fp)).WithField("ext", utils.JSON(query)).WithField("error", err).Error("find file policy list failed")
		return nil, resp.CodeInternalServerError
	}
	// 理论上审核过的申请一定会有文件策略关系记录
	// 如果未找到文件关系记录则是程序或数据异常

	policyIDs := make([]uint64, 0, 0)
	fileID2PolicyID := make(map[string]uint64, 0)
	for _, fp := range filePolicyList {
		policyIDs = append(policyIDs, fp.PolicyID)
		fileID2PolicyID[fp.FileID] = fp.PolicyID
	}

	query = &dao.QueryExtra{
		Conditions: map[string]interface{}{
			"id in ?": policyIDs,
		},
	}
	policies, err := dao.NewPolicy().Find(query, nil)
	if err != nil {
		log.Logger().WithField("ext", utils.JSON(query)).WithField("error", err).Error("find policy list failed")
		return nil, resp.CodeInternalServerError
	}
	policyID2Policy := make(map[uint64]*dao.Policy)
	for _, p := range policies {
		policyID2Policy[p.ID] = p
	}

	ret := make([]*entity.ApplyFileListResponse, 0, len(afs))
	for _, af := range afs {
		item := &entity.ApplyFileListResponse{
			FileID:           af.FileID,
			FileName:         af.FileName,
			Thumbnail:        fileID2Thumbnail[af.FileID],
			ApplyID:          af.ID,
			Proposer:         af.Proposer,
			ProposerID:       af.ProposerID,
			ProposerAddress:  af.ProposerAddress,
			FileOwner:        af.FileOwner,
			FileOwnerID:      af.FileOwnerID,
			FileOwnerAddress: af.FileOwnerAddress,
			Status:           af.Status,
			Remark:           af.Remark,
			StartAt:          af.StartAt.Unix(),
			EndAt:            af.FinishAt.Unix(),
			CreatedAt:        af.CreatedAt.Unix(),
		}
		if af.Status == dao.ApplyStatusApproved {
			if policy := policyID2Policy[fileID2PolicyID[af.FileID]]; policy != nil {
				item.PolicyID = policy.ID
				item.PolicyLabelID = policy.PolicyLabelID
				item.Hrac = policy.Hrac
			}
		}
		ret = append(ret, item)
	}
	return ret, resp.CodeSuccess
}

func RevokeApply(proposerID string, applyIDs []uint64) (code int) {
	af := &dao.ApplyFile{
		ProposerID: proposerID,
		Status:     dao.ApplyStatusApplying,
	}
	query := &dao.QueryExtra{
		Conditions: map[string]interface{}{
			"id in ?": applyIDs,
		},
	}
	applyFileList, err := af.FindAny(query, nil)
	if err != nil {
		log.Logger().WithField("applyFile", utils.JSON(af)).WithField("ext", utils.JSON(query)).WithField("error", err).Error("find apply file list failed")
		return resp.CodeInternalServerError
	}
	if len(applyFileList) == 0 {
		// todo
		// 不存在申请记录
		// 记录申请者不是当前用户
		// 申请不是待审核状态
		return resp.CodeSuccess
	}
	ids := make([]uint64, 0, len(applyFileList))
	for _, f := range applyFileList {
		ids = append(ids, f.ID)
	}

	if err := af.BatchDelete(ids); err != nil {
		return resp.CodeInternalServerError
	}

	return resp.CodeSuccess
}

func ApplyDetail(applyID uint64) (*entity.ApplyDetailResponse, int) {
	af := &dao.ApplyFile{
		ID: applyID,
	}
	apply, err := af.Get()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, resp.CodeApplyNotExist
		}
		log.Logger().WithField("applyID", applyID).WithField("error", err).Error("get apply file failed")
		return nil, resp.CodeInternalServerError
	}

	f := &dao.File{
		FileID: apply.FileID,
	}
	file, err := f.Get()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, resp.CodeFileNotExist
		}
		log.Logger().WithField("fileID", apply.FileID).WithField("error", err).Error("get file failed")
		return nil, resp.CodeInternalServerError
	}

	pl := &dao.PolicyLabel{
		PolicyLabelID: file.PolicyLabelID,
	}
	policyLabel, err := pl.Get()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, resp.CodePolicyLabelNotExist
		}
		log.Logger().WithField("policyLabelID", file.PolicyLabelID).WithField("error", err).Error("get policy label failed")
		return nil, resp.CodeInternalServerError
	}
	return &entity.ApplyDetailResponse{
		StartAt:       apply.StartAt.Unix(),
		EndAt:         apply.FinishAt.Unix(),
		PolicyLabel:   policyLabel.Label,
		PolicyLabelID: policyLabel.PolicyLabelID,
	}, resp.CodeSuccess
}

func ApproveApply(accountID string, applyID uint64, remark string, policy entity.Policy) (code int) {
	af := &dao.ApplyFile{
		ID:          applyID,
		FileOwnerID: accountID,
	}
	apply, err := af.Get()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// todo 提示无权操作
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

	fp := &dao.FilePolicy{
		FileID:     apply.FileID,
		CreatorID:  apply.FileOwnerID,
		ConsumerID: apply.ProposerID,
	}
	filePolicy, err := fp.Get()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Logger().WithField("file policy", fp).WithField("error", err).Error("get file policy failed")
		return resp.CodeInternalServerError
	}

	if filePolicy.ID == 0 {
		// 文件和策略的关系不存在，创建策略记录和文件策略关系记录
		f := &dao.File{
			FileID: apply.FileID,
		}
		file, err := f.Get()
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return resp.CodeFileNotExist
			}
			log.Logger().WithField("file", f).WithField("error", err).Error("get file failed")
			return resp.CodeInternalServerError
		}

		p := &dao.Policy{
			Hrac:             policy.Hrac,
			PolicyLabelID:    file.PolicyLabelID,
			Creator:          apply.FileOwner,
			CreatorID:        apply.FileOwnerID,
			CreatorAddress:   apply.FileOwnerAddress,
			Consumer:         apply.Proposer,
			ConsumerID:       apply.ProposerID,
			ConsumerAddress:  apply.ProposerAddress,
			EncryptedPK:      policy.EncryptedPK,
			EncryptedAddress: policy.EncryptedAddress,
			Gas:              policy.Gas,
			TxHash:           policy.TxHash,
			StartAt:          apply.StartAt,
			EndAt:            apply.FinishAt,
		}
		policyID, err := p.Create()
		if err != nil {
			log.Logger().WithField("policy", p).WithField("error", err).Error("create policy failed")
			return resp.CodeInternalServerError
		}

		fp := &dao.FilePolicy{
			FileID:     apply.FileID,
			PolicyID:   policyID,
			CreatorID:  apply.FileOwnerID,
			ConsumerID: apply.ProposerID,
			StartAt:    apply.StartAt,
			EndAt:      apply.FinishAt,
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
			EndAt: apply.FinishAt,
		}
		if err := p.Updates(newPolicy); err != nil {
			log.Logger().WithField("policy", p).WithField("new policy", newPolicy).WithField("error", err).Error("update policy failed")
			return resp.CodeInternalServerError
		}

		newFilePolicy := &dao.FilePolicy{
			EndAt: apply.FinishAt,
		}
		if err := fp.Updates(newFilePolicy); err != nil {
			log.Logger().WithField("file policy", p).WithField("new file policy", newFilePolicy).WithField("error", err).Error("update file policy failed")
			return resp.CodeInternalServerError
		}
	}

	newAf := &dao.ApplyFile{
		Status: dao.ApplyStatusApproved,
		Remark: remark,
	}
	if err := af.Updates(newAf); err != nil {
		return resp.CodeInternalServerError
	}
	return resp.CodeSuccess
}

func RejectApply(accountID string, applyID uint64, remark string) (code int) {
	af := &dao.ApplyFile{
		ID: applyID,
	}

	apply, err := af.Get()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp.CodeApplyNotExist
		}
		return resp.CodeInternalServerError
	}
	if apply.FileOwnerID != accountID {
		return resp.CodeUnauthorized
	}
	if apply.Status == dao.ApplyStatusApproved {
		return resp.CodeApplyApproved
	}
	if apply.Status == dao.ApplyStatusRejected {
		return resp.CodeSuccess
	}

	newAf := &dao.ApplyFile{
		Status: dao.ApplyStatusRejected,
		Remark: remark,
	}
	if err := af.Updates(newAf); err != nil {
		return resp.CodeInternalServerError
	}
	return resp.CodeSuccess
}

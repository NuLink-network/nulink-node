package logic

import (
	"errors"
	"github.com/NuLink-network/nulink-node/dao"
	"github.com/NuLink-network/nulink-node/entity"
	"github.com/NuLink-network/nulink-node/resource/log"
	"github.com/NuLink-network/nulink-node/resp"
	"gorm.io/gorm"
)

func RevokePolicy(accountID string, policyID uint64) (code int) {
	p := &dao.Policy{
		ID:        policyID,
		CreatorID: accountID,
	}
	policy, err := p.Get()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp.CodePolicyNotExist
		}
		log.Logger().WithField("policy", p).WithField("error", err).Error("get policy failed")
		return resp.CodeInternalServerError
	}
	if policy.CreatorID != accountID {
		return resp.CodePolicyNotYours
	}

	fp := &dao.FilePolicy{
		PolicyID: policyID,
	}
	filePolicy, err := fp.Get()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return resp.CodeInternalServerError
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 如果策略记录存在，那么文件和策略的对应关系也应该存在。如果不存在则是程序异常。
		log.Logger().WithField("filePolicy", fp).Error("policy exist but file policy does not exist")
	}
	if _, err := fp.Delete(); err != nil {
		return resp.CodeInternalServerError
	}

	af := &dao.ApplyFile{
		FileID: filePolicy.FileID,
	}
	if _, err := af.Delete(); err != nil {
		return resp.CodeInternalServerError
	}

	_, err = p.Delete()
	if err != nil {
		return resp.CodeInternalServerError
	}
	return resp.CodeSuccess
}

func PolicyList(policyID uint64, policyLabelID, creatorID, consumerID string, page, pageSize int) (ret []*entity.PolicyListResponse, count int64, codee int) {
	p := &dao.Policy{
		ID:            policyID,
		PolicyLabelID: policyLabelID,
		CreatorID:     creatorID,
		ConsumerID:    consumerID,
	}
	ps, count, err := p.Find(nil, dao.Paginate(page, pageSize))
	if err != nil {
		log.Logger().WithField("policy", p).WithField("error", err).Error("find policy failed")
		return nil, count, resp.CodeInternalServerError
	}
	if count == 0 || len(ps) == 0 {
		return []*entity.PolicyListResponse{}, 0, resp.CodeSuccess
	}

	accountIDs := make([]string, 0, len(ps))
	for _, p := range ps {
		accountIDs = append(accountIDs, p.CreatorID)
		accountIDs = append(accountIDs, p.ConsumerID)
	}
	accounts, err := dao.NewAccount().FindAccountByAccountIDs(accountIDs)
	if err != nil {
		log.Logger().WithField("accountIDs", accountIDs).WithField("error", err).Error("find account by account ids failed")
		return nil, 0, resp.CodeInternalServerError
	}
	if len(accounts) == 0 {
		return nil, 0, resp.CodeAccountNotExist
	}

	ret = make([]*entity.PolicyListResponse, 0, len(ps))
	for _, p := range ps {
		creator := accounts[p.CreatorID]
		if creator == nil {
			return nil, 0, resp.CodeAccountNotExist
		}
		consumer := accounts[p.ConsumerID]
		if consumer == nil {
			return nil, 0, resp.CodeAccountNotExist
		}
		ret = append(ret, &entity.PolicyListResponse{
			Hrac:            p.Hrac,
			PolicyID:        p.ID,
			Creator:         accounts[p.CreatorID].Name,
			CreatorID:       p.CreatorID,
			CreatorAddress:  accounts[p.CreatorID].EthereumAddr,
			Consumer:        accounts[p.ConsumerID].Name,
			ConsumerID:      p.ConsumerID,
			ConsumerAddress: accounts[p.ConsumerID].EthereumAddr,
			Gas:             p.Gas,
			TxHash:          p.TxHash,
			EncryptedPK:     p.EncryptedPK,
			StartAt:         p.StartAt.Unix(),
			EndAt:           p.EndAt.Unix(),
			CreatedAt:       p.CreatedAt.Unix(),
		})
	}
	return ret, count, resp.CodeSuccess
}

func FileDetailList(policyID uint64, page, pageSize int) (ret []*entity.FileDetailListResponse, count int64, code int) {
	p := &dao.Policy{
		ID: policyID,
	}
	policy, err := p.Get()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 0, resp.CodePolicyNotExist
		}
		log.Logger().WithField("id", policyID).WithField("error", err).Error("get policy failed")
		return nil, 0, resp.CodeInternalServerError
	}

	fp := &dao.FilePolicy{
		PolicyID: policyID,
	}
	filePolicyList, count, err := fp.FindAny(nil, dao.Paginate(page, pageSize))
	if err != nil {
		log.Logger().WithField("policyID", policyID).WithField("error", err).Error("get file policy list failed")
		return nil, 0, resp.CodeInternalServerError
	}
	if count == 0 || len(filePolicyList) == 0 {
		return []*entity.FileDetailListResponse{}, count, resp.CodeSuccess
	}

	filePolicyListLength := len(filePolicyList)
	fileIDs := make([]string, 0, filePolicyListLength)
	for _, fp := range filePolicyList {
		fileIDs = append(fileIDs, fp.FileID)
	}

	query := &dao.QueryExtra{
		Conditions: map[string]interface{}{
			"file_id in ?": fileIDs,
		},
	}
	files, _, err := dao.NewFile().FindAny(query, nil)
	if err != nil {
		log.Logger().WithField("query", query).WithField("error", err).Error("get file list failed")
		return nil, 0, resp.CodeInternalServerError
	}
	if len(files) == 0 {
		return nil, 0, resp.CodeFileNotExist
	}

	accountIDs := make([]string, 0, len(files))
	fileID2file := make(map[string]*dao.File, len(files))
	for _, f := range files {
		accountIDs = append(accountIDs, f.OwnerID)
		fileID2file[f.FileID] = f
	}

	accounts, err := dao.NewAccount().FindAccountByAccountIDs(accountIDs)
	if err != nil {
		log.Logger().WithField("accountIDs", accountIDs).WithField("error", err).Error("find account by account ids failed")
		return nil, 0, resp.CodeInternalServerError
	}
	if len(accounts) == 0 {
		return nil, 0, resp.CodeAccountNotExist
	}

	ret = make([]*entity.FileDetailListResponse, 0, filePolicyListLength)
	for _, fp := range filePolicyList {
		file := fileID2file[fp.FileID]
		if file == nil {
			return nil, 0, resp.CodeFileNotExist
		}
		ret = append(ret, &entity.FileDetailListResponse{
			FileID:        fp.FileID,
			FileName:      file.Name,
			Owner:         accounts[file.OwnerID].Name,
			OwnerID:       file.OwnerID,
			OwnerAvatar:   accounts[file.OwnerID].Avatar,
			Address:       file.Address,
			Thumbnail:     file.Thumbnail,
			CreatedAt:     file.CreatedAt.Unix(),
			PolicyID:      fp.PolicyID,
			PolicyHrac:    policy.Hrac,
			PolicyStartAt: policy.StartAt.Unix(),
			PolicyEndAt:   policy.EndAt.Unix(),
		})
	}
	return ret, count, resp.CodeSuccess
}

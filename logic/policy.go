package logic

import (
	"github.com/NuLink-network/nulink-node/dao"
	"github.com/NuLink-network/nulink-node/entity"
	"github.com/NuLink-network/nulink-node/resource/log"
	"github.com/NuLink-network/nulink-node/resp"
)

func RevokePolicy(accountID, policyID string) (code int) {
	policy := &dao.Policy{
		PolicyID:  policyID,
		CreatorID: accountID,
	}
	rows, err := policy.Delete()
	if err != nil {
		return resp.CodeInternalServerError
	}
	if rows == 0 {
		return resp.CodePolicyNotExist
	}

	fp := &dao.FilePolicy{
		PolicyID: policyID,
	}
	filePolicy, err := fp.Get()
	if err != nil {
		return resp.CodeInternalServerError
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

	return resp.CodeSuccess
}

func PolicyList(policyID, creatorID, consumerID string, status uint8, page, pageSize int) ([]*entity.PolicyListResponse, int) {
	p := &dao.Policy{
		PolicyID:   policyID,
		CreatorID:  creatorID,
		ConsumerID: consumerID,
	}
	if status != dao.PolicyStatusAll {
		p.Status = status
	}
	ps, err := p.Find(Paginate(page, pageSize))
	if err != nil {
		log.Logger().WithField("policy", p).WithField("error", err).Error("find policy failed")
		return nil, resp.CodeInternalServerError
	}

	ret := make([]*entity.PolicyListResponse, 0, len(ps))
	for _, p := range ps {
		ret = append(ret, &entity.PolicyListResponse{
			Hrac:       p.Hrac,
			Label:      p.Label,
			PolicyID:   p.PolicyID,
			Creator:    p.Creator,
			CreatorID:  p.CreatorID,
			Consumer:   p.Consumer,
			ConsumerID: p.ConsumerID,
			//EncryptedPK:      p.EncryptedPK,
			//EncryptedAddress: p.EncryptedAddress,
			Status:    p.Status,
			Gas:       p.Gas,
			TxHash:    p.TxHash,
			CreatedAt: p.CreatedAt,
		})
	}
	return ret, resp.CodeSuccess
}

func FileDetailList(creatorID, consumerID string, status uint8, page, pageSize int) ([]*entity.FileDetailListResponse, int) {
	p := &dao.Policy{
		CreatorID:  creatorID,
		ConsumerID: consumerID,
	}
	if status != dao.PolicyStatusAll {
		p.Status = status
	}
	policyIDs, err := p.FindPolicyIDs()
	if err != nil {
		log.Logger().WithField("policy", *p).WithField("error", err).Error("find policy ids failed")
		return nil, resp.CodeInternalServerError
	}

	fileIDs, err := dao.NewFilePolicy().FindFileIDsByPolicyIDs(policyIDs, Paginate(page, pageSize))
	if err != nil {
		log.Logger().WithField("policyIDs", policyIDs).WithField("error", err).Error("find file ids by policy ids failed")
		return nil, resp.CodeInternalServerError
	}

	files, err := dao.NewFile().FindByFileIDs(fileIDs, nil)
	if err != nil {
		log.Logger().WithField("fileIDs", fileIDs).WithField("error", err).Error("find file by file ids failed")
		return nil, resp.CodeInternalServerError
	}

	ret := make([]*entity.FileDetailListResponse, 0, 10)
	for _, f := range files {
		ret = append(ret, &entity.FileDetailListResponse{
			OwnerID:   f.OwnerID,
			FileName:  f.Name,
			Address:   f.Address,
			Thumbnail: f.Thumbnail,
			CreatedAt: f.CreatedAt.Unix(),
		})
	}
	return ret, resp.CodeSuccess
}

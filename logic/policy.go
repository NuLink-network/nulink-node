package logic

import (
	"github.com/NuLink-network/nulink-node/dao"
	"github.com/NuLink-network/nulink-node/entity"
	"github.com/NuLink-network/nulink-node/resource/log"
	"github.com/NuLink-network/nulink-node/resp"
)

func RevokePolicy(accountID string, policyID uint64) (code int) {
	policy := &dao.Policy{
		ID:        policyID,
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

func PolicyList(policyID uint64, creatorID, consumerID string, page, pageSize int) ([]*entity.PolicyListResponse, int) {
	p := &dao.Policy{
		ID:         policyID,
		CreatorID:  creatorID,
		ConsumerID: consumerID,
	}
	ps, err := p.Find(nil, dao.Paginate(page, pageSize))
	if err != nil {
		log.Logger().WithField("policy", p).WithField("error", err).Error("find policy failed")
		return nil, resp.CodeInternalServerError
	}

	ret := make([]*entity.PolicyListResponse, 0, len(ps))
	for _, p := range ps {
		ret = append(ret, &entity.PolicyListResponse{
			Hrac:       p.Hrac,
			PolicyID:   p.ID,
			Creator:    p.Creator,
			CreatorID:  p.CreatorID,
			Consumer:   p.Consumer,
			ConsumerID: p.ConsumerID,
			Gas:        p.Gas,
			TxHash:     p.TxHash,
			StartAt:    p.StartAt.Unix(),
			EndAt:      p.EndAt.Unix(),
			CreatedAt:  p.CreatedAt,
		})
	}
	return ret, resp.CodeSuccess
}

func FileDetailList(policyID uint64, creatorID, consumerID string, page, pageSize int) ([]*entity.FileDetailListResponse, int) {
	p := &dao.Policy{
		ID:         policyID,
		CreatorID:  creatorID,
		ConsumerID: consumerID,
	}
	policyIDs, err := p.FindPolicyIDs()
	if err != nil {
		log.Logger().WithField("policy", *p).WithField("error", err).Error("find policy ids failed")
		return nil, resp.CodeInternalServerError
	}

	fileIDs, err := dao.NewFilePolicy().FindFileIDsByPolicyIDs(policyIDs, dao.Paginate(page, pageSize))
	if err != nil {
		log.Logger().WithField("policy ids", policyIDs).WithField("error", err).Error("find file ids by policy ids failed")
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
			FileID:        f.FileID,
			FileName:      f.Name,
			Owner:         f.Owner,
			OwnerID:       f.OwnerID,
			Address:       f.Address,
			Thumbnail:     f.Thumbnail,
			CreatedAt:     f.CreatedAt.Unix(),
			PolicyID:      0,
			PolicyHrac:    "",
			PolicyStartAt: 0,
			PolicyEndAt:   0,
		})
	}
	return ret, resp.CodeSuccess
}

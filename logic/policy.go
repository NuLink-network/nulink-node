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

func PolicyList(policyID uint64, policyLabelID, creatorID, consumerID string, page, pageSize int) ([]*entity.PolicyListResponse, int) {
	p := &dao.Policy{
		ID:            policyID,
		PolicyLabelID: policyLabelID,
		CreatorID:     creatorID,
		ConsumerID:    consumerID,
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
			CreatedAt:  p.CreatedAt.Unix(),
		})
	}
	return ret, resp.CodeSuccess
}

func FileDetailList(policyID uint64, creatorID, consumerID string, page, pageSize int) ([]*entity.FileDetailListResponse, int) {
	fp := &dao.FilePolicy{
		PolicyID:   policyID,
		CreatorID:  creatorID,
		ConsumerID: consumerID,
	}
	query := &dao.QueryExtra{
		DistinctStr: []string{"file_id"},
	}
	filePolicyList, err := fp.FindAny(query, dao.Paginate(page, pageSize))
	if err != nil {
		log.Logger().WithField("filePolicy", fp).WithField("ext", query).WithField("error", err).Error("get file policy list failed")
		return nil, resp.CodeInternalServerError
	}

	filePolicyListLength := len(filePolicyList)
	fileIDs := make([]string, 0, filePolicyListLength)
	policyIDs := make([]uint64, 0, filePolicyListLength)
	for _, fp := range filePolicyList {
		fileIDs = append(fileIDs, fp.FileID)
		policyIDs = append(policyIDs, fp.PolicyID)
	}

	query = &dao.QueryExtra{
		Conditions: map[string]interface{}{
			"file_id in ?": fileIDs,
		},
	}
	files, err := dao.NewFile().FindAny(query, nil)
	if err != nil {
		log.Logger().WithField("ext", query).WithField("error", err).Error("get file list failed")
		return nil, resp.CodeInternalServerError
	}
	fileID2file := make(map[string]*dao.File, len(files))
	for _, f := range files {
		fileID2file[f.FileID] = f
	}

	query = &dao.QueryExtra{
		Conditions: map[string]interface{}{
			"file_id in ?": fileIDs,
		},
	}
	policies, err := dao.NewPolicy().Find(query, nil)
	if err != nil {
		log.Logger().WithField("ext", query).WithField("error", err).Error("get file list failed")
		return nil, resp.CodeInternalServerError
	}
	policyID2policy := make(map[uint64]*dao.Policy, len(files))
	for _, p := range policies {
		policyID2policy[p.ID] = p
	}

	ret := make([]*entity.FileDetailListResponse, 0, filePolicyListLength)
	for _, fp := range filePolicyList {
		file := fileID2file[fp.FileID]
		policy := policyID2policy[fp.PolicyID]
		ret = append(ret, &entity.FileDetailListResponse{
			FileID:        fp.FileID,
			FileName:      file.Name,
			Owner:         file.Owner,
			OwnerID:       file.OwnerID,
			Address:       file.Address,
			Thumbnail:     file.Thumbnail,
			CreatedAt:     file.CreatedAt.Unix(),
			PolicyID:      fp.PolicyID,
			PolicyHrac:    policy.Hrac,
			PolicyStartAt: policy.StartAt.Unix(),
			PolicyEndAt:   policy.EndAt.Unix(),
		})
	}
	return ret, resp.CodeSuccess
}

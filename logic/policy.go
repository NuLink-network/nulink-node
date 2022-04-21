package logic

import (
	"github.com/NuLink-network/nulink-node/dao"
	"github.com/NuLink-network/nulink-node/entity"
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
	return resp.CodeSuccess
}

func PolicyList(policyID, creatorID, consumerID string, status uint8, page, pageSize int) ([]*entity.PolicyListResponse, error) {
	p := &dao.Policy{
		PolicyID:   policyID,
		CreatorID:  creatorID,
		ConsumerID: consumerID,
	}
	if status != dao.PolicyStatusAll {
		p.Status = status
	}
	ps, err := p.Find(page, pageSize)
	if err != nil {
		return nil, err
	}

	ret := make([]*entity.PolicyListResponse, 0, len(ps))
	for _, p := range ps {
		ret = append(ret, &entity.PolicyListResponse{
			Hrac:             p.Hrac,
			Label:            p.Label,
			PolicyID:         p.PolicyID,
			Creator:          p.Creator,
			CreatorID:        p.CreatorID,
			Consumer:         p.Consumer,
			ConsumerID:       p.ConsumerID,
			EncryptedPK:      p.EncryptedPK,
			EncryptedAddress: p.EncryptedAddress,
			Status:           p.Status,
			Gas:              p.Gas,
			TxHash:           p.TxHash,
			CreatedAt:        p.CreatedAt,
		})
	}
	return ret, nil
}

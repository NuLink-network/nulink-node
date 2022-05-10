package logic

import (
	"github.com/NuLink-network/nulink-node/dao"
	"github.com/NuLink-network/nulink-node/entity"
	"github.com/NuLink-network/nulink-node/resource/log"
	"github.com/NuLink-network/nulink-node/resp"
)

func PolicyLabelList(creatorID string, page entity.Paginate) ([]*entity.PolicyLabelListResponse, int) {
	pl := &dao.PolicyLabel{
		CreatorID: creatorID,
	}
	policyLabelList, err := pl.Find(dao.Paginate(page.Page, page.PageSize))
	if err != nil {
		log.Logger().WithField("policyLabel", pl).WithField("error", err).Error("get policy label list failed")
		return nil, resp.CodeInternalServerError
	}

	ret := make([]*entity.PolicyLabelListResponse, 0, len(policyLabelList))
	for _, pl := range policyLabelList {
		ret = append(ret, &entity.PolicyLabelListResponse{
			Label:     pl.Label,
			LabelID:   pl.PolicyLabelID,
			Creator:   pl.Creator,
			CreatorID: pl.CreatorID,
			CreateAt:  pl.CreatedAt.Unix(),
		})
	}
	return ret, resp.CodeSuccess
}

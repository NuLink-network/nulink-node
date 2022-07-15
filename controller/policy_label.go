package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/NuLink-network/nulink-node/entity"
	"github.com/NuLink-network/nulink-node/logic"
	"github.com/NuLink-network/nulink-node/resp"
)

func PolicyLabelList(c *gin.Context) {
	req := &entity.PolicyLabelListRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		resp.ParameterErr(c)
		return
	}

	list, count, code := logic.PolicyLabelList(req.CreatorID, req.Paginate)
	if code != resp.CodeSuccess {
		resp.Error(c, code)
		return
	}
	resp.SuccessList(c, list, count)
}

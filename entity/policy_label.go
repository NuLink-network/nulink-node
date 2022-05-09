package entity

// ========================= request =========================

type PolicyLabelListRequest struct {
	CreatorID string   `json:"creator_id" binding:"required"`
	Paginate  Paginate `json:"paginate"`
}

// ========================= response =========================

type PolicyLabelListResponse struct {
	Label   string `json:"label"`
	LabelID string `json:"label_id"`
}

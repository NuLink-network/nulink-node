package entity

// ========================= request =========================

type File struct {
	ID       string `json:"id" binding:"required"`
	MD5      string `json:"md5" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Suffix   string `json:"suffix"`
	Category string `json:"category" binding:"required"`
	Address  string `json:"address" binding:"required"`
}

type UploadFileRequest struct {
	Files     []File `json:"files" binding:"required"`
	AccountID string `json:"account_id" binding:"required"`
	PolicyID  uint64 `json:"policy_id" binding:"required"`
	Signature string `json:"signature" binding:"required"`
}

type CreatePolicyAndUploadFileRequest struct {
	Files         []File `json:"files" binding:"required"`
	AccountID     string `json:"account_id" binding:"required"`
	PolicyLabelID string `json:"policy_label_id" binding:"required"`
	PolicyLabel   string `json:"policy_label" binding:"required"`
	EncryptedPK   string `json:"encrypted_pk" binding:"required"`
	Signature     string `json:"signature" binding:"required"`
}

type Paginate struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type GetFileListRequest struct {
	AccountID string   `json:"account_id"`
	FileName  string   `json:"file_name"`
	Paginate  Paginate `json:"paginate"`
}

type GetOthersFileListRequest struct {
	AccountID string   `json:"account_id" binding:"required"`
	FileName  string   `json:"file_name"`
	Category  string   `json:"category"`
	Format    string   `json:"format"`
	Desc      bool     `json:"desc"`
	Paginate  Paginate `json:"paginate"`
}

type DeleteFileRequest struct {
	FileIDs   []string `json:"file_ids" binding:"required"`
	AccountID string   `json:"account_id" binding:"required"`
	Signature string   `json:"signature"`
}

// ========================= response =========================

type GetFileListResponse struct {
	FileID    string `json:"file_id"`
	FileName  string `json:"file_name"`
	Address   string `json:"address"`
	Owner     string `json:"owner"`
	OwnerID   string `json:"owner_id"`
	Thumbnail string `json:"thumbnail"`
	CreatedAt int64  `json:"created_at"`
}

type GetOthersFileListResponse struct {
	FileID    string `json:"file_id"`
	FileName  string `json:"file_name"`
	Address   string `json:"address"`
	Owner     string `json:"owner"`
	OwnerID   string `json:"owner_id"`
	Thumbnail string `json:"thumbnail"`
	CreatedAt int64  `json:"created_at"`
}

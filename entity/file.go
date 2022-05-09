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

type FileDetailRequest struct {
	FileID     string `json:"file_id" binding:"required"`
	ConsumerID string `json:"consumer_id" binding:"required"`
}

type FileDownloadInfoRequest struct {
	FileID string `json:"file_id" binding:"required"`
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

type FileDetailResponse struct {
	FileID          string `json:"file_id"`
	FileName        string `json:"file_name"`
	Address         string `json:"address"`
	Thumbnail       string `json:"thumbnail"`
	FileCreatedAt   int64  `json:"file_created_at"`
	PolicyID        uint64 `json:"policy_id"`
	Hrac            string `json:"hrac"`
	Label           string `json:"label"`
	Creator         string `json:"creator"`
	CreatorID       string `json:"creator_id"`
	Consumer        string `json:"consumer"`
	ConsumerID      string `json:"consumer_id"`
	Gas             string `json:"gas"`
	TxHash          string `json:"tx_hash"`
	PolicyStartAt   int64  `json:"policy_start_at"`
	PolicyEndAt     int64  `json:"policy_end_at"`
	PolicyCreatedAt int64  `json:"policy_created_at"`
	ApplyID         uint64 `json:"apply_id"`
	ApplyStartAt    int64  `json:"apply_start_at"`
	ApplyEndAt      int64  `json:"apply_end_at"`
	ApplyCreatedAt  int64  `json:"apply_created_at"`
	VerifyPK        string `json:"verify_pk"`
}

type FileDownloadInfoResponse struct {
	PolicyEncryptedPK string `json:"policy_encrypted_pk"`
}

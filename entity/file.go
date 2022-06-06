package entity

// ========================= request =========================

type File struct {
	ID        string `json:"id" binding:"required"`
	MD5       string `json:"md5" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Suffix    string `json:"suffix"`
	Category  string `json:"category" binding:"required"`
	Address   string `json:"address" binding:"required"`
	Thumbnail string `json:"thumbnail"`
}

type UploadFileRequest struct {
	Files         []File `json:"files" binding:"required"`
	AccountID     string `json:"account_id" binding:"required"`
	PolicyID      uint64 `json:"policy_id"`
	PolicyLabelID string `json:"policy_label_id"`
	Signature     string `json:"signature" binding:"required"`
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
	AccountID string   `json:"account_id" binding:"required"`
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
	// 文件信息
	FileID         string `json:"file_id"`
	FileName       string `json:"file_name"`
	Thumbnail      string `json:"thumbnail"`
	Creator        string `json:"creator"`
	CreatorID      string `json:"creator_id"`
	CreatorAddress string `json:"creator_address"`
	FileCreatedAt  int64  `json:"file_created_at"`
	// 申请信息
	ApplyID         uint64 `json:"apply_id"`
	ProposerAddress string `json:"proposer_address"`
	Status          uint8  `json:"status"`
	ApplyStartAt    int64  `json:"apply_start_at"`
	ApplyEndAt      int64  `json:"apply_end_at"`
	ApplyCreatedAt  int64  `json:"apply_created_at"`
	// 策略信息
	PolicyID        uint64 `json:"policy_id"`
	Hrac            string `json:"hrac"`
	Consumer        string `json:"consumer"`
	ConsumerID      string `json:"consumer_id"`
	Gas             string `json:"gas"`
	TxHash          string `json:"tx_hash"`
	PolicyCreatedAt int64  `json:"policy_created_at"`
	// 下载信息
	FileIPFSAddress        string `json:"file_ipfs_address"`
	PolicyEncryptedPK      string `json:"policy_encrypted_pk"`
	PolicyEncryptedAddress string `json:"encrypted_treasure_map_ipfs_address"`
	AliceVerifyPK          string `json:"alice_verify_pk"`
}

type FileDownloadInfoResponse struct {
	PolicyEncryptedPK string `json:"policy_encrypted_pk"`
}

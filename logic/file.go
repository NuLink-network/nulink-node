package logic

import (
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/NuLink-network/nulink-node/dao"
	"github.com/NuLink-network/nulink-node/entity"
	"github.com/NuLink-network/nulink-node/resource/log"
	"github.com/NuLink-network/nulink-node/resp"
	"github.com/NuLink-network/nulink-node/utils"
)

func DuplicateFilename(accountID string, fileNames []string) ([]string, int) {
	f := &dao.File{
		OwnerID: accountID,
	}
	query := &dao.QueryExtra{
		Conditions: map[string]interface{}{
			"name in ?": fileNames,
		},
	}
	fileList, err := f.FindAny(query, nil)
	if err != nil {
		log.Logger().WithField("file", utils.JSON(f)).WithField("ext", utils.JSON(query)).WithField("error", err).Error("get file list failed")
		return nil, resp.CodeInternalServerError
	}

	names := make([]string, 0)
	for _, f := range fileList {
		names = append(names, f.Name)
	}
	return names, resp.CodeSuccess
}

func UploadFile(accountID string, policyID uint64, files []entity.File) (code int) {
	p := &dao.Policy{ID: policyID}
	policy, err := p.Get()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp.CodePolicyNotExist
		}
		return resp.CodeInternalServerError
	}
	if policy.CreatorID != accountID {
		return resp.CodePolicyNotYours
	}

	fs := make([]*dao.File, 0, len(files))
	fps := make([]*dao.FilePolicy, 0, len(files))
	for _, f := range files {
		fs = append(fs, &dao.File{
			FileID:        f.ID,
			MD5:           f.MD5,
			Name:          f.Name,
			Suffix:        f.Suffix,
			Category:      f.Category,
			Address:       f.Address,
			Owner:         policy.Creator,
			OwnerID:       accountID,
			PolicyLabelID: policy.PolicyLabelID,
		})

		fps = append(fps, &dao.FilePolicy{
			FileID:     f.ID,
			PolicyID:   policyID,
			CreatorID:  accountID,
			ConsumerID: policy.ConsumerID,
			StartAt:    policy.StartAt,
			EndAt:      policy.EndAt,
		})
	}

	if err = dao.Tx(fs, fps); err != nil {
		log.Logger().WithField("file", utils.JSON(fs)).WithField("filePolicy", utils.JSON(fps)).WithField("error", err).Error("batch create file and file policy failed")
		return resp.CodeInternalServerError
	}
	return resp.CodeSuccess
}

func CreatePolicyAndUploadFile(accountID, policyLabelID, policyLabel, encryptedPK string, files []entity.File) (code int) {
	acc := &dao.Account{AccountID: accountID}
	account, err := acc.Get()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp.CodeAccountNotExist
		}
		log.Logger().WithField("account", *acc).WithField("error", err).Error("get account failed")
		return resp.CodeInternalServerError
	}

	pl := &dao.PolicyLabel{
		PolicyLabelID: policyLabelID,
		Label:         policyLabel,
		CreatorID:     accountID,
		EncryptedPK:   encryptedPK,
	}

	fs := make([]*dao.File, 0, len(files))
	for _, f := range files {
		fs = append(fs, &dao.File{
			FileID:        f.ID,
			MD5:           f.MD5,
			Name:          f.Name,
			Suffix:        f.Suffix,
			Category:      f.Category,
			Address:       f.Address,
			Owner:         account.Name,
			OwnerID:       accountID,
			PolicyLabelID: policyLabelID,
		})
	}
	if err = dao.Tx(pl, fs); err != nil {
		log.Logger().WithField("policyLabel", utils.JSON(pl)).WithField("files", utils.JSON(fs)).WithField("error", err).Error("create policy label and files failed")
		return resp.CodeInternalServerError
	}
	return resp.CodeSuccess
}

func GetFileList(accountID string, fileName string, page, pageSize int) (ret []*entity.GetFileListResponse, code int) {
	file := &dao.File{
		OwnerID: accountID,
	}

	query := &dao.QueryExtra{
		Conditions: map[string]interface{}{},
	}
	if !utils.IsEmpty(fileName) {
		query.Conditions["name like ?"] = "%" + fileName + "%"
	}

	files, err := file.FindAny(query, dao.Paginate(page, pageSize))
	if err != nil {
		log.Logger().WithField("file", file).WithField("error", err).Error("find files failed")
		return nil, resp.CodeInternalServerError
	}

	ret = make([]*entity.GetFileListResponse, 0, 10)
	for _, f := range files {
		ret = append(ret, &entity.GetFileListResponse{
			FileID:    f.FileID,
			FileName:  f.Name,
			Address:   f.Address,
			Thumbnail: f.Thumbnail,
			Owner:     f.Owner,
			OwnerID:   f.OwnerID,
			CreatedAt: f.CreatedAt.Unix(),
		})
	}
	return ret, resp.CodeSuccess
}

func GetOthersFileList(accountID string, fileName, category, format string, desc bool, page, pageSize int) (ret []*entity.GetOthersFileListResponse, code int) {
	file := &dao.File{
		Category: category,
	}
	conditions := map[string]interface{}{
		"owner_id != ?": accountID,
	}
	if !utils.IsEmpty(fileName) {
		conditions["name like ?"] = "%" + fileName + "%"
	}
	if format == utils.OtherFormat {
		conditions["suffix not in ?"] = utils.OtherFormatExcludeSuffix
	} else {
		s, ok := utils.FileFormat2Suffix[format]
		if ok {
			conditions["suffix in ?"] = s
		}
	}
	query := &dao.QueryExtra{
		Conditions: conditions,
	}
	if desc {
		query.OrderStr = "id desc"
	}
	files, err := file.FindAny(query, dao.Paginate(page, pageSize))
	if err != nil {
		log.Logger().WithField("file", file).WithField("error", err).Error("find others file failed")
		return nil, resp.CodeInternalServerError
	}

	ret = make([]*entity.GetOthersFileListResponse, 0, len(files))
	for _, f := range files {
		ret = append(ret, &entity.GetOthersFileListResponse{
			FileID:    f.FileID,
			FileName:  f.Name,
			Address:   f.Address,
			Thumbnail: f.Thumbnail,
			Owner:     f.Owner,
			OwnerID:   f.OwnerID,
			CreatedAt: f.CreatedAt.Unix(),
		})
	}
	return ret, resp.CodeSuccess
}

func DeleteFile(accountID string, fileIDs []string) (code int) {
	// todo signature verification
	applyFile := &dao.ApplyFile{
		FileOwnerID: accountID,
	}
	if err := applyFile.DeleteByFileIDs(fileIDs); err != nil {
		log.Logger().WithField("applyFile", applyFile).WithField("error", err).Error("delete apply file failed")
		return resp.CodeInternalServerError
	}

	filePolicy := &dao.FilePolicy{
		CreatorID: accountID,
	}
	if err := filePolicy.DeleteByFileIDs(fileIDs); err != nil {
		log.Logger().WithField("filePolicy", filePolicy).WithField("error", err).Error("delete file policy failed")
		return resp.CodeInternalServerError
	}

	file := &dao.File{
		OwnerID: accountID,
	}
	if err := file.DeleteByFilesIDs(fileIDs); err != nil {
		log.Logger().WithField("file", file).WithField("error", err).Error("delete file failed")
		return resp.CodeInternalServerError
	}
	return resp.CodeSuccess
}

func FileDetail(fileID, consumerID string) (ret *entity.FileDetailResponse, code int) {
	// 返回文件信息，策略信息，申请信息，文件拥有者 VerifyPK
	f := &dao.File{
		FileID: fileID,
	}
	file, err := f.Get()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, resp.CodeFileNotExist
		}
		log.Logger().WithField("file", f).WithField("error", err).Error("get file failed")
		return nil, resp.CodeInternalServerError
	}

	af := &dao.ApplyFile{
		FileID:     fileID,
		ProposerID: consumerID,
	}
	query := &dao.QueryExtra{
		Conditions: map[string]interface{}{
			"status != ?": dao.ApplyStatusRejected,
		},
	}
	applyFile, err := af.GetAny(query)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Logger().WithField("applyFile", af).WithField("error", err).Error("get apply file failed")
		return nil, resp.CodeInternalServerError
	}

	// 因为上传文件时关联已发布策略策略的使用者自动获得该文件的使用
	// 这种情况不存在申请记录表只存在文件策略关联纪律
	fp := &dao.FilePolicy{
		FileID:     fileID,
		ConsumerID: consumerID,
	}
	filePolicy, err := fp.Get()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, resp.CodeApplyUnApproved
		}
		log.Logger().WithField("filePolicy", fp).WithField("error", err).Error("get file policy failed")
		return nil, resp.CodeInternalServerError
	}

	// 0 0
	// 申请记录和文件策略关联纪律都不存在，直接返回文件信息
	if applyFile.ID == 0 && filePolicy.ID == 0 {
		return &entity.FileDetailResponse{
			FileID:        file.FileID,
			FileName:      file.Name,
			Thumbnail:     file.Thumbnail,
			Creator:       file.Owner,
			CreatorID:     file.OwnerID,
			FileCreatedAt: file.CreatedAt.Unix(),
		}, resp.CodeSuccess
	}

	// 1 0
	// 申请未通过(表示申请记录存在，文件策略关联记录还未创建)，返回文件信息和申请信息
	if applyFile.Status != dao.ApplyStatusApproved {
		return &entity.FileDetailResponse{
			FileID:         file.FileID,
			FileName:       file.Name,
			Thumbnail:      file.Thumbnail,
			Creator:        file.Owner,
			CreatorID:      file.OwnerID,
			FileCreatedAt:  file.CreatedAt.Unix(),
			ApplyID:        applyFile.ID,
			Status:         applyFile.Status,
			ApplyStartAt:   applyFile.StartAt.Unix(),
			ApplyEndAt:     applyFile.FinishAt.Unix(),
			ApplyCreatedAt: applyFile.CreatedAt.Unix(),
		}, resp.CodeSuccess
	}

	p := &dao.Policy{
		ID: filePolicy.PolicyID,
	}
	policy, err := p.Get()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, resp.CodeApplyUnApproved
		}
		log.Logger().WithField("policy", p).WithField("error", err).Error("get policy failed")
		return nil, resp.CodeInternalServerError
	}

	// 0 1
	// 申请记录不存在文件策略记录存在表示使用者是自动获取的文件使用权限
	// 返回文件信息，策略信息
	if applyFile.ID == 0 && filePolicy.ID == 1 {
		ret = &entity.FileDetailResponse{
			FileID:          file.FileID,
			FileName:        file.Name,
			Thumbnail:       file.Thumbnail,
			Creator:         file.Owner,
			CreatorID:       file.OwnerID,
			FileCreatedAt:   file.CreatedAt.Unix(),
			PolicyID:        policy.ID,
			Hrac:            policy.Hrac,
			Consumer:        policy.Consumer,
			ConsumerID:      policy.ConsumerID,
			Gas:             policy.Gas,
			TxHash:          policy.TxHash,
			PolicyCreatedAt: policy.CreatedAt.Unix(),
		}
	}

	// 1 1
	// 申请记录和文件策略关联记录都存在表示用户是通过申请获取到的文件使用权限
	// 返回文件信息，申请信息，策略信息
	if applyFile.ID == 1 && filePolicy.ID == 1 {
		ret = &entity.FileDetailResponse{
			FileID:          file.FileID,
			FileName:        file.Name,
			Thumbnail:       file.Thumbnail,
			Creator:         file.Owner,
			CreatorID:       file.OwnerID,
			FileCreatedAt:   file.CreatedAt.Unix(),
			ApplyID:         applyFile.ID,
			Status:          applyFile.Status,
			ApplyStartAt:    applyFile.StartAt.Unix(),
			ApplyEndAt:      applyFile.FinishAt.Unix(),
			ApplyCreatedAt:  applyFile.CreatedAt.Unix(),
			PolicyID:        policy.ID,
			Hrac:            policy.Hrac,
			Consumer:        policy.Consumer,
			ConsumerID:      policy.ConsumerID,
			Gas:             policy.Gas,
			TxHash:          policy.TxHash,
			PolicyCreatedAt: policy.CreatedAt.Unix(),
		}
	}

	// apply has expired
	if applyFile.FinishAt.Before(time.Now()) {
		return ret, resp.CodeSuccess
	}

	acc := &dao.Account{
		AccountID: file.OwnerID,
	}
	owner, err := acc.Get()
	if err != nil {
		log.Logger().WithField("policy", p).WithField("error", err).Error("get policy failed")
		return nil, resp.CodeInternalServerError
	}

	ret.FileIPFSAddress = file.Address
	ret.PolicyEncryptedPK = policy.EncryptedPK
	ret.PolicyEncryptedAddress = policy.EncryptedAddress
	ret.AliceVerifyPK = owner.VerifyPK
	return ret, resp.CodeSuccess
}

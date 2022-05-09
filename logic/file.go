package logic

import (
	"errors"
	"github.com/NuLink-network/nulink-node/dao"
	"github.com/NuLink-network/nulink-node/entity"
	"github.com/NuLink-network/nulink-node/resource/log"
	"github.com/NuLink-network/nulink-node/resp"
	"github.com/NuLink-network/nulink-node/utils"
	"gorm.io/gorm"
)

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
		return resp.CodeInternalServerError
	}
	return resp.CodeSuccess
}

func CreatePolicyAndUploadFile(accountID, policyLabelID, policyLabel, encryptedPK string, files []entity.File) (code int) {
	pl := &dao.PolicyLabel{PolicyLabelID: policyLabelID}
	isExist, err := pl.IsExist()
	if err != nil {
		log.Logger().WithField("policy label", pl).WithField("error", err).Error("get policy label failed")
		return resp.CodeInternalServerError
	}
	if isExist {
		return resp.CodePolicyIsExist
	}

	acc := &dao.Account{Account: accountID}
	account, err := acc.Get()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp.CodeAccountNotExist
		}
		log.Logger().WithField("account", acc).WithField("error", err).Error("get account failed")
		return resp.CodeInternalServerError
	}

	policy := &dao.PolicyLabel{
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
	if err = dao.Tx(policy, fs); err != nil {
		return resp.CodeInternalServerError
	}
	return resp.CodeSuccess
}

func GetFileList(accountID string, fileName string, page, pageSize int) (ret []*entity.GetFileListResponse, code int) {
	file := &dao.File{
		OwnerID: accountID,
		Name:    fileName,
	}

	query := &dao.QueryExtra{
		Conditions: map[string]interface{}{},
	}
	if utils.IsEmpty(fileName) {
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
	if utils.IsEmpty(fileName) {
		conditions["name like ?"] = "%" + fileName + "%"
	}
	if format == utils.OtherFormat {
		conditions["suffix not in ?"] = utils.OtherFormatExcludeSuffix
	} else {
		conditions["suffix in ?"] = utils.FileFormat2Suffix[format]
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

func FileDetail(fileID, consumerID string) (ret []*entity.FileDetailResponse, code int) {
	// 返回文件信息，策略信息，申请信息，文件拥有者 VerifyPK
	f := &dao.File{
		FileID: fileID,
	}
	file, err := f.Get()
	if err != nil {
		log.Logger().WithField("file", f).WithField("error", err).Error("get file failed")
		return nil, resp.CodeInternalServerError
	}

	fp := &dao.FilePolicy{
		FileID:     fileID,
		ConsumerID: consumerID,
	}
	filePolicy, err := fp.Get()
	if err != nil {
		log.Logger().WithField("filePolicy", fp).WithField("error", err).Error("get file policy failed")
		return nil, resp.CodeInternalServerError
	}

	p := &dao.Policy{
		ID: filePolicy.PolicyID,
	}
	policy, err := p.Get()
	if err != nil {
		log.Logger().WithField("policy", p).WithField("error", err).Error("get policy failed")
		return nil, resp.CodeInternalServerError
	}

	af := &dao.ApplyFile{
		FileID:     fileID,
		ProposerID: consumerID,
	}
	// todo
}

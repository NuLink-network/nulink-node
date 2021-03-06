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
	fileList, _, err := f.FindAny(query, nil)
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
			FileID:    f.ID,
			MD5:       f.MD5,
			Name:      f.Name,
			Suffix:    f.Suffix,
			Category:  f.Category,
			Address:   f.Address,
			Thumbnail: f.Thumbnail,
			//Owner:         policy.Creator,
			OwnerID: accountID,
			//OwnerAddress:  policy.CreatorAddress,
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
	// todo isExist
	_, err := acc.Get()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp.CodeAccountNotExist
		}
		log.Logger().WithField("account", utils.JSON(acc)).WithField("error", err).Error("get account failed")
		return resp.CodeInternalServerError
	}

	p := &dao.PolicyLabel{PolicyLabelID: policyLabelID}
	isExist, err := p.IsExist()
	if err != nil {
		log.Logger().WithField("policyLabel", utils.JSON(p)).WithField("error", err).Error("get policy label failed")
		return resp.CodeInternalServerError
	}
	if isExist {
		return resp.CodePolicyLabelIsExist
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
			FileID:    f.ID,
			MD5:       f.MD5,
			Name:      f.Name,
			Suffix:    f.Suffix,
			Category:  f.Category,
			Address:   f.Address,
			Thumbnail: f.Thumbnail,
			//Owner:         account.Name,
			OwnerID: accountID,
			//OwnerAddress:  account.EthereumAddr,
			PolicyLabelID: policyLabelID,
		})
	}
	if err = dao.Tx(pl, fs); err != nil {
		log.Logger().WithField("policyLabel", utils.JSON(pl)).WithField("files", utils.JSON(fs)).WithField("error", err).Error("create policy label and files failed")
		return resp.CodeInternalServerError
	}
	return resp.CodeSuccess
}

func GetFileList(accountID string, fileName string, page, pageSize int) (ret []*entity.GetFileListResponse, count int64, code int) {
	file := &dao.File{
		OwnerID: accountID,
	}

	query := &dao.QueryExtra{
		Conditions: map[string]interface{}{},
	}
	if !utils.IsEmpty(fileName) {
		query.Conditions["name like ?"] = "%" + fileName + "%"
	}

	files, count, err := file.FindAny(query, dao.Paginate(page, pageSize))
	if err != nil {
		log.Logger().WithField("file", file).WithField("error", err).Error("find files failed")
		return nil, count, resp.CodeInternalServerError
	}
	if count == 0 || len(files) == 0 {
		return []*entity.GetFileListResponse{}, count, resp.CodeSuccess
	}

	accountIDs := make([]string, 0, len(files))
	for _, f := range files {
		accountIDs = append(accountIDs, f.OwnerID)
	}
	accounts, err := dao.NewAccount().FindAccountByAccountIDs(accountIDs)
	if err != nil {
		log.Logger().WithField("accountIDs", accountIDs).WithField("error", err).Error("find account by account ids failed")
		return nil, 0, resp.CodeInternalServerError
	}
	if len(accounts) == 0 {
		return nil, 0, resp.CodeAccountNotExist
	}

	ret = make([]*entity.GetFileListResponse, 0, len(files))
	for _, f := range files {
		ret = append(ret, &entity.GetFileListResponse{
			FileID:    f.FileID,
			FileName:  f.Name,
			Address:   f.Address,
			Thumbnail: f.Thumbnail,
			Owner:     accounts[f.OwnerID].Name,
			OwnerID:   f.OwnerID,
			CreatedAt: f.CreatedAt.Unix(),
		})
	}
	return ret, count, resp.CodeSuccess
}

func GetOthersFileList(accountID string, include bool, fileName, category, format string, desc bool, page, pageSize int) (ret []*entity.GetOthersFileListResponse, count int64, code int) {
	file := &dao.File{
		Category: category,
	}

	conditions := make(map[string]interface{})
	if !include {
		conditions["owner_id != ?"] = accountID
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
	files, count, err := file.FindAny(query, dao.Paginate(page, pageSize))
	if err != nil {
		log.Logger().WithField("file", file).WithField("error", err).Error("find others file failed")
		return nil, count, resp.CodeInternalServerError
	}
	if count == 0 || len(files) == 0 {
		return []*entity.GetOthersFileListResponse{}, count, resp.CodeSuccess
	}

	accountIDs := make([]string, 0, len(files))
	for _, f := range files {
		accountIDs = append(accountIDs, f.OwnerID)
	}
	accounts, err := dao.NewAccount().FindAccountByAccountIDs(accountIDs)
	if err != nil {
		log.Logger().WithField("accountIDs", accountIDs).WithField("error", err).Error("find account by account ids failed")
		return nil, 0, resp.CodeInternalServerError
	}
	if len(accounts) == 0 {
		return nil, 0, resp.CodeAccountNotExist
	}

	if !include {
		ret = make([]*entity.GetOthersFileListResponse, 0, len(files))
		for _, f := range files {
			ret = append(ret, &entity.GetOthersFileListResponse{
				FileID:      f.FileID,
				FileName:    f.Name,
				Address:     f.Address,
				Thumbnail:   f.Thumbnail,
				Owner:       accounts[f.OwnerID].Name,
				OwnerID:     f.OwnerID,
				OwnerAvatar: accounts[f.OwnerID].Avatar,
				CreatedAt:   f.CreatedAt.Unix(),
			})
		}
		return ret, count, resp.CodeSuccess
	}

	r1 := make([]*entity.GetOthersFileListResponse, 0)
	r2 := make([]*entity.GetOthersFileListResponse, 0)
	for _, f := range files {
		r := &entity.GetOthersFileListResponse{
			FileID:      f.FileID,
			FileName:    f.Name,
			Address:     f.Address,
			Thumbnail:   f.Thumbnail,
			Owner:       accounts[f.OwnerID].Name,
			OwnerID:     f.OwnerID,
			OwnerAvatar: accounts[f.OwnerID].Avatar,
			CreatedAt:   f.CreatedAt.Unix(),
		}
		if f.OwnerID == accountID {
			r1 = append(r1, r)
			continue
		}
		r2 = append(r2, r)
	}
	ret = append(r1, r2...)
	return ret, count, resp.CodeSuccess
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
	// ApplyFile.ProposerID == Policy.ConsumerID == FilePolicy.ConsumerID
	acc := &dao.Account{AccountID: consumerID}
	proposer, err := acc.Get()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, resp.CodeAccountNotExist
		}
		log.Logger().WithField("accountID", consumerID).WithField("error", err).Error("get account failed")
		return nil, resp.CodeInternalServerError
	}

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

	acc = &dao.Account{AccountID: file.OwnerID}
	owner, err := acc.Get()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, resp.CodeAccountNotExist
		}
		log.Logger().WithField("accountID", file.OwnerID).WithField("error", err).Error("get account failed")
		return nil, resp.CodeInternalServerError
	}

	// ???????????????????????????????????????????????????????????????
	if file.OwnerID == consumerID {
		policyLabel := &dao.PolicyLabel{
			PolicyLabelID: file.PolicyLabelID,
		}
		pl, err := policyLabel.Get()
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, resp.CodePolicyLabelNotExist
			}
			log.Logger().WithField("policyLabelID", utils.JSON(policyLabel)).WithField("error", err).Error("get policy label failed")
			return nil, resp.CodeInternalServerError
		}
		return &entity.FileDetailResponse{
			FileID:            file.FileID,
			FileName:          file.Name,
			Thumbnail:         file.Thumbnail,
			Creator:           owner.Name,
			CreatorID:         file.OwnerID,
			CreatorAvatar:     owner.Avatar,
			CreatorAddress:    owner.EthereumAddr,
			FileIPFSAddress:   file.Address,
			FileCreatedAt:     file.CreatedAt.Unix(),
			PolicyEncryptedPK: pl.EncryptedPK,
			AliceVerifyPK:     owner.VerifyPK,
		}, resp.CodeSuccess
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

	// 1 0
	// ?????????????????????????????????????????????????????????????????????????????????????????????????????????
	if applyFile.ID != 0 && applyFile.Status != dao.ApplyStatusApproved {
		return &entity.FileDetailResponse{
			FileID:          file.FileID,
			FileName:        file.Name,
			Thumbnail:       file.Thumbnail,
			Creator:         owner.Name,
			CreatorID:       file.OwnerID,
			CreatorAvatar:   owner.Avatar,
			CreatorAddress:  owner.EthereumAddr,
			FileIPFSAddress: file.Address,
			FileCreatedAt:   file.CreatedAt.Unix(),
			ApplyID:         applyFile.ID,
			ProposerAddress: proposer.EthereumAddr,
			Status:          applyFile.Status,
			ApplyStartAt:    applyFile.StartAt.Unix(),
			ApplyEndAt:      applyFile.FinishAt.Unix(),
			ApplyCreatedAt:  applyFile.CreatedAt.Unix(),
		}, resp.CodeSuccess
	}

	// ??????????????????????????????????????????????????????????????????????????????????????????
	// ?????????????????????????????????????????????????????????????????????
	fp := &dao.FilePolicy{
		FileID:     fileID,
		ConsumerID: consumerID,
	}
	filePolicy, err := fp.Get()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Logger().WithField("filePolicy", fp).WithField("error", err).Error("get file policy failed")
		return nil, resp.CodeInternalServerError
	}

	// 0 0
	// ??????????????????????????????????????????????????????????????????????????????
	if applyFile.ID == 0 && filePolicy.ID == 0 {
		return &entity.FileDetailResponse{
			FileID:          file.FileID,
			FileName:        file.Name,
			Thumbnail:       file.Thumbnail,
			Creator:         owner.Name,
			CreatorID:       file.OwnerID,
			CreatorAvatar:   owner.Avatar,
			CreatorAddress:  owner.EthereumAddr,
			FileIPFSAddress: file.Address,
			FileCreatedAt:   file.CreatedAt.Unix(),
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
	// ???????????????????????????????????????????????????????????????????????????????????????????????????
	// ????????????????????????????????????????????????
	if applyFile.ID == 0 && filePolicy.ID != 0 {
		ret = &entity.FileDetailResponse{
			FileID:          file.FileID,
			FileName:        file.Name,
			Thumbnail:       file.Thumbnail,
			Creator:         owner.Avatar,
			CreatorID:       file.OwnerID,
			CreatorAvatar:   owner.Avatar,
			CreatorAddress:  owner.EthereumAddr,
			FileIPFSAddress: file.Address,
			FileCreatedAt:   file.CreatedAt.Unix(),
			ProposerAddress: proposer.EthereumAddr,
			Status:          dao.ApplyStatusApproved,
			ApplyStartAt:    filePolicy.StartAt.Unix(),
			ApplyEndAt:      filePolicy.EndAt.Unix(),
			ApplyCreatedAt:  filePolicy.CreatedAt.Unix(),
			PolicyID:        policy.ID,
			Hrac:            policy.Hrac,
			Consumer:        proposer.Name,
			ConsumerID:      policy.ConsumerID,
			Gas:             policy.Gas,
			TxHash:          policy.TxHash,
			PolicyCreatedAt: policy.CreatedAt.Unix(),
		}
	}

	// 1 1
	// ?????????????????????????????????????????????????????????????????????????????????????????????????????????
	// ????????????????????????????????????????????????
	if applyFile.ID != 0 && filePolicy.ID != 0 {
		ret = &entity.FileDetailResponse{
			FileID:          file.FileID,
			FileName:        file.Name,
			Thumbnail:       file.Thumbnail,
			Creator:         owner.Name,
			CreatorID:       file.OwnerID,
			CreatorAvatar:   owner.Avatar,
			CreatorAddress:  owner.EthereumAddr,
			FileIPFSAddress: file.Address,
			FileCreatedAt:   file.CreatedAt.Unix(),
			ApplyID:         applyFile.ID,
			ProposerAddress: proposer.EthereumAddr,
			Status:          applyFile.Status,
			ApplyStartAt:    applyFile.StartAt.Unix(),
			ApplyEndAt:      applyFile.FinishAt.Unix(),
			ApplyCreatedAt:  applyFile.CreatedAt.Unix(),
			PolicyID:        policy.ID,
			Hrac:            policy.Hrac,
			Consumer:        proposer.Name,
			ConsumerID:      policy.ConsumerID,
			Gas:             policy.Gas,
			TxHash:          policy.TxHash,
			PolicyCreatedAt: policy.CreatedAt.Unix(),
		}
	}

	// has expired
	if filePolicy.EndAt.Before(time.Now()) {
		return ret, resp.CodeSuccess
	}

	//ret.FileIPFSAddress = file.Address
	ret.PolicyEncryptedPK = policy.EncryptedPK
	ret.PolicyEncryptedAddress = policy.EncryptedAddress
	ret.AliceVerifyPK = owner.VerifyPK
	return ret, resp.CodeSuccess
}

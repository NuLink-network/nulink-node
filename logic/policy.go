package logic

import (
	"github.com/NuLink-network/nulink-node/dao"
)

func CreatePolicy(accountID uint64, label, encryptedPK, verifyPK, signature string) error {
	policy := &dao.Policy{
		AccountID:   accountID,
		Label:       label,
		EncryptedPK: encryptedPK,
		VerifyPK:    verifyPK,
		Signature:   signature,
	}
	_, err := policy.Create()
	if err != nil {
		return err
	}
	return nil
}

func RevokePolicy(accountID, policyID uint64, signature string) error {
	policy := &dao.Policy{ID: policyID}
	newPolicy := &dao.Policy{IsPublish: false}
	return policy.Updates(newPolicy)
}

package logic

import (
	"time"

	"github.com/NuLink-network/nulink-node/dao"
)

func ApplyFile(fileIDs []uint64, proposerID uint64, signature string, startAt, finishAt time.Time) error {
	af := &dao.AppleFile{}
	afs := make([]*dao.AppleFile, 0, len(fileIDs))
	for _, fid := range fileIDs {
		afs = append(afs, &dao.AppleFile{
			FileID:     fid,
			ProposerID: proposerID,
			Signature:  signature,
			StartAt:    startAt,
			FinishAt:   finishAt,
		})
	}
	return af.BatchCreate(afs)
}

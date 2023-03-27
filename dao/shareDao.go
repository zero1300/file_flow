package dao

import (
	"context"
	"file_flow/ent"
	"file_flow/ent/centralstoragepool"
	"file_flow/ent/share"
	"file_flow/ent/userstoragepool"
	"file_flow/global"
	"file_flow/models"
)

var shareDao *ShareDao

type ShareDao struct {
	db *ent.Client
}

func NewShareDao() *ShareDao {
	if shareDao == nil {
		shareDao = &ShareDao{
			db: global.Client,
		}
	}
	return shareDao
}

func (s ShareDao) CreateShare(userFileId, expiration int) (*ent.Share, error) {
	return s.db.Share.Create().SetUserFileID(userFileId).SetExpiration(expiration).SetClickNumber(0).Save(context.Background())
}

func (s ShareDao) GetShareInfo(shareId int) (*models.ShareInfo, error) {
	shareInfo := new(models.ShareInfo)

	sharePO, err := s.db.Share.Query().Where(share.ID(shareId)).First(context.Background())
	if err != nil {
		return nil, err
	}
	shareInfo.Id = sharePO.ID
	shareInfo.CreateAt = sharePO.CreateAt
	shareInfo.Expiration = sharePO.Expiration
	userFilePo, err := s.db.UserStoragePool.Query().Where(userstoragepool.ID(sharePO.UserFileID)).First(context.Background())
	if err != nil {
		return nil, err
	}
	shareInfo.Filename = userFilePo.Filename
	centralFilePo, err := s.db.CentralStoragePool.Query().Where(centralstoragepool.ID(userFilePo.RepoID)).First(context.Background())
	if err != nil {
		return nil, err
	}
	shareInfo.Size = centralFilePo.Size
	shareInfo.Path = centralFilePo.Path
	return shareInfo, nil
}

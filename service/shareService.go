package service

import (
	"errors"
	"file_flow/dao"
	"file_flow/models"
	"time"
)

var shareService *ShareService

type ShareService struct {
	shareDao *dao.ShareDao
	fileDao  *dao.FileDao
}

func NewShareService() *ShareService {
	if shareService == nil {
		shareService = &ShareService{
			shareDao: dao.NewShareDao(),
			fileDao:  dao.NewFileDao(),
		}
	}
	return shareService
}

func (s ShareService) CreateShare(userFileId, expiration, uid int) (int, error) {
	count, err := s.fileDao.CountUserFile(userFileId, uid)
	if err != nil || count == 0 {
		return -1, errors.New("创建文件共享失败: " + "权限不足")
	}
	share, err := s.shareDao.CreateShare(userFileId, expiration)
	if err != nil {
		return -1, errors.New("创建文件共享失败: " + "数据库异常")
	}
	return share.ID, nil
}

func (s ShareService) GetShareInfo(shareId int) (*models.ShareInfo, error) {
	info, err := s.shareDao.GetShareInfo(shareId)
	if err != nil {
		return nil, err
	}
	add := info.CreateAt.Add(time.Second * time.Duration(info.Expiration))
	if add.Before(time.Now()) {
		return nil, errors.New("文件已过期间")
	}
	return info, err
}

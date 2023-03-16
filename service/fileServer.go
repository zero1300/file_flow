package service

import (
	"crypto/md5"
	"errors"
	"file_flow/common/upload"
	"file_flow/dao"
	"file_flow/ent"
	"file_flow/models"
	"fmt"
	"github.com/google/uuid"
	"mime/multipart"
	"path"
)

var fileService *FileService

type FileService struct {
	fileDao *dao.FileDao
}

func NewFileService() *FileService {
	if fileService == nil {
		fileService = &FileService{fileDao: dao.NewFileDao()}
	}
	return fileService
}

func (f FileService) UploadFile(fileHeader *multipart.FileHeader, uid int) error {

	file, err := fileHeader.Open()
	if err != nil {
		return errors.New("文件异常: " + err.Error())
	}
	b := make([]byte, fileHeader.Size)
	_, err = file.Read(b)
	if err != nil {
		return errors.New("读取文件异常: " + err.Error())
	}
	hash := md5.Sum(b)
	_, err = f.fileDao.GetFileByHash(fmt.Sprintf("%x", hash))
	if !ent.IsNotFound(err) {
		return errors.New("文件已存在")
	}
	newUUID, err2 := uuid.NewUUID()
	if err2 != nil {
		return errors.New("uuid生成失败异常: " + err2.Error())
	}
	ext := path.Ext(fileHeader.Filename)

	objectName := newUUID.String() + ext
	info, err2 := upload.PutObject(objectName, fileHeader)
	if err2 != nil {
		return errors.New("上传文件失败: : " + err2.Error())
	}

	var one ent.CentralStoragePool
	one.Filename = objectName
	one.Path = info.Bucket + "/" + info.Key
	one.Ext = ext
	one.Hash = fmt.Sprintf("%x", hash)
	one.Size = float64(fileHeader.Size)
	onePo, err := f.fileDao.AddFile(one)
	if err != nil {
		return errors.New("上传文件失败: : " + err.Error())
	}

	f.fileDao.AddRelation(onePo, uid)
	return nil
}

func (f FileService) GetUserFile(uid, parentId int, p models.Paginate) ([]models.File, int, error) {
	return f.fileDao.GetUserFiles(uid, parentId, p)
}

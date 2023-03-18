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

func (f FileService) UploadFile(fileHeader *multipart.FileHeader, uid, parentId int) error {

	if !f.isDir(parentId) {
		return errors.New("上传文件失败: 非文件夹")
	}
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
	filePo, err := f.fileDao.GetFileByHash(fmt.Sprintf("%x", hash))
	var objectName string
	var onePo *ent.CentralStoragePool
	if ent.IsNotFound(err) {
		newUUID, err2 := uuid.NewUUID()
		if err2 != nil {
			return errors.New("uuid生成失败异常: " + err2.Error())
		}
		ext := path.Ext(fileHeader.Filename)

		objectName = newUUID.String() + ext
		info, err2 := upload.PutObject(objectName, fileHeader)
		if err2 != nil {
			return errors.New("上传文件失败: " + err2.Error())
		}
		unique, err := f.fileDao.CheckFilenameUnique(objectName, parentId, uid)
		if err != nil || unique != 0 {
			return errors.New("上传文件失败: " + "文件已存在")
		}

		var one ent.CentralStoragePool
		one.Filename = objectName
		one.Path = info.Bucket + "/" + info.Key
		one.Ext = ext
		one.Hash = fmt.Sprintf("%x", hash)
		one.Size = float64(fileHeader.Size)
		onePo, err = f.fileDao.AddFile(one)
		if err != nil {
			return errors.New("上传文件失败: : " + err.Error())
		}
	} else {
		unique, err := f.fileDao.CheckFilenameUnique(filePo.Filename, parentId, uid)
		if err != nil || unique != 0 {
			return errors.New("上传文件失败: " + "文件已存在")
		}
		onePo = filePo
	}
	f.fileDao.AddRelation(onePo, uid, parentId)
	return nil
}

func (f FileService) GetUserFile(uid, parentId int, p models.Paginate) ([]models.File, int, error) {
	return f.fileDao.GetUserFiles(uid, parentId, p)
}

func (f FileService) NewFolder(name string, parentId, uid int) error {
	if !f.isDir(parentId) {
		return errors.New("上传文件失败: 非文件夹")
	}
	unique, err := f.fileDao.CheckFilenameUnique(name, parentId, uid)
	if err != nil || unique != 0 {
		return errors.New("创建文件夹失败: " + err.Error())
	}
	_, err = f.fileDao.CreateFolder(name, parentId, uid)
	if err != nil {
		return errors.New("创建文件夹失败: " + err.Error())
	}
	return nil
}

func (f FileService) DelUserFile(uid, id int) error {
	count, err := f.fileDao.GetUserFile(id, uid)
	if err != nil || ent.IsNotFound(err) || count.Ext == "dir" {
		return errors.New("删除文件失败")
	}
	err = f.fileDao.DelFile(id)
	return nil
}

func (f FileService) MoveFile(id, uid, pid int) error {
	_, err := f.fileDao.GetUserFile(id, uid)
	if err != nil {
		return errors.New("移动文件失败")
	}
	if pid != 0 {
		userFile, err := f.fileDao.GetUserFile(pid, uid)
		if err != nil || userFile.Ext != "dir" {
			return errors.New("移动文件失败")
		}
		if !f.moveCheck(id, pid) {
			return errors.New("移动文件失败")
		}
	}

	_, err = f.fileDao.MoveFile(id, pid)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func (f FileService) isDir(id int) bool {
	if id == 0 {
		return true
	}
	file, err := f.fileDao.GetFileById(id)
	if err != nil {
		return false
	}
	return file.Ext == "dir"
}

func (f FileService) moveCheck(source, target int) bool {
	dir, err := f.fileDao.GetFileById(target)
	if err != nil {
		return false
	}
	if dir.ParentID == source {
		return false
	}
	var pid = dir.ParentID
	for pid != 0 {
		curDir, _ := f.fileDao.GetFileById(pid)
		if curDir.ParentID == source {
			return false
		}
		pid = curDir.ParentID
	}
	return true
}

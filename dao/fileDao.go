package dao

import (
	"context"
	"file_flow/dao/calc"
	"file_flow/ent"
	"file_flow/ent/centralstoragepool"
	"file_flow/ent/userstoragepool"
	"file_flow/global"
	"file_flow/models"
	"fmt"
)

var fileDao *FileDao

type FileDao struct {
	db *ent.Client
}

func NewFileDao() *FileDao {
	if fileDao == nil {
		fileDao = &FileDao{
			db: global.Client,
		}
	}
	return fileDao
}

func (f FileDao) GetFileByHash(hash string) (*ent.CentralStoragePool, error) {
	return f.db.CentralStoragePool.Query().Where(centralstoragepool.Hash(hash)).First(context.Background())
}

func (f FileDao) AddFile(file ent.CentralStoragePool) (*ent.CentralStoragePool, error) {
	return f.db.CentralStoragePool.Create().SetExt(file.Ext).SetFilename(file.Filename).SetHash(file.Hash).SetPath(file.Path).SetSize(file.Size).Save(context.Background())
}

func (f FileDao) AddRelation(file *ent.CentralStoragePool, uid, parentId int) {
	_, err := f.db.UserStoragePool.Create().SetExt(file.Ext).
		SetFilename(file.Filename).
		SetRepoID(file.ID).
		SetParentID(parentId).
		SetUID(uid).Save(context.Background())
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (f FileDao) GetUserFiles(uid int, parentId int, p models.Paginate) ([]models.File, int, error) {
	calc.Offset(&p)
	total := f.db.UserStoragePool.Query().Where(userstoragepool.UID(uid), userstoragepool.ParentID(parentId)).CountX(context.Background())
	files, err := f.db.UserStoragePool.Query().Where(userstoragepool.UID(uid), userstoragepool.ParentID(parentId)).Limit(p.PageSize).Offset(p.Page).Order(ent.Desc(userstoragepool.FieldCreateAt)).All(context.Background())
	filesModels := f.getFileInfoExpand(files)
	if err != nil {
		return nil, 0, err
	}
	return filesModels, total, nil
}

func (f FileDao) getFileInfoExpand(files []*ent.UserStoragePool) []models.File {
	filesModels := make([]models.File, 0)
	for i := 0; i < len(files); i++ {
		var filesModel models.File
		filesModel.ID = files[i].ID
		filesModel.Filename = files[i].Filename
		filesModel.CreateAt = files[i].CreateAt
		filesModel.Ext = files[i].Ext
		if files[i].RepoID != 0 {
			// RepoID 表示是文件夹
			f := f.db.CentralStoragePool.GetX(context.Background(), files[i].RepoID)
			filesModel.Path = f.Path
			filesModel.Size = f.Size
		}
		filesModels = append(filesModels, filesModel)
	}
	return filesModels
}

func (f FileDao) CreateFolder(name string, parentId, uid int) (*ent.UserStoragePool, error) {
	return f.db.UserStoragePool.Create().SetFilename(name).SetParentID(parentId).SetUID(uid).SetRepoID(0).SetExt("dir").Save(context.Background())
}

func (f FileDao) CheckFilenameUnique(name string, parentId, uid int) (int, error) {
	return f.db.UserStoragePool.Query().Where(userstoragepool.Filename(name), userstoragepool.ParentID(parentId), userstoragepool.UID(uid)).Count(context.Background())
}

func (f FileDao) DelFile(id int) error {
	return f.db.UserStoragePool.DeleteOneID(id).Exec(context.Background())
}

func (f FileDao) CountUserFile(id, uid int) (int, error) {
	return f.db.UserStoragePool.Query().Where(userstoragepool.UID(uid), userstoragepool.ID(id)).Count(context.Background())
}

func (f FileDao) GetUserFile(id, uid int) (*ent.UserStoragePool, error) {
	return f.db.UserStoragePool.Query().Where(userstoragepool.UID(uid), userstoragepool.ID(id)).First(context.Background())
}

func (f FileDao) GetFileById(id int) (*ent.UserStoragePool, error) {
	return f.db.UserStoragePool.Query().Where(userstoragepool.ID(id)).First(context.Background())
}

func (f FileDao) MoveFile(id, newPid int) (*ent.UserStoragePool, error) {
	return f.db.UserStoragePool.UpdateOneID(id).SetParentID(newPid).Save(context.Background())
}

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

func (f FileDao) AddRelation(file *ent.CentralStoragePool, uid int) {
	_, err := f.db.UserStoragePool.Create().SetExt(file.Ext).
		SetFilename(file.Filename).
		SetRepoID(file.ID).
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
		f := f.db.CentralStoragePool.GetX(context.Background(), files[i].RepoID)
		filesModel.Path = f.Path
		filesModel.Size = f.Size
		filesModels = append(filesModels, filesModel)
	}
	return filesModels
}

package dao

import (
	"context"
	"file_flow/ent"
	"file_flow/ent/centralstoragepool"
	"file_flow/global"
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

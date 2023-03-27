package models

import "file_flow/ent"

type ShareInfo struct {
	Id         int    `json:"id"`
	Filename   string `json:"filename"`
	Expiration int    `json:"expiration"`
	ent.CentralStoragePool
}

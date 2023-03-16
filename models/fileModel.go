package models

import "file_flow/ent"

type File struct {
	ent.UserStoragePool
	Path string  `json:"path"`
	Size float64 `json:"size"`
}

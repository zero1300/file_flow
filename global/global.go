package global

import (
	"file_flow/ent"
	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
)

var Client *ent.Client
var Redis *redis.Client
var Minio *minio.Client

package global

import (
	"file_flow/ent"
	"github.com/redis/go-redis/v9"
)

var Client *ent.Client
var Redis *redis.Client

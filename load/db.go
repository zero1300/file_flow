package load

import (
	"context"
	"file_flow/ent"
	_ "file_flow/ent/runtime"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
)

func InitDB() *ent.Client {
	client, err := ent.Open(viper.GetString("db.driverName"), viper.GetString("db.dsn"))
	if err != nil {
		panic("数据库连接失败: " + err.Error())
	}
	ctx := context.Background()
	// 运行自动迁移工具来创建所有Schema资源
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}

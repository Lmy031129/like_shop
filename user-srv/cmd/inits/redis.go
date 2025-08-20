package inits

import (
	"context"
	"fmt"
	"log"
	"user-srv/cmd/config"
	"user-srv/cmd/globar"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func InitRedis() {
	data := config.ConfigAppData.Redis
	addr := fmt.Sprintf("%s:%d", data.Host, data.Port)
	globar.Rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: data.Password, // no password set
		DB:       0,             // use default DB
	})

	err := globar.Rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
	log.Println("Redis is success")
}

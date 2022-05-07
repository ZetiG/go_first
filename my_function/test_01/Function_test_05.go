package test_01

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func init() {
	println("------------------------>>> Test 05 <<<------------------------")

	// connect
	rdb, ctx := redisConnect()

	// set
	redisOperator(rdb, ctx)

	// get
	redisGet(rdb, ctx)

	// close
	redisClose(rdb, ctx)
}

// 连接
func redisConnect() (*redis.Client, context.Context) {
	// 创建上下文，用于redis连接传递
	ctx := context.Background()

	// redis 连接
	rdb := redis.NewClient(&redis.Options{
		Addr:     "106.13.22.217:6379", // use default Addr
		Password: "",                   // no password set
		DB:       0,                    //use default DB
	})

	pong, err := rdb.Ping(ctx).Result()
	checkErr(err)
	fmt.Println(pong, err)

	return rdb, ctx
}

// set操作
func redisOperator(rdb *redis.Client, ctx context.Context) {
	err := rdb.Set(ctx, "key_123", "value_321", 0).Err()
	checkErr(err)

}

// get
func redisGet(rdb *redis.Client, ctx context.Context) {
	key := "key_123"
	val, err := rdb.Get(ctx, key).Result()
	checkErr(err)

	fmt.Println(key, val)
}

// close
func redisClose(rdb *redis.Client, ctx context.Context) {

	err := rdb.Close()
	checkErr(err)
}

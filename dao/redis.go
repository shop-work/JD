/*******
* @Author:qingmeng
* @Description:
* @File:redis
* @Date2022/3/25
 */

package dao

import (
	"fmt"
	"github.com/go-redis/redis"
)

//连接池

var RedisDB *redis.Client

func InitRedisClient() {
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := RedisDB.Ping().Result()
	if err != nil {
		fmt.Println("init redis err:", err)
		return
	}

}

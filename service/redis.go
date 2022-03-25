/*******
* @Author:qingmeng
* @Description:
* @File:redis
* @Date2022/3/25
 */

package service

import (
	"shop/dao"
	"time"
)

type RedisService struct {
}

func (redis *RedisService) RedisGetValue(key string) (string, error) {
	return dao.RedisDB.Get(key).Result()
}

func (redis *RedisService) RedisSetValue(key string, value string) error {
	return dao.RedisDB.Set(key, value, time.Minute*2).Err()
}

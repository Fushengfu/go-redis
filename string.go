package go_redis

import (
	"github.com/garyburd/redigo/redis"
	"log"
)

/**
 *  存储数据
 */
func (this *RedisClient) Set(key, str string, args ...int) bool {
	Rds := this.Conn.Get()
	defer Rds.Close()
	_, err := Rds.Do("Set", key, str)
	if err != nil {
		log.Println(err.Error())
		return false
	}

	expire := 3 * 3600

	if len(args) > 0 {
		expire = args[0]
	}

	if expire == 0 {
		return true
	}

	this.RedisExpire(key, expire)

	return true
}

/**
 *  存储数据
 */
func (this *RedisClient) Get(key string) string {
	Rds := this.Conn.Get()
	defer Rds.Close()
	res, err := redis.String(Rds.Do("GET", key))
	if err != nil {
		log.Println(err.Error())
		return ""
	}

	return res
}

/**
 *  存储数据判断
 */
func (this *RedisClient) SetNx(key, str string, args ...int) bool {
	Rds := this.Conn.Get()
	defer Rds.Close()
	_, err := Rds.Do("SETNX", key, str)
	if err != nil {
		log.Println(err.Error())
		return false
	}

	expire := 3 * 3600

	if len(args) > 0 {
		expire = args[0]
	}

	this.RedisExpire(key, expire)

	return true
}

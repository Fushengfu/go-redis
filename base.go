package go_redis

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"log"
	"strings"
)

type RedisClient struct {
	Conn *redis.Pool
	OK   string
}

/**
 *  获取客户端连接
 */
func NewRedisClient(conn *redis.Pool) *RedisClient {
	return &RedisClient{
		Conn: conn,
		OK:   "OK",
	}
}

/**
 *  获取redis数据
 */
func (this *RedisClient) Command(cmd string, keys ...interface{}) (str string, err error) {
	Rds := this.Conn.Get()
	defer Rds.Close()

	switch strings.ToUpper(cmd) {
	case "HMGET":
		res, err := redis.Strings(Rds.Do(cmd, keys...))
		if err == nil {
			return res[0], nil
		}
	case "HKEYS":
		res, err := redis.Strings(Rds.Do(cmd, keys...))
		if err == nil {
			if strs, er := json.Marshal(res); er == nil {
				return string(strs), nil
			}
		}
	case "GET":
		res, err := redis.String(Rds.Do(cmd, keys...))
		if err == nil {
			return res, nil
		}
	case "RPOP":
		res, err := redis.String(Rds.Do(cmd, keys...))
		if err == nil {
			return res, nil
		}

	case "RPOPLPUSH":
		res, err := redis.String(Rds.Do(cmd, keys...))
		if err == nil {
			return res, nil
		}

	case "LREM":
		res, err := redis.String(Rds.Do(cmd, keys...))
		if err == nil {
			return res, nil
		}

	default:

	}

	return str, err
}

/**
 *  获取redis数据
 */
func (this *RedisClient) Del(key string) (bool, error) {
	Rds := this.Conn.Get()
	defer Rds.Close()

	return redis.Bool(Rds.Do("DEL", key))
}

/**
 *  设置时效性
 */
func (this *RedisClient) RedisExpire(key string, expire int) {
	Rds := this.Conn.Get()
	defer Rds.Close()
	_, err := Rds.Do("expire", key, expire)
	if err != nil {
		log.Println(err.Error())
	}
}

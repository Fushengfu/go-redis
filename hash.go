package go_redis

import (
	"github.com/garyburd/redigo/redis"
	"log"
)

/**
 *  删除一个或多个哈希表字段
 */
func (this *RedisClient) Hdel(key string, agrs ...interface{}) int {
	Rds := this.Conn.Get()
	defer Rds.Close()
	agrs = append([]interface{}{key}, agrs...)
	res, err := redis.Int(Rds.Do("HDEL", agrs...))
	if err != nil {
		log.Println(err.Error())
		return 0
	}

	return res
}

/**
 *  查看哈希表 key 中，指定的字段是否存在。
 */
func (this *RedisClient) Hexists(key string, agrs ...interface{}) int {
	Rds := this.Conn.Get()
	defer Rds.Close()
	agrs = append([]interface{}{key}, agrs...)
	res, err := redis.Int(Rds.Do("HEXISTS", agrs...))
	if err != nil {
		log.Println(err.Error())
		return 0
	}

	return res
}

/**
 *  获取存储在哈希表中指定字段的值。
 */
func (this *RedisClient) Hget(key string, agrs ...interface{}) string {
	Rds := this.Conn.Get()
	defer Rds.Close()
	agrs = append([]interface{}{key}, agrs...)
	res, err := redis.String(Rds.Do("HGET", agrs...))
	if err != nil {
		log.Println(err.Error())
		return ""
	}

	return res
}

/**
 *  获取存储在哈希表中指定字段的值。
 */
func (this *RedisClient) Hgetall(key string) map[string]string {
	Rds := this.Conn.Get()
	defer Rds.Close()
	res, err := redis.StringMap(Rds.Do("HGETALL", key))
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return res
}

/**
 *  命令用于获取哈希表中的所有域。
 */
func (this *RedisClient) Hkeys(key string) []string {
	Rds := this.Conn.Get()
	defer Rds.Close()
	res, err := redis.Strings(Rds.Do("HKEYS", key))
	if err != nil {
		log.Println(err.Error())
		return []string{}
	}

	return res
}

/**
 *  用于获取哈希表中字段的数量。
 */
func (this *RedisClient) Hlen(key string) int {
	Rds := this.Conn.Get()
	defer Rds.Close()
	res, err := redis.Int(Rds.Do("HLEN", key))
	if err != nil {
		log.Println(err.Error())
		return 0
	}

	return res
}

/**
 *  获取所有给定字段的值。
 */
func (this *RedisClient) Hmget(key string, field interface{}) string {
	Rds := this.Conn.Get()
	defer Rds.Close()
	res, err := redis.Strings(Rds.Do("HMGET", key, field))
	if err != nil {
		log.Println(err.Error(), ":", field)
		return ""
	}

	return res[0]
}

/**
 *  此命令会覆盖哈希表中已存在的字段。
 */
func (this *RedisClient) Hmset(key string, field, value interface{}) bool {
	Rds := this.Conn.Get()
	defer Rds.Close()
	res, err := redis.String(Rds.Do("HMSET", key, field, value))
	if err != nil {
		log.Println(err.Error())
		return false
	}

	return res == this.OK
}

package go_redis

import (
	"github.com/garyburd/redigo/redis"
	"reflect"
)

/**
 *  删除一个或多个哈希表字段
 */
func (this *RedisClient) Hdel(key string, agrs ...interface{}) int {
	Rds := this.Conn.Get()
	defer Rds.Close()

	agrs = append([]interface{}{key}, agrs...)
	res, err := redis.Int(Rds.Do("HDEL", agrs...))
	if err == redis.ErrNil {
		//log.Println(err.Error())
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
	if err == redis.ErrNil {
		//log.Println(err.Error())
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
	if err == redis.ErrNil {
		//log.Println(err.Error())
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
	if err == redis.ErrNil {
		//log.Println(err.Error())
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
	if err == redis.ErrNil {
		//log.Println(err.Error())
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
	if err == redis.ErrNil {
		//log.Println(err.Error())
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
	if err == redis.ErrNil {
		//log.Println(err.Error(), ":", field)
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
	if err == redis.ErrNil {
		//log.Println(err.Error())
		return false
	}

	return res == this.OK
}

/**
 *  迭代哈希表中的键值对。
 */
func (this *RedisClient) Hscan(key string, cursor, match, count interface{}) (int, []interface{}) {
	Rds := this.Conn.Get()
	defer Rds.Close()

	if cursor == nil {
		cursor = 0
	}

	if count == nil {
		count = 10
	}

	if match == nil {
		match = "*"
	}

	res, err := redis.Values(Rds.Do("HSCAN", key, cursor, "MATCH", match, "COUNT", count))
	if err == redis.ErrNil {
		return 0, nil
	}

	items := []interface{}{}

	for _, item := range res {
		if reflect.TypeOf(item).String() == "[]interface {}" {
			items = item.([]interface{})
		} else {
			cursor = this.ToInt(item.([]byte))
		}
	}

	return this.ToInt(cursor), items
}

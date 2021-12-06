package go_redis

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"reflect"
)

/**
 *  堵塞读取数据
 */
func (this *RedisClient) Blpop(key string, timeout int) []string {
	Rds := this.Conn.Get()
	defer Rds.Close()
	res, err := redis.Strings(Rds.Do("BLPOP", key, timeout))
	if err != nil {
		log.Println(err.Error())
		return []string{}
	}

	return res
}

/**
 *  堵塞读取数据
 */
func (this *RedisClient) Brpop(key string, timeout int) []string {
	Rds := this.Conn.Get()
	defer Rds.Close()
	res, err := redis.Strings(Rds.Do("BRPOP", key, timeout))
	if err != nil {
		log.Println(err.Error())
		return []string{}
	}

	return res
}

/**
 *  堵塞读取数据
 */
func (this *RedisClient) BrPopLpush(key, key1 string, timeout int) string {
	Rds := this.Conn.Get()
	defer Rds.Close()
	res, err := redis.String(Rds.Do("BRPOPLPUSH", key, key1, timeout))
	if err == nil {
		return res
	}
	return ""
}

/**
 *  通过索引获取列表中的元素
 */
func (this *RedisClient) Lindex(key string, index int) string {
	Rds := this.Conn.Get()
	defer Rds.Close()
	res, err := redis.String(Rds.Do("LINDEX", key, index))
	if err == nil {
		return res
	}
	return ""
}

/**
 *  读取列表长度
 */
func (this *RedisClient) Llen(key string) int {
	Rds := this.Conn.Get()
	defer Rds.Close()
	res, err := redis.Int(Rds.Do("LLEN", key))
	if err == nil {
		return res
	}
	return 0
}

/**
 *  读取列表数据
 */
func (this *RedisClient) Lpop(key string) string {
	Rds := this.Conn.Get()
	defer Rds.Close()
	res, err := redis.String(Rds.Do("LPOP", key))
	if err != nil {
		log.Println(err.Error())
		return ""
	}

	return res
}

/**
 *  读取列表数据
 */
func (this *RedisClient) Rpop(key string) string {
	Rds := this.Conn.Get()
	defer Rds.Close()
	res, err := redis.String(Rds.Do("RPOP", key))
	if err != nil {
		log.Println(err.Error())
		return ""
	}

	return res
}

/**
 *  删除列表数据
 */
func (this *RedisClient) Lrem(key string, count int, value interface{}) int {
	Rds := this.Conn.Get()
	defer Rds.Close()
	res, err := redis.Int(Rds.Do("LREM", key, count, value))
	if err != nil {
		log.Println(err.Error())
		return 0
	}

	return res
}

/**
 *  移除列表的最后一个元素，并将该元素添加到另一个列表
 */
func (this *RedisClient) RpopLpush(key, key1 string) (string, error) {
	Rds := this.Conn.Get()
	defer Rds.Close()
	res, err := redis.String(Rds.Do("RPOPLPUSH", key, key1))
	return res, err
}

/**
 *  插入列表数据
 */
func (this *RedisClient) Lpush(key string, args ...interface{}) int {
	Rds := this.Conn.Get()
	defer Rds.Close()
	args = append([]interface{}{key}, args...)
	res, err := Rds.Do("lpush", args...)

	if err != nil {
		log.Println(err.Error())
		return 0
	}
	num := 0
	mType := reflect.TypeOf(res).String()
	if mType == "int64" {
		num = int(res.(int64))
	} else if mType == "int64" {
		num = int(res.(int32))
	} else {
		num = res.(int)
	}

	return num
}

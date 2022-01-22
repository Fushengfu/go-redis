package go_redis

import (
	"github.com/garyburd/redigo/redis"
	"log"
)

/**
 *  将一个或多个值插入到列表头部
 */
func (this *RedisClient) Lpush(key string, args ...interface{}) int {
	Rds := this.Conn.Get()
	defer Rds.Close()

	args = append([]interface{}{key}, args...)
	reply, err := Rds.Do("LPUSH", args...)

	if err != nil {
		log.Println(err.Error())
		return 0
	}

	return this.ToInt(reply)
}

/**
 *  将一个或多个值插入到列表尾部
 */
func (this *RedisClient) Rpush(key string, args ...interface{}) int {
	Rds := this.Conn.Get()
	defer Rds.Close()

	args = append([]interface{}{key}, args...)
	reply, err := Rds.Do("RPUSH", args...)

	if err != nil {
		log.Println(err.Error())
		return 0
	}

	return this.ToInt(reply)
}

/**
 *  读取列表数据
 */
func (this *RedisClient) Lpop(key string) interface{} {
	Rds := this.Conn.Get()
	defer Rds.Close()
	res, err := Rds.Do("LPOP", key)
	log.Println(res, err)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	if nil == res {
		return nil
	}

	return string(res.([]byte))
}

/**
 *  读取列表数据
 */
func (this *RedisClient) Rpop(key string) interface{} {
	Rds := this.Conn.Get()
	defer Rds.Close()
	reply, err := Rds.Do("RPOP", key)
	log.Println(reply, err)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	if nil == reply {
		return nil
	}

	return string(reply.([]byte))
}

/**
 *  堵塞读取数据
 */
func (this *RedisClient) Blpop(key string, timeout int) []interface{} {
	Rds := this.Conn.Get()
	defer Rds.Close()

	reply, err := Rds.Do("BLPOP", key, timeout)

	if err != nil {
		log.Println(err.Error())
		return nil
	}

	if nil == reply {
		return nil
	}

	var items []interface{}

	for _, item := range reply.([]interface{}) {
		items = append(items, string(item.([]byte)))
	}

	return items
}

/**
 *  堵塞读取数据
 */
func (this *RedisClient) Brpop(key string, timeout int) []interface{} {
	Rds := this.Conn.Get()
	defer Rds.Close()

	reply, err := Rds.Do("BRPOP", key, timeout)

	if err != nil {
		log.Println(err.Error())
		return nil
	}

	if nil == reply {
		return nil
	}

	var items []interface{}

	for _, item := range reply.([]interface{}) {
		items = append(items, string(item.([]byte)))
	}

	return items
}

/**
 *  堵塞读取数据
 */
func (this *RedisClient) BrPopLpush(key, key1 string, timeout int) interface{} {
	Rds := this.Conn.Get()
	defer Rds.Close()

	res, err := Rds.Do("BRPOPLPUSH", key, key1, timeout)

	if nil != err {
		log.Println(err.Error())
		return ""
	}

	return this.ToString(res)
}

/**
 *  通过索引获取列表中的元素
 */
func (this *RedisClient) Lindex(key string, index int) string {
	Rds := this.Conn.Get()
	defer Rds.Close()

	res, err := Rds.Do("LINDEX", key, index)
	log.Println(res, err)

	if nil != err {
		log.Println(err.Error())
		return ""
	}
	//
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

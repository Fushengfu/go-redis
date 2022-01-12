package go_redis

import (
	"github.com/garyburd/redigo/redis"
	"log"
)

/**
 *  存储键值对数据
 */
func (this *RedisClient) Set(key, value string, args ...int) bool {

	//判断不能传空值
	if "" == value {
		log.Println("value 值不能为空")
		return false
	}

	Rds := this.Conn.Get()
	defer Rds.Close()

	//插入键值对
	ok, err := Rds.Do("Set", key, value)

	//插入异常
	if nil != err {
		log.Println(err.Error())
		return false
	}

	//插入失败
	if "OK" != ok {
		return false
	}

	expire := 3 * 3600

	if len(args) > 0 {
		expire = args[0]
	}

	if expire == 0 {
		return true
	}

	//设置默认有效时间
	this.Expire(key, expire)

	return true
}

/**
 *  获取键值对数据
 */
func (this *RedisClient) Get(key string) string {
	Rds := this.Conn.Get()
	defer Rds.Close()

	//获取数据
	res, err := redis.String(Rds.Do("GET", key))
	if nil != err {
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
	result, err := redis.Int(Rds.Do("SETNX", key, str))
	if err != nil {
		log.Println("SETNX", key, err.Error())
		return false
	}

	expire := 3 * 3600

	if len(args) > 0 {
		expire = args[0]
	}

	if result == 0 {
		return false
	}

	this.Expire(key, expire)

	return true
}

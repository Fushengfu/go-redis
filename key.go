package go_redis

import (
	"github.com/garyburd/redigo/redis"
	"log"
)

/**
 *  删除键
 */
func (this *RedisClient) Del(key string) (num int) {
	Rds := this.Conn.Get()
	defer Rds.Close()

	effectNum, err := redis.Int(Rds.Do("DEL", key))
	if err == redis.ErrNil {
		//log.Println("DEL ERROR", err.Error())
		return 0
	}

	return effectNum
}

/**
 *  设置时效性
 */
func (this *RedisClient) Expire(key string, expire int) int {
	Rds := this.Conn.Get()
	defer Rds.Close()

	ret, err := redis.Int(Rds.Do("expire", key, expire))
	if err == redis.ErrNil {
		//log.Println(err.Error())
		return 0
	}

	return ret
}

/**
 *  获取键名对应的有效时间
 */
func (this *RedisClient) Ttl(keyName string) int {
	Rds := this.Conn.Get()
	defer Rds.Close()

	ret, err := Rds.Do("TTL", keyName)
	if err != nil {
		log.Println(err.Error())
		return 0
	}

	expire := this.ToInt(ret)

	if expire == -2 {
		return 0
	}

	return expire
}

/**
 *  返回 key 所储存的值的类型
 */
func (this *RedisClient) Type(keyName string) string {
	Rds := this.Conn.Get()
	defer Rds.Close()

	reply, err := redis.String(Rds.Do("TYPE", keyName))
	if err == redis.ErrNil {
		//log.Println(err.Error())
		return "none"
	}

	return reply
}

/**
 *  返回 key 所储存的值的类型
 */
func (this *RedisClient) Scan(cursor interface{}, pattern interface{}, count int) (int, []string) {
	Rds := this.Conn.Get()
	defer Rds.Close()

	reply, err := redis.Values(Rds.Do("SCAN", cursor, "MATCH", pattern, "COUNT", count))

	if err == redis.ErrNil {
		//log.Println(err.Error())
		return 0, []string{}
	}

	var cur int
	var items []string

	for k, v := range reply {
		if 0 == k {
			cur = this.ToInt(string(v.([]byte)))
		} else {
			for _, item := range v.([]interface{}) {
				items = append(items, string(item.([]byte)))
			}
		}
	}

	return cur, items
}

/**
 *  检查给定 key 是否存在
 */
func (this *RedisClient) Exists(keyName string) bool {
	Rds := this.Conn.Get()
	defer Rds.Close()

	reply, err := Rds.Do("EXISTS", keyName)
	if err != nil {
		log.Println(err.Error())
		return false
	}

	return this.ToInt(reply) > 0
}

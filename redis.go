package go_redis

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"strconv"
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
	client := new(RedisClient)
	client.OK = "OK"
	client.Conn = conn
	return client
}

/**
 *  转字符串
 */
func (this *RedisClient) ToString(data interface{}) string {
	if data == nil {
		return ""
	}

	switch data.(type) {
	case string:
		return data.(string)

	case int64:
		return strconv.FormatInt(data.(int64), 10)

	case int:
		return strconv.Itoa(data.(int))

	case float64:
		return strconv.Itoa(int(int64(data.(float64))))

	case float32:
		return strconv.Itoa(int(int32(data.(float32))))
	case []byte:
		return string(data.([]byte))
	default:
		return fmt.Sprintf("%v", data)
	}
}

/**
 *  转整型
 */
func (this *RedisClient) ToInt(data interface{}) int {
	if data == nil {
		return 0
	}

	switch data.(type) {
	case int:
		return data.(int)

	case int64:
		return int(data.(int64))

	case []byte:
		n, _ := strconv.ParseInt(string(data.([]byte)), 10, 0)
		return int(n)

	case float64:
		return int(int64(data.(float64)))

	case float32:
		return int(int32(data.(float32)))

	case string:
		va, _ := strconv.Atoi(data.(string))
		return va
	default:
		return 0
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

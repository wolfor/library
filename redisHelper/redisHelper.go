// redisHelper project redisHelper.go
package redisHelper

import (
	"github.com/go-redis/redis"
)

type RedisHelper struct {
	client *redis.Client
}

//redis访问帮助
//parameter :
//redisAddr 参数格式: 127.0.0.1:30000
func NewRedisHelper(redisAddr, password string, db int) *RedisHelper {
	this := new(RedisHelper)

	this.client = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: password, // no password set
		DB:       db,       // use default DB
	})

	return this
}

func (this *RedisHelper) Get(key string) string {
	val, err := this.client.Get(key).Result()

	if err == redis.Nil {
		return ""
	} else if err != nil {
		return ""
	} else {
		return val
	}
}

func (this *RedisHelper) Set(key string, val interface{}) bool {
	var isSuccess bool

	err := this.client.Set(key, val, 0).Err()
	if err != nil {
		isSuccess = false
	} else {
		isSuccess = true
	}

	return isSuccess
}

//channel of redis
/*********************************************************************************/

//向redis的channel发布消息
func (this *RedisHelper) Publish(channel string, message interface{}) bool {
	var isSuccess bool

	err := this.client.Publish(channel, message).Err()
	if err != nil {
		isSuccess = false
	} else {
		isSuccess = true
	}

	return isSuccess
}

//get data channel of redis
//监听channel发布的消息
func (this *RedisHelper) Subscribe(channel string) string {

	pubSub := this.client.Subscribe(channel)

	if pubSub == nil {
		return ""
	}

	msg, err := pubSub.ReceiveMessage()

	if err != nil {
		return ""
	}

	return msg.String()
}

//queue of redis
/*********************************************************************************/

//从队列的左侧入队一个或多个元素
func (this *RedisHelper) LPush(queue string, values ...interface{}) {
	this.client.LPush(queue, values...)
}

//从队列的右侧入队一个或多个元素
func (this *RedisHelper) RPush(queue string, values ...interface{}) {
	this.client.RPush(queue, values...)
}

//获取队列的元素数量
func (this *RedisHelper) LLen(queue string) int64 {
	result := this.client.LLen(queue)

	return result.Val()
}

//从队列中获取指定返回的元素
func (this *RedisHelper) LRange(queue string, start, stop int64) []string {
	//example : get all queue all item
	//	result := this.client.LRange(queue, 0, -1)

	result := this.client.LRange(queue, start, stop)

	return result.Val()
}

//从队列左侧出队一个元素
func (this *RedisHelper) LPop(queue string) {
	this.client.LPop(queue)
}

//从队列右侧出队一个元素
func (this *RedisHelper) RPop(queue string) {
	this.client.RPop(queue)
}

//修剪(trim)一个已存在的队列
//超过范围的下标并不会产生错误：如果 start 超过列表尾部，或者 start > end，结果会是列表变成空表（即该 key 会被移除）。
//如果 end 超过列表尾部，Redis 会将其当作列表的最后一个元素。
func (this *RedisHelper) LTrim(queue string, start, stop int64) {
	this.client.LTrim(queue, start, stop)
}

package gbootstrap

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisClient struct {
	client *redis.Client
}

// InitRedis 初始化Redis客户端
func InitRedis() (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:         cfg.Redis.URL,      // Redis地址
		Password:     cfg.Redis.Password, // Redis密码
		DB:           cfg.Redis.DB,       // Redis数据库
		DialTimeout:  5 * time.Second,    // 连接超时时间
		ReadTimeout:  3 * time.Second,    // 读取超时时间
		WriteTimeout: 3 * time.Second,    // 写入超时时间
		PoolSize:     10,                 // 连接池大小
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to redis: %w", err)
	}

	return &RedisClient{client: client}, nil
}

// Set 方法将键值对存储到Redis，并设置过期时间。
// 如果未指定过期时间（duration == 0），则默认设置为30天。
// key: 存储的键。
// value: 存储的值，可以是任意类型。
// duration: 键值对的过期时间。如果为0，则使用默认值30天。
func (r *RedisClient) Set(ctx context.Context, key string, value any, duration time.Duration) error {
	if duration == time.Duration(0) {
		duration = 30 * 24 * time.Hour // 设置默认过期时间为30天
	}
	return r.client.Set(ctx, key, value, duration).Err()
}

// Get 方法从Redis获取指定键的值。
// key: 要获取的键。
// 返回键对应的值和操作过程中可能出现的错误。
func (r *RedisClient) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

// Delete 方法从Redis中删除一个或多个键。
// keys: 要删除的键的列表。
// 返回操作过程中可能出现的错误。
func (r *RedisClient) Delete(ctx context.Context, keys ...string) error {
	return r.client.Del(ctx, keys...).Err()
}

// HSet 在Redis中为哈希表设置字段值。
// key: 哈希表的键名。
// field: 字段名。
// value: 字段值，可以是任意类型。
// 返回操作过程中可能出现的错误。
func (r *RedisClient) HSet(ctx context.Context, key, field string, value any) error {
	return r.client.HSet(ctx, key, field, value).Err()
}

// HGet 从Redis的哈希表中获取字段的值。
// key: 哈希表的键名。
// field: 字段名。
// 返回字段的值和操作过程中可能出现的错误。
func (r *RedisClient) HGet(ctx context.Context, key, field string) (string, error) {
	return r.client.HGet(ctx, key, field).Result()
}

// LPush 将一个或多个值插入到Redis列表的头部。
// key: 列表的键名。
// values: 要插入的值，可以是多个。
// 返回操作过程中可能出现的错误。
func (r *RedisClient) LPush(ctx context.Context, key string, values ...any) error {
	return r.client.LPush(ctx, key, values...).Err()
}

// RPop 移除并返回Redis列表的最后一个元素。
// key: 列表的键名。
// 返回列表的最后一个元素和操作过程中可能出现的错误。
func (r *RedisClient) RPop(ctx context.Context, key string) (string, error) {
	return r.client.RPop(ctx, key).Result()
}

// SAdd 将一个或多个成员元素加入到集合中。
// key: 集合的键名。
// members: 要添加的成员元素，可以是多个。
// 返回操作过程中可能出现的错误。
func (r *RedisClient) SAdd(ctx context.Context, key string, members ...any) error {
	return r.client.SAdd(ctx, key, members...).Err()
}

// SMembers 返回集合中的所有成员。
// key: 集合的键名。
// 返回集合中的所有成员和操作过程中可能出现的错误。
func (r *RedisClient) SMembers(ctx context.Context, key string) ([]string, error) {
	return r.client.SMembers(ctx, key).Result()
}

// Publish 向Redis发布消息。
// channel: 消息发布到的频道。
// message: 要发布的消息内容，可以是任意类型。
// 返回操作过程中可能出现的错误。
func (r *RedisClient) Publish(ctx context.Context, channel string, message any) error {
	return r.client.Publish(ctx, channel, message).Err()
}

// Subscribe 订阅Redis频道的消息。
// channel: 要订阅的频道。
// 返回一个PubSub对象，用于接收频道消息。
func (r *RedisClient) Subscribe(ctx context.Context, channel string) *redis.PubSub {
	return r.client.Subscribe(ctx, channel)
}

// ExecuteTransaction 执行一个Redis事务，传入的action包含需要在事务中执行的所有Redis命令。
// 如果action执行中出现错误，事务将被取消，否则提交事务。
func (r *RedisClient) ExecuteTransaction(ctx context.Context, action func(pipe redis.Pipeliner) error) error {
	// 使用TxPipeline开启一个事务
	_, err := r.client.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		return action(pipe)
	})

	if err != nil {
		// 处理错误，可以在这里记录日志或进行其他错误处理
		return fmt.Errorf("transaction failed: %w", err)
	}

	// 如果没有错误，事务自动提交
	return nil
}

// BatchSet 使用批量操作设置多个键值对。
// pairs: 键值对映射，键为string类型，值为any类型。
// 返回操作过程中可能出现的错误。
func (r *RedisClient) BatchSet(ctx context.Context, pairs map[string]any) error {
	_, err := r.client.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		for key, value := range pairs {
			pipe.Set(ctx, key, value, 0)
		}
		return nil
	})
	return err
}

// GetExpireDuration 获取过期时间
func (r *RedisClient) GetExpireDuration(ctx context.Context, key string) (time.Duration, error) {
	return r.client.TTL(ctx, key).Result()
}

//// 发布消息
//err := redisClient.Publish("myChannel", "Hello World!")
//if err != nil {
//    log.Fatalf("Failed to publish: %v", err)
//}
//
//// 订阅频道
//pubsub := redisClient.Subscribe("myChannel")
//// 获取订阅消息
//for msg := range pubsub.Channel() {
//    fmt.Println(msg.Channel, msg.Payload)
//}

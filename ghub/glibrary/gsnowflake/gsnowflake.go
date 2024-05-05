package gsnowflake

import (
	"github.com/bwmarrin/snowflake"
	"log"
	"sync"
	"time"
)

// SnowflakeGenerator struct to hold the snowflake node
type SnowflakeGenerator struct {
	node *snowflake.Node
}

var (
	generator *SnowflakeGenerator
	once      sync.Once
	lock      sync.Mutex
)

// GetInstance 创建并返回一个 SnowflakeGenerator 的单例
func GetInstance(nodeNumber int64) *SnowflakeGenerator {
	once.Do(func() {
		// 设置起始时间为2020年1月1日
		snowflake.Epoch = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano() / 1e6
		node, err := snowflake.NewNode(nodeNumber)
		if err != nil {
			log.Fatalf("Snowflake Node creation failed: %v", err)
		}
		generator = &SnowflakeGenerator{node: node}
	})
	return generator
}

// Generate 返回一个唯一的 snowflake ID
func (s *SnowflakeGenerator) Generate() int64 {
	lock.Lock()
	defer lock.Unlock()
	return s.node.Generate().Int64()
}

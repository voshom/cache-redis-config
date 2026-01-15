package helpers

import (
	"fmt"
	"strings"
)

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	Database int
}

func GetRedisConfig(host string, port int, password string, database int) *RedisConfig {
	return &RedisConfig{
		Host:     host,
		Port:     port,
		Password: password,
		Database: database,
	}
}

func (c *RedisConfig) String() string {
	return fmt.Sprintf("redis://%s:%d@%s:%d?db=%d",
		strings.TrimPrefix(c.Password, "redis://"),
		c.Port,
		c.Host,
		c.Database,
	)
}
package redisClient

import "github.com/garyburd/redigo/redis"

func Connect() redis.Conn {
	pool, _ := redis.Dial("tcp", "127.0.0.1:6379")
	return pool
}

func PoolConnect() redis.Conn {
	pool := &redis.Pool{
		MaxIdle:     1,
		MaxActive:   10,
		IdleTimeout: 180,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "127.0.0.1:6379")
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
	return pool.Get()
}

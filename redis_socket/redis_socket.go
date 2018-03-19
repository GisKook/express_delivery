package redis_socket

import (
	"github.com/garyburd/redigo/redis"
	"github.com/giskook/express_delivery/conf"
	"log"
	"time"
)

type RedisSocket struct {
	conf *conf.Redis
	pool *redis.Pool
}

func (r *RedisSocket) dial() (redis.Conn, error) {
	c, err := redis.Dial("tcp", r.conf.Addr)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	if len(r.conf.Passwd) > 0 {
		if _, err := c.Do("AUTH", r.conf.Passwd); err != nil {
			log.Println(err.Error())
			c.Close()
			return nil, err
		}
	}

	return c, err
}

func (r *RedisSocket) test_on_borrow(c redis.Conn, t time.Time) error {
	if time.Since(t) < time.Minute {
		return nil
	}

	_, err := c.Do("PING")

	return err
}

func NewRedisSocket(config *conf.Redis) (*RedisSocket, error) {
	return &RedisSocket{
		conf: config,
	}, nil
}

func (r *RedisSocket) InitPool() {
	r.pool = redis.NewPool(r.dial, r.conf.MaxIdle)
	r.pool.TestOnBorrow = r.test_on_borrow
}

func (socket *RedisSocket) GetConn() redis.Conn {
	return socket.pool.Get()
}

func (socket *RedisSocket) Close() {
	socket.pool.Close()
}

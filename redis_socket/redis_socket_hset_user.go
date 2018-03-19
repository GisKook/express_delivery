package redis_socket

import ()

const (
	// 余额
	HKEY_USER_BALANCE string = "balance"
)

func (r *RedisSocket) HsetUser(openid string, balance float32) error {
	conn := r.GetConn()
	defer conn.Close()
	_, err := conn.Do("HSET", openid, HKEY_USER_BALANCE, balance)

	return err
}

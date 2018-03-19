package redis_socket

import ()

const (
	HKEY_SESSION_KEY string = "session_key"
	HKEY_OPENID      string = "openid"
)

func (r *RedisSocket) HsetSession(session_id, session_key, openid string) error {
	conn := r.GetConn()
	defer conn.Close()
	_, err := conn.Do("HSET", session_id, HKEY_SESSION_KEY, session_key, HKEY_OPENID, openid)

	return err
}

func (r *RedisSocket) ExpireSession(session_id string, duration int) error {
	conn := r.GetConn()
	defer conn.Close()
	_, err := conn.Do("EXPIRE", session_id, duration)

	return err
}

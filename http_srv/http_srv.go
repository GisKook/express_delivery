package http_srv

import (
	"github.com/giskook/express_delivery/conf"
	"github.com/giskook/express_delivery/db"
	"github.com/giskook/express_delivery/redis_socket"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type HttpSrv struct {
	conf   *conf.Conf
	router *mux.Router
	db     *db.DBSocket
	redis  *redis_socket.RedisSocket
}

func NewHttpSrv(conf *conf.Conf) *HttpSrv {
	_db, err := db.NewDBSocket(conf.DB)
	if err != nil {
		log.Println(err.Error())

		return nil
	}

	redis, e := redis_socket.NewRedisSocket(conf.Redis)
	if e != nil {
		log.Println(e.Error())
		_db.Close()

		return nil
	}
	redis.InitPool()

	return &HttpSrv{
		conf:   conf,
		router: mux.NewRouter(),
		db:     _db,
		redis:  redis,
	}
}

func (h *HttpSrv) Start() {
	s := h.router.PathPrefix("/api").Subrouter()
	h.init_wechat_v1(s)

	h.router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(h.conf.Http.Assets))))

	if err := http.ListenAndServeTLS(h.conf.Http.Addr, h.conf.Http.CertFile, h.conf.Http.KeyFile, h.router); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func (h *HttpSrv) init_wechat_v1(r *mux.Router) {
	s := r.PathPrefix("/v1").Subrouter()
	s.HandleFunc("/session", h.handler_session)
	s.HandleFunc("/message", h.handler_message)
}

package http_srv

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/giskook/express_delivery/base"
	"log"
	"net/http"
)

const (
	WECHAT_USER_CODE string = "code"
	WECHAT_LOGIN_FMT string = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
)

func (h *HttpSrv) handler_session_post(r *http.Request) (string, error) {
	defer func() {
		if x := recover(); x != nil {
		}
	}()
	r.ParseForm()
	defer r.Body.Close()

	var req_code struct {
		Code string `json:"code"`
	}
	decoder := json.NewDecoder(r.Body)
	e := decoder.Decode(&req_code)
	if e != nil {
		return "", base.NewErr(e, base.ERR_HTTP_WECHAT_LOGIN_CODE, base.ERR_HTTP_WECHAT_LOGIN_DESC)
	}

	if req_code.Code == "" {
		return "", base.ERROR_HTTP_LACK_PARAMTERS
	}
	url := fmt.Sprintf(WECHAT_LOGIN_FMT, h.conf.App.AppID, h.conf.App.AppSecret, req_code.Code)

	resp, err := http.Get(url)
	if err != nil {
		return "", base.NewErr(err, base.ERR_HTTP_WECHAT_LOGIN_CODE, base.ERR_HTTP_WECHAT_LOGIN_DESC)
	}

	var result struct {
		Session_key string `json:"session_key,omitempty"`
		Openid      string `json:"openid,omitempty"`
		Errcode     int    `json:"errcode,omitempty"`
		Errmsg      string `json:"errmsg,omitempty"`
	}

	err = unmarshal_json_resp(resp, &result)
	if err != nil || result.Errcode != 0 {
		return "", base.NewErr(err, base.ERR_HTTP_WECHAT_LOGIN_CODE, base.ERR_HTTP_WECHAT_LOGIN_DESC)
	}

	session_id := make([]byte, 32)
	_, err = rand.Read(session_id)
	if err != nil {
		return "", base.NewErr(err, base.ERR_HTTP_WECHAT_GEN_SESSION_CODE, base.ERR_HTTP_WECHAT_GEN_SESSION_DESC)

	}
	session := base64.StdEncoding.EncodeToString(session_id)
	if h.redis.HsetSession(session, result.Session_key, result.Openid) != nil || h.redis.ExpireSession(session, h.conf.App.SessionLifeCycle) != nil {
		return "", base.NewErr(nil, base.ERR_HTTP_WECHAT_GEN_SESSION_CODE, base.ERR_HTTP_WECHAT_GEN_SESSION_DESC)

	}
	h.redis.HsetUser(result.Openid, 0.0)

	return session, nil

}

func (h *HttpSrv) handler_session(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if x := recover(); x != nil {
			EncodeErrResponse(w, base.ERROR_HTTP_INNER_PANIC)
		}
	}()

	var result struct {
		Session string `json:"session, omitempty"`
		Errcode int    `json:"errcode, omitempty"`
		Errmsg  string `json:"errmsg, omitempty"`
	}

	RecordReq(r)
	if r.Method == http.MethodPost {
		session, err := h.handler_session_post(r)
		if err != nil {
			result.Errcode = err.(*base.ExpressDeliveryError).Code
			result.Errmsg = err.(*base.ExpressDeliveryError).Describe
		} else {
			result.Session = session
		}

		log.Println(result)
		marshal_json_resp(w, result)
	}
}

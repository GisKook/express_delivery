package http_srv

import (
	"encoding/json"
	"fmt"
	"github.com/giskook/express_delivery/base"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
)

type GeneralResponse struct {
	Code int    `json:"errcode"`
	Desc string `json:"errmsg"`
}

var (
	GRS GeneralResponse = GeneralResponse{Code: base.ERR_NONE_CODE, Desc: base.ERR_NONE_DESC}
)

func EncodeErrResponse(w http.ResponseWriter, err *base.ExpressDeliveryError) {
	gr := &GeneralResponse{
		Code: err.Code,
		Desc: err.Desc(),
	}
	marshal_json_resp(w, gr)
}

func RecordReq(r *http.Request) {
	v, e := httputil.DumpRequest(r, true)
	if e != nil {
		log.Println(e.Error())
		return
	}
	log.Println(string(v))
}

// MarshalJson 把对象以json格式放到response中
func marshal_json_resp(w http.ResponseWriter, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	fmt.Fprint(w, string(data))
	return nil
}

// UnMarshalJson 从request中取出对象
func unmarshal_json_req(req *http.Request, v interface{}) error {
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	//json.Unmarshal([]byte(bytes.NewBuffer(result).String()), v)
	json.Unmarshal(result, v)
	return nil
}

func unmarshal_json_resp(resp *http.Response, v interface{}) error {
	result, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return err
	}

	log.Println(string(result))
	json.Unmarshal(result, v)

	return nil
}

package http_srv

import (
	"encoding/base64"
	"encoding/json"
	"github.com/giskook/express_delivery/base"
	"github.com/otiai10/gosseract"
	"io/ioutil"
	"log"
	"net/http"
)

const ()

func (h *HttpSrv) handler_message_post(r *http.Request) (string, error) {
	defer func() {
		if x := recover(); x != nil {
		}
	}()
	r.ParseForm()
	defer r.Body.Close()

	var req struct {
		ImageData string `json:"image_data"`
		Serial    int    `json:"serial"`
		Content   string `json:"content"`
	}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&req)

	if req.ImageData == "" ||
		req.Serial == 0 ||
		req.Content == "" {
		return "", base.ERROR_HTTP_LACK_PARAMTERS
	}
	image_data, er := base64.StdEncoding.DecodeString(req.ImageData)

	if er != nil {
		log.Println(er)
	}
	client := gosseract.NewClient()
	defer client.Close()
	client.SetWhitelist("0123456789")
	client.SetPageSegMode(gosseract.PSM_SINGLE_LINE)
	client.SetImageFromBytes(image_data[0:])
	text, _ := client.Text()
	log.Println(text)
	ioutil.WriteFile("./images/"+text+".jpg", image_data[0:], 0644)

	return text, nil
}

func (h *HttpSrv) handler_message(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if x := recover(); x != nil {
			EncodeErrResponse(w, base.ERROR_HTTP_INNER_PANIC)
		}
	}()

	var result struct {
		MsgID   string `json:"msgid, omitempty"`
		Errcode int    `json:"errcode, omitempty"`
		Errmsg  string `json:"errmsg, omitempty"`
	}

	//RecordReq(r)
	if r.Method == http.MethodPost {
		v, err := h.handler_message_post(r)
		if err != nil {
			result.Errcode = err.(*base.ExpressDeliveryError).Code
			result.Errmsg = err.(*base.ExpressDeliveryError).Describe
		}
		result.Errcode = 1
		result.Errmsg = v
		marshal_json_resp(w, result)
	}
}

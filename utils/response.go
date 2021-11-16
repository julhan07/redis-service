package utils

import (
	"encoding/json"
	"net/http"
)

type responseApi struct {
	Code    int          `json:"code"`
	Msg     string       `json:"msg"`
	Payload *interface{} `json:"payload,omitempty"`
}

func CreateResponseSuccess(w http.ResponseWriter, code int, pld interface{}) {

	data, _ := json.Marshal(&responseApi{
		Code:    code,
		Payload: &pld,
	})

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func CreateResponseError(w http.ResponseWriter, code int, msg error) {

	data, _ := json.Marshal(&responseApi{
		Code: code,
		Msg:  msg.Error(),
	})

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

}

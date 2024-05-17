package server

import (
	"errors"
	kerr "github.com/go-kratos/kratos/v2/errors"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	"net/http"
	myError "sastoj/pkg/error"
	"strings"
)

type Response struct {
	Success bool        `json:"success" form:"success"`
	Code    int         `json:"code" form:"code"`
	Message string      `json:"message" form:"message"`
	Data    interface{} `json:"data" form:"data"`
}

type ErrResponse struct {
	Success bool        `json:"success" form:"success"`
	Code    int         `json:"code" form:"code"`
	Message string      `json:"message" form:"message"`
	Data    interface{} `json:"data" form:"data"`
}

func ErrorEncoder(w http.ResponseWriter, r *http.Request, err error) {
	var kerror *kerr.Error
	var loerr *myError.LocalError
	if !errors.As(err, &kerror) {
		err = myError.ServerError
	} else {
		loerr = myError.FromError(err)
	}
	reply := &ErrResponse{}
	reply.Success = false
	reply.Code = loerr.Code
	reply.Data = nil
	reply.Message = loerr.Message

	codec, _ := khttp.CodecForRequest(r, "Accept")
	body, err := codec.Marshal(reply)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", strings.Join([]string{"application", codec.Name()}, "/"))
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func ResponseEncoder(w http.ResponseWriter, r *http.Request, v interface{}) error {
	reply := &Response{}
	reply.Success = true
	reply.Code = 0
	reply.Data = v
	reply.Message = "ok"

	codec, _ := khttp.CodecForRequest(r, "Accept")
	data, err := codec.Marshal(reply)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", strings.Join([]string{"application", codec.Name()}, "/"))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return nil
}

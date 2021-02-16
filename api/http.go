package api

import (
	"fmt"
	"project5/serializer/json"
	"project5/serializer/msgpack"
	"io/ioutil"
	"github.com/pkg/errors"
	"project5/shortener"
	"net/http"
	"github.com/go-chi/chi"
)


func NewHandler(service shortener.RedirectSerivce) RedirectHandler {
	return &handler{service}
}

type RedirectHandler interface {
	GET(http.ResponseWriter, *http.Request)
	POST(http.ResponseWriter, *http.Request)
}

type handler struct {
	redirectService shortener.RedirectSerivce
}

func (h handler) GET(res http.ResponseWriter, req *http.Request) {
	code := chi.URLParam(req, "code")
	redirect, err := h.redirectService.Find(code)
	if err != nil {
		if errors.Cause(err) == shortener.ErrRedirectNotFound {
			http.Error(res, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError )
		return
	}
	http.Redirect(res, req, redirect.URL, http.StatusMovedPermanently)
}

func (h handler) POST(res http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")
	requestBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	var serializer shortener.Serializeer
	if contentType == "application/msg-pack" {
		serializer = msgpack.Redirect{}		
	} else {
		serializer = json.Redirect{}
	}
	redirect, err := serializer.Decode(requestBody)
	if err != nil {
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	
	err = h.redirectService.Store(redirect)
	if err != nil {		
		if errors.Cause(err) == shortener.ErrRedirectInvalid {
			http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError )
		return
	}

	responseBody, err := serializer.Encode(redirect)
	if err != nil {
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", contentType)
	res.WriteHeader(http.StatusCreated)
	_, err = res.Write(responseBody)
	if err != nil {
		fmt.Println(err)
	}
}


package main

import (
	"encoding/json"
	"github.com/stevepartridge/go/log"
	"net/http"
	"strings"
	"time"
)

const (
	ContentType    = "Content-Type"
	ContentLength  = "Content-Length"
	ContentJSON    = "application/json"
	ContentHTML    = "text/html"
	ContentXHTML   = "application/xhtml+xml"
	ContentBinary  = "application/octet-stream"
	defaultCharset = "UTF-8"
)

type ApiResponse struct {
	Version  ApiVersion             `json:"version"`
	Error    bool                   `json:"error"`
	Message  string                 `json:"message,omitempty"`
	Messages []string               `json:"messages,omitempty"`
	Data     map[string]interface{} `json:"data"`
}

type ApiVersion struct {
	Current string `json:"current"`
	BuiltAt string `json:"built_at"`
}

var builtAt = "##BUILT_AT##"
var buildVersion = "##VERSION##"
var apiVersion = ApiVersion{
	"0.0.0",
	time.Now().Format("Jan 2, 2006 at 3:04pm (MST)"),
}

func init() {
	if !strings.Contains(builtAt, "BUILT_AT") {
		apiVersion.BuiltAt = builtAt
	}
	if !strings.Contains(buildVersion, "VERSION") {
		apiVersion.Current = buildVersion
	}
}

func (a *Api) ResponseJSON(res http.ResponseWriter, req *http.Request, status int, data map[string]interface{}) {

	response := ApiResponse{
		apiVersion,
		false,
		"",
		nil,
		data,
	}

	result, err := json.MarshalIndent(response, "", "  ")

	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}

	// json rendered fine, write out the result
	res.Header().Set(ContentType, ContentJSON)
	res.WriteHeader(status)
	res.Write(result)
}

func (a *Api) ResponseErrorJSON(res http.ResponseWriter, req *http.Request, status int, messages []string) {

	var msg = ""
	if len(messages) <= 0 {
		msg = messages[0]
		messages = nil
	}
	response := ApiResponse{
		apiVersion,
		true,
		msg,
		messages,
		nil,
	}

	result, err := json.MarshalIndent(response, "", "  ")

	if err != nil {
		log.Error(err)
		http.Error(res, err.Error(), 500)
		return
	}

	// json rendered fine, write out the result
	res.Header().Set(ContentType, ContentJSON)
	res.WriteHeader(status)
	res.Write(result)
}

func (a *Api) Entry(res http.ResponseWriter, req *http.Request) {
	log.Info("API Entry")
	a.ResponseJSON(res, req, 200, map[string]interface{}{"hello": "world"})
}

func (a *Api) NotFound(res http.ResponseWriter, req *http.Request) {
	log.Info("Not Found", "404", req.URL.Path)

	a.ResponseErrorJSON(res, req, 404, []string{"Not Found."})
}

package main

import (
	"github.com/emicklei/go-restful"
	"testing"
	"net/http"
	"fmt"
	"time"
)

func TestRouteExtractParameter(t *testing.T) {
	// setup service
	ws := new(restful.WebService)
	ws.Route(ws.GET("/test/{param}").To(DummyHandler))
	restful.Add(ws)

	http.ListenAndServe(":8080", nil)
}

func DummyHandler(rq *restful.Request, rp *restful.Response) {
	Result:= rq.PathParameter("param")
	fmt.Println(time.Now(),":",Result)
	rp.Write([]byte(Result))
}

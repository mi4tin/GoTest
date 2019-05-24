package main

import (
	"testing"
	"net/http"
)

func TestHttp1(t *testing.T){
	http.HandleFunc("/hello/", nil)
}
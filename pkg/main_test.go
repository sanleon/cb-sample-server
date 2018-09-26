package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
)

func TestDefaultHandler(t *testing.T)  {

	ts := httptest.NewServer(http.HandlerFunc(DefaultHandler))

	defer ts.Close()

	res, err := http.Get(ts.URL)

	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Error("Status code error")
		return
	}
	//b, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//if string(b) != ", Hello, World\nVersion:v0.1.0\n" {
	//	t.Fatal("response is not Hello, World")
	//}

}

func TestStatusHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(StatusHandler))

	defer ts.Close()

	res, err := http.Get(ts.URL)

	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Error("Status code error")
		return
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != "ok" {
		t.Fatal("response is not ok")
	}

}

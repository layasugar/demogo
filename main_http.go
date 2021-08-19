package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io"
	"log"
	"net/http"
)

const RespSuc = `{
    "data": {},
    "message": "操作成功",
    "status_code": 200
}`

func MainListenHttp() {
	http.HandleFunc("/reload", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(RespSuc))
	})
	http.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(RespSuc))
	})
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(RespSuc))
	})
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/test-hook", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err == nil {
			fmt.Println(string(body))
		}

		_, _ = w.Write([]byte(RespSuc))
	})
	log.Printf("http listen: %s", "0.0.0.0:10080")
	log.Panic(http.ListenAndServe("0.0.0.0:10080", nil))
}

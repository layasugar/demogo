package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		log.Print("asdasds")
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
	log.Printf("http listen: %s", "0.0.0.0:80")
	log.Panic(http.ListenAndServe("0.0.0.0:80", nil))
}

func MainGin() {
	r := gin.Default()
	r.Use(Cors()) //开启中间件 允许使用跨域请求

	var data = struct {
		Data       string `json:"data"`
		Message    string `json:"message"`
		StatusCode int64  `json:"status_code"`
	}{
		"我是data", "操作成功", 200,
	}

	r.POST("/test", func(ctx *gin.Context) {
		log.Print("有请求进来咯")
		ctx.JSON(http.StatusOK, &data)
	})

	r.Run(":80")
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			//接收客户端发送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token, session, Content-Type")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			//允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic info is: %v", err)
			}
		}()

		c.Next()
	}
}
